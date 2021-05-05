package helpers

func minInt(numbers []int) (minValue int) {
	if len(numbers) > 0 {
		minValue = numbers[0]
	}
	for i, e := range numbers {
		if i == 0 || e < minValue {
			minValue = e
		}
	}

	return
}

func maxInt(numbers []int) (maxValue int) {
	if len(numbers) > 0 {
		maxValue = numbers[0]
	}
	for i, e := range numbers {
		if i == 0 || e > maxValue {
			maxValue = e
		}
	}

	return
}
