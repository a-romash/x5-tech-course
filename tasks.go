package homework

// PairInput соответствует объекту nums/target из условия контеста.
type PairInput struct {
	Nums   []int
	Target int
}

// longestPrefix возвращает самый длинный общий префикс строк.
func longestPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}

	prefix := words[0]
	for _, word := range words {
		prefix = getPrefix(prefix, word)

		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func getPrefix(word, other string) string {
	prefix := ""

	runeWord, runeOther := []rune(word), []rune(other)
	for idx := range min(len(runeWord), len(runeOther)) {
		if runeWord[idx] != runeOther[idx] {
			return prefix
		}

		prefix += string(runeWord[idx])
	}

	return prefix
}

// pairIndices возвращает индексы двух разных элементов по возрастанию.
func pairIndices(input PairInput) [2]int {
	remained := make(map[int]int)

	for idx, val := range input.Nums {
		if _, ok := remained[val]; ok {
			return [2]int{remained[val], idx}
		}

		remained[input.Target - val] = idx
	}

	return [2]int{}
}

// hasDuplicate сообщает, встречается ли какое-либо значение повторно.
func hasDuplicate(nums []int) bool {
	numCounter := make(map[int]int)
	
	for _, num := range nums {
		numCounter[num]++

		if numCounter[num] > 1 {
			return true
		}
	}

	return false
}

// compactSorted создаёт независимый слайс уникальных значений.
func compactSorted(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	
	compacted := make([]int, 1, len(nums))
	compacted[0] = nums[0]

	for _, num := range nums {
		if compacted[len(compacted) - 1] != num {
			compacted = append(compacted, num)
		}
	}

	return compacted
}

// balancedBrackets проверяет пары и порядок закрытия скобок.
func balancedBrackets(text string) bool {
	if len(text) == 0 {
		return true
	}

	closedToOpen := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}

	stack := make([]rune, 0, len(text)/2)
	for _, bracket := range []rune(text) {
		if _, ok := closedToOpen[bracket]; !ok {
			stack = append(stack, bracket)
			continue
		}

		if len(stack) == 0 || stack[len(stack) - 1] != closedToOpen[bracket] {
			return false
		}

		stack = stack[:len(stack) - 1]
	}

	return len(stack) == 0
}
