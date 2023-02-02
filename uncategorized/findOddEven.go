package uncategorized

func One(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func Two(number int) string {
	switch number % 2 {
	case 0:
		return "Even"
	default:
		return "Odd"
	}
}

func Three(number int) string {
	return [2]string{"Even", "Odd"}[number&1]
}

func Four(number int) string {
	return map[bool]string{true: "Even", false: "Odd"}[number%2 == 0]
}

func Five(number int) string {
	parity := map[int]string{0: "Even", 1: "Odd"}
	return parity[number%2]
}
