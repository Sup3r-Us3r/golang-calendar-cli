package util

// RepeatString: Returns the entered value repeatedly
func RepeatString(value string, count int) string {
	currentValue := value

	for i := 0; i <= count; i++ {
		currentValue += value
	}

	return currentValue
}
