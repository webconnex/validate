package validate

func IsValidEmail(email string) bool {
	var count = len(email)
	if count == 0 {
		return false
	}
	var last rune
	var at, dot int
	for i, c := range email {
		if c == ',' ||
			c == '/' ||
			c == '`' ||
			(c >= ':' && c <= '?') ||
			(c >= '[' && c <= '^') ||
			(c <= '*' || c >= '{') {
			break // Invalid characters
		}
		if c == '.' {
			if count == 1 || // '.' is last
				last == 0 || // '.' is first
				last == '.' || // '.' after '.'
				last == '@' || // '.' is after '@'
				(at > 0 && last == '-') { // '.' is after '-' in domain
				break
			}
			dot = i
		} else if c == '-' {
			if count == 1 || // '-' is last
				last == '@' || // '-' is after '@'
				(at > 0 && last == '.') { // '-' is after '.' in domain
				break
			}
		} else if c == '@' {
			if count == 1 || // '@' is last
				last == 0 || // '@' is first
				at > 0 { // '@' appears twice
				break
			}
			at = i
		} else if at > 0 && (c == '_' || c == '+') {
			break // Invalid characters in domain
		}
		last = c
		count--
	}
	return count == 0 && // Reached End
		at > 0 && // We have a '@'
		dot > at && // We have a '.' after '@'
		IsValidTLD(email[dot+1:])
}