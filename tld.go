package validate

func IsValidTLD(str string) bool {
	length := len(str)
	match := 0
	for _, tld := range tldList {
		if len(tld) != length {
			continue
		}
		match = 0
		for i := 0; i < length; i++ {
			if tld[i] == str[i] || tld[i]+32 == str[i] {
				match++
			}
		}
		if match == length {
			break
		}
	}
	return match == length
}
