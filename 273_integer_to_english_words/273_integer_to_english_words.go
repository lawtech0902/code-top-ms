package integertoenglishwords

import "strings"

func numberToWords(num int) string {
	const (
		BILLION  = 1000000000
		MILLION  = 1000000
		THOUSAND = 1000
		HUNDRED  = 100
		TEN      = 10
	)

	m := map[int]string{
		0:  "Zero",
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
		20: "Twenty",
		30: "Thirty",
		40: "Forty",
		50: "Fifty",
		60: "Sixty",
		70: "Seventy",
		80: "Eighty",
		90: "Ninety",
	}

	if v, ok := m[num]; ok {
		return v
	}

	var res []string
	// billion
	base := num / BILLION
	num = num % BILLION
	if base > 0 {
		res = append(res, numberToWords(base))
		res = append(res, "Billion")
	}

	// million
	base = num / MILLION
	num = num % MILLION
	if base > 0 {
		res = append(res, numberToWords(base))
		res = append(res, "Million")
	}

	// thousand
	base = num / THOUSAND
	num = num % THOUSAND
	if base > 0 {
		res = append(res, numberToWords(base))
		res = append(res, "Thousand")
	}

	// hundred
	base = num / HUNDRED
	num = num % HUNDRED
	if base > 0 {
		res = append(res, numberToWords(base))
		res = append(res, "Hundred")
	}

	// tens
	if num >= 20 {
		base = num / TEN
		num = num % TEN
		res = append(res, numberToWords(base*TEN))
	}

	// ones
	if num > 0 {
		res = append(res, numberToWords(num))
	}

	return strings.Join(res, " ")
}
