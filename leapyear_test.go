package main

import "testing"

func TestLeapYear(t *testing.T) {
	tests := []struct {
		year     int
		expected bool
	}{
		{2000, true},  // 被 400 整除 → 闰年
		{1900, false}, // 被 100 整除但不被 400 整除 → 平年
		{2024, true},  // 被 4 整除但不被 100 整除 → 闰年
		{2023, false}, // 不能被 4 整除 → 平年
		{1, false},    // 公元1年不是闰年
		{1600, true},  // 被 400 整除 → 闰年
	}

	for _, test := range tests {
		actual := leapyear(test.year)
		if actual != test.expected {
			t.Errorf("leapyear(%d) = %v; want %v", test.year, actual, test.expected)
		}
	}
}
