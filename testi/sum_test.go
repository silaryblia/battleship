package main

import "testing"

func TestSum(t *testing.T) {
	// Arrange настройка тестированных данных

	testTable := []struct {
		name     string
		a        int
		b        int
		c        int
		expected int
	}{
		{"111", 1, 1, 1, 3},
		{
			name:     "123",
			a:        1,
			b:        2,
			c:        3,
			expected: 6},
		{"555", 5, 5, 5, 15}}

	for _, testCase := range testTable {
		//	t.Run(testCase.name, func(t *testing.T) {
		result := sum(testCase.a, testCase.b, testCase.c)
		t.Logf("Testing with a=%d, b=%d, c=%d, expected=%d", testCase.a, testCase.b, testCase.c, testCase.expected)
		if result != testCase.expected {
			t.Errorf("Incorrect result for %s. Expected %d, got %d", testCase.name, testCase.expected, result)
		}
	}
}

//
//	// Act вызов тестируемого кода
//	result := sum(a, b, c)
//
//	// Assert проверка возвращаемых результатов
//	if result != expected {
//		t.Errorf("Incorrect result. Expected %d, got %d", expected, result)
//	}
//}

//func TestSum(t *testing.T) {
//	// Arrange настройка тестированных данных
//	a := 1
//	b := 2
//	c := 12
//	expected := 15
//
//	// Act вызов тестируемого кода
//	result := sum(a, b, c)
//
//	// Assert проверка возвращаемых результатов
//	if result != expected {
//		t.Errorf("Incorrect result. Expected %d, got %d", expected, result)
//	}
//}
