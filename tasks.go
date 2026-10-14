package lesson2

// FizzBuzz возвращает FizzBuzz, Fizz, Buzz или пустую строку для обычного числа.
// Для n <= 0 возвращает invalid.
func FizzBuzz(n int) string {
	// TODO: Реализовать через switch без выражения.
	switch {
	case n <= 0:
		return "invalid"
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return ""
	}
}

// Sum возвращает сумму. Для пустого среза и nil возвращает 0.
func Sum(values []int64) int64 {
	// TODO: Пройти по значениям циклом.
	if len(values) == 0 {
		return 0
	}

	var result int64 = 0

	for _, value := range values {
		result += value
	}

	return result
}

// Min возвращает минимум. Для пустого среза и nil возвращает 0 по условию задачи.
func Min(values []int64) int64 {
	if len(values) == 0 {
		return 0
	}

	minimum := values[0]

	for _, value := range values {
		minimum = min(minimum, value)
	}

	return minimum
}

// Max возвращает максимум. Для пустого среза и nil возвращает 0 по условию задачи.
func Max(values []int64) int64 {
	if len(values) == 0 {
		return 0
	}

	maximum := values[0]

	for _, value := range values {
		maximum = max(maximum, value)
	}

	return maximum
}

// Average возвращает среднее как float64. Для пустого среза и nil возвращает 0.
func Average(values []int64) float64 {
	if len(values) == 0 {
		return 0
	}

	sum := Sum(values)

	return float64(sum) / float64(len(values))
}

// Category возвращает invalid, small, regular или large по сумме в копейках.
func Category(amount int64) string {
	switch {
	case amount >= 50_000:
		return "large"
	case amount >= 10_000:
		return "regular"
	case amount >= 0:
		return "small"
	default:
		return "invalid"
	}
}

// RuneCount считает кодовые точки Unicode в строке через range.
func RuneCount(text string) int {
	// TODO: Не использовать len(text) как количество кодовых точек.
	return len([]rune(text)) // :)
}
