package main


type (
	tag byte

	token struct {
		typ tag
		val string
	}

	lexer struct {
		src  []byte
		ch   rune
		pos  int
		rd   int

		ln   int
		col  int
	}

)

func (l *lexer) look() (c rune, w int) {
	c, w = rune(l.src[l.rd]), 1

	if l.rd < len(l.src) && c >= utf8.RuneSelf {
		c, w = utf8.DecodeRune(l.src[l.rd:])
	}

	return
}

func (l *lexer) line() {
	l.col += 1

	if l.ch == '\n' {
		l.ln += 1
		l.col = 0
	}
}

func (l *lexer) step() {
	c, w := l.look()

	l.pos = l.rd
	l.rd += w
	l.ch = c
	l.line()
}

var keywords = map[string]tag {}

func (l *lexer) ident() (tag, string) {
	p := l.pos
	s := l.src[p:l.pos]

	if tk, ok := keywords[s]; ok {
		return tk, s
	}
	
	return NAME, s
}


