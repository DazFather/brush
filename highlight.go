package brush

import (
	"fmt"
	"regexp"
)

type section struct {
	from, to int
	style
}

func (b Brush[color]) newSection(from, to int) section {
	return section{
		from:  from,
		to:    to,
		style: b.extract(),
	}
}

func (b Painted) newSection(from int) section {
	return section{
		from:  from,
		to:    from + len(b.content),
		style: b.style,
	}
}

func (s section) evaluateOn(content string) string {
	return s.style.apply(content[s.from:s.to])
}

func (s *section) shift(offset int) {
	s.from += offset
	s.to += offset
}

// Highlighted represents a string that contains informations about
// the foreground and background color output of some parts of it.
// the styling can also be different for different subset of the content
type Highlighted struct {
	sectors []section
	content string
}

// Join different values toghether into a single item maintaining all styling
func Join(values ...any) (joined Highlighted) {
	joined.Append(values...)
	return
}

// Highlight only the matching part of the given string with the color of the brush
func (b Brush[color]) Highlight(s string, find *regexp.Regexp) Highlighted {
	var result = Highlighted{content: s}

	found := find.FindAllStringIndex(s, -1)
	if found == nil {
		return result
	}

	result.sectors = make([]section, len(found))
	for i, indexes := range found {
		result.sectors[i] = b.newSection(indexes[0], indexes[1])
	}

	return result
}

// HighlightFunc is like Highlight but allows you to replace the part that will match
// the returned string of the repl function will then be highlighted using brush styling
func (b Brush[color]) HighlightFunc(s string, find *regexp.Regexp, repl func(string) string) Highlighted {
	var result = Highlighted{}

	found := find.FindAllStringIndex(s, -1)
	if found == nil {
		return result
	}

	var from, to int

	for _, indexes := range found {
		from = indexes[0]
		if to < indexes[1] {
			result.content += s[to:from]
		}
		to = indexes[1]

		replacement := repl(s[from:to])
		if replSize := len(replacement); replSize > 0 {
			size := len(result.content)
			result.sectors = append(result.sectors, b.newSection(
				size,
				size+replSize,
			))
			result.content += replacement
		}
	}

	if to < len(s) {
		result.content += s[to:]
	}

	return result
}

// Embed lets create an Highlighted item by joining the given values.
// Painted items will maintain their styling.
// Highlighted items will maintain their styling only on the subset that contains info about it,
// for other values (and the subset of Highlighted items that do not specify info) the brush style will be enforced
func (b Brush[color]) Embed(values ...any) Highlighted {
	var res Highlighted

	for _, rawValue := range values {
		size := len(res.content)

		switch v := rawValue.(type) {
		case Painted:
			res.addSections(v.newSection(size))
			res.content += v.content
		case Highlighted:
			last := size
			for _, sec := range v.sectors {
				sec.shift(size)
				if last < sec.from {
					res.addSections(b.newSection(last, sec.from))
				}
				res.addSections(sec)
				last = sec.to
			}
			res.content += v.content
			if size = len(res.content); last < size {
				res.addSections(b.newSection(last, size))
			}
		case string:
			res.addSections(b.newSection(size, size+len(v)))
			res.content += v
		case fmt.Stringer:
			s := v.String()
			res.addSections(b.newSection(size, size+len(s)))
			res.content += s
		default:
			s := fmt.Sprint(v)
			res.addSections(b.newSection(size, size+len(s)))
			res.content += s
		}

	}

	return res
}

func (h *Highlighted) addSections(s ...section) {
	h.sectors = append(h.sectors, s...)
}

// Append lets you add some items at the end of the Highlighted content
func (h *Highlighted) Append(values ...any) {
	for i := range values {
		size := len(h.content)

		switch v := values[i].(type) {
		case Painted:
			h.addSections(v.newSection(size))
			h.content += v.content
		case Highlighted:
			for _, sec := range v.sectors {
				sec.shift(size)
				h.addSections(sec)
			}
			h.content += v.content
		default:
			h.content += fmt.Sprint(v)
		}
	}
}

// String evaluates the content by applying the different styling where specified
func (h Highlighted) String() string {
	var (
		res  string
		last int
	)

	for _, sec := range h.sectors {
		if last < sec.from {
			res += h.content[last:sec.from]
		}
		res += sec.evaluateOn(h.content)
		last = sec.to
	}

	if last < len(h.content) {
		res += h.content[last:]
	}

	return res
}
