package parse // import "github.com/tdewolff/parse"

// Copy returns a copy of the given byte slice.
func Copy(src []byte) (dst []byte) {
	dst = make([]byte, len(src))
	copy(dst, src)
	return
}

// ToLower converts all characters in the byte slice from A-Z to a-z.
func ToLower(src []byte) []byte {
	for i := 0; i < len(src); i++ {
		c := src[i]
		if c >= 'A' && c <= 'Z' {
			src[i] = c + ('a' - 'A')
		}
	}
	return src
}

// Equal returns true when s matches the target.
func Equal(s, target []byte) bool {
	if len(s) != len(target) {
		return false
	}
	for i := 0; i < len(target); i++ {
		if s[i] != target[i] {
			return false
		}
	}
	return true
}

// EqualFold returns true when s matches case-insensitively the targetLower (which must be lowercase).
func EqualFold(s, targetLower []byte) bool {
	if len(s) != len(targetLower) {
		return false
	}
	for i := 0; i < len(targetLower); i++ {
		c := targetLower[i]
		if s[i] != c && (c < 'A' && c > 'Z' || s[i]+('a'-'A') != c) {
			return false
		}
	}
	return true
}

var whitespaceTable = [256]bool{
	// ASCII
	false, false, false, false, false, false, false, false,
	false, true, true, false, true, true, false, false, // tab, new line, form feed, carriage return
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	true, false, false, false, false, false, false, false, // space
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	// non-ASCII
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
}

// IsWhitespace returns true for space, \n, \r, \t, \f.
func IsWhitespace(c byte) bool {
	return whitespaceTable[c]
}

// IsAllWhitespace returns true when the entire byte slice consists of space, \n, \r, \t, \f.
func IsAllWhitespace(b []byte) bool {
	for i := 0; i < len(b); i++ {
		if !IsWhitespace(b[i]) {
			return false
		}
	}
	return true
}

// TrimWhitespace removes any leading and trailing whitespace characters.
func TrimWhitespace(b []byte) []byte {
	n := len(b)
	start := n
	for i := 0; i < n; i++ {
		if !IsWhitespace(b[i]) {
			start = i
			break
		}
	}
	end := n
	for i := n - 1; i >= start; i-- {
		if !IsWhitespace(b[i]) {
			end = i + 1
			break
		}
	}
	return b[start:end]
}

// ReplaceMultipleWhitespace replaces character series of space, \n, \t, \f, \r into a single space or newline (when the serie contained a \n or \r).
func ReplaceMultipleWhitespace(b []byte) []byte {
	j := 0
	prevWS := false
	hasNewline := false
	for i := 0; i < len(b); i++ {
		c := b[i]
		if IsWhitespace(c) {
			prevWS = true
			if c == '\n' || c == '\r' {
				hasNewline = true
			}
		} else {
			if prevWS {
				prevWS = false
				if hasNewline {
					hasNewline = false
					b[j] = '\n'
				} else {
					b[j] = ' '
				}
				j++
			}
			b[j] = b[i]
			j++
		}
	}
	if prevWS {
		if hasNewline {
			b[j] = '\n'
		} else {
			b[j] = ' '
		}
		j++
	}
	return b[:j]
}
