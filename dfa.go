//go:build dfa

package dfaregexp

const (
	numState = 6

	q0 = iota // start: no char read yet.
	q1        // valid, last char was letter or digit.
	q2        // started with symbol, need at least one letter/digit.
	q3        // valid, last char was underscore '_'.
	q4        // last char was hyphen '-', not accepting (can't be suffix).
	qx        // trap.
)

var (
	accepted uint8
	dfa      [numState][256]uint8
)

func init() {
	accepted = (1 << q1) | (1 << q3)
	for q := range numState {
		for b := range 256 {
			c := byte(b)
			dfa[q][b] = qx
			switch q {
			default:
				continue
			case q0:
				if isLetter(c) {
					dfa[q][b] = q1
				} else if c == '_' || c == '-' {
					// need letter/digit for next char.
					dfa[q][b] = q2
				}
			case q1:
				if isLetter(c) || isDigit(c) {
					dfa[q][b] = q1
				} else if c == '_' {
					dfa[q][b] = q3
				} else if c == '-' {
					dfa[q][b] = q4
				}
			case q2, q3, q4:
				if isLetter(c) || isDigit(c) {
					dfa[q][b] = q1
				}
			}
		}
	}
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func Valid(s string) bool {
	if len(s) == 0 {
		return false
	}
	state := uint8(q0)
	for _, c := range s {
		state = dfa[state][c]
		if state == qx {
			return false
		}
	}
	return (accepted>>state)&1 != 0
}
