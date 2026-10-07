func isPalindrome(s string) bool {
	i := 0
	j := len(s)-1

	for i < j{
		l := s[i]
		r := s[j]
		if !checkCharacterValid(l) {
			i++
			continue
		}else if !checkCharacterValid(r) {
			j--
			continue
		}
		if toLower(l) != toLower(r) {
			return false
		}
		i++
		j--
	}
	return true
}

func checkCharacterValid (r byte)bool{
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func toLower(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        return c + ('a' - 'A')
    }
    return c
}