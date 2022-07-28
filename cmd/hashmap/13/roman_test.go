package _3

import "testing"

func TestRoman(t *testing.T) {
	testcases := []struct {
		number int
		roman  string
	}{
		{0, ""},
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{9, "IX"},
		{14, "XIV"},
		{19, "XIX"},
		{24, "XXIV"},
		{34, "XXXIV"},
		{39, "XXXIX"},
		{40, "XL"},
		{41, "XLI"},
		{44, "XLIV"},
		{44, "XLIV"},
		{99, "XCIX"},
		{400, "CD"},
		{900, "CM"},
		{1990, "MCMXC"},
		{1993, "MCMXCIII"},
		{2018, "MMXVIII"},
		{1111, "MCXI"},
		{2222, "MMCCXXII"},
		{444, "CDXLIV"},
		{555, "DLV"},
		{666, "DCLXVI"},
		{999, "CMXCIX"},
	}

	for _, testcase := range testcases {
		roman := intToRoman(testcase.number)
		if roman != testcase.roman {
			t.Errorf("%d expected to convert to %s, got: %s.", testcase.number, testcase.roman, roman)
		}
	}
}
