package encrypt

func Nimbus(str string) string {
	result := ""
	for _, c := range(str) {
		asciiCode := int(c)
		newAsciiCode := asciiCode + 3
		result += string(newAsciiCode)
	}
	return result
}
