func isAnagram(s string, t string) bool {
    
    if len(s) != len(t) {
		return false
	}
	characterCounter := make(map[string]int)
	compareCharacterCounter := make(map[string]int)
	firstArray := strings.Split(s, "")
	lastArray := strings.Split(t, "")

	for _, character := range firstArray {
		characterCounter[character] = characterCounter[character] + 1
	}

	for _, character := range lastArray {
		compareCharacterCounter[character] = compareCharacterCounter[character] + 1
	}

	for key, element := range characterCounter {
		if compareCharacterCounter[key] != element {
			return false
		}
	}

	return true
}