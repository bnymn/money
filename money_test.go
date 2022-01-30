package money

import (
	"testing"
)

func TestNewMoney(t *testing.T) {
	money := NewMoney("USD", 100)

	if money.currencyCode != "USD" {
		t.Error("currencyCode is not set correctly")
	}

	if money.amountInCents != 100 {
		t.Error("amountInCents is not set correctly")
	}
}

func TestGetIntegralPart(t *testing.T) {
	var testDataTable = []struct {
		currency      string
		amountInCents int64
		result        int64
	}{
		{"USD", 9820, 98},
		{"USD", 90001, 900},
		{"USD", 9020, 90},
		{"USD", 92233, 922},
	}

	for _, testData := range testDataTable {
		t.Run("Testing GetIntegralPart", func(t *testing.T) {
			money := NewMoney(testData.currency, testData.amountInCents)
			integralPart := GetIntegralPart(*money)
			if integralPart != testData.result {
				t.Errorf("integral part (%d) is not calculated correctly", integralPart)
			}
		})
	}
}

func TestGetFactionalPart(t *testing.T) {
	var testDataTable = []struct {
		currency      string
		amountInCents int64
		result        int64
	}{
		{"USD", 9820, 20},
		{"USD", 9821, 21},
		{"USD", 90001, 1},
		{"USD", 1023, 23},
		{"USD", 9820, 20},
		{"USD", 9802, 2},
		{"USD", 5500, 0},
	}

	for _, testData := range testDataTable {
		t.Run("Testing GetFractinalPart", func(t *testing.T) {
			money := NewMoney(testData.currency, testData.amountInCents)
			fractionalPart := GetFractionalPart(*money)
			if fractionalPart != testData.result {
				t.Errorf("fractional part (%d) is not calculated correctly", fractionalPart)
			}
		})
	}
}

func TestSumMoney(t *testing.T) {
	var testDataTable = []struct {
		m1             Money
		m2             Money
		expectedResult *Money
	}{
		{*NewMoney("USD", 9820), *NewMoney("USD", 5523), NewMoney("USD", 15343)},
		{*NewMoney("USD", 100), *NewMoney("USD", 1), NewMoney("USD", 101)},
		{*NewMoney("USD", 9820), *NewMoney("USD", 5500), NewMoney("USD", 15320)},
		{*NewMoney("USD", 9820), *NewMoney("EUR", 5500), nil},
	}
	for _, testData := range testDataTable {
		t.Run("Testing SumMoney", func(t *testing.T) {
			actualResult, error := SumMoney(testData.m1, testData.m2)

			if error != nil && actualResult != nil {
				t.Error("the result should be nil in case of error")
			}

			if actualResult != nil && actualResult.amountInCents != testData.expectedResult.amountInCents {
				t.Errorf("amount does not match (%d != %d)", actualResult.amountInCents, testData.expectedResult.amountInCents)
			}
		})
	}
}

func TestSubtractMoney(t *testing.T) {
	var testDataTable = []struct {
		m1             Money
		m2             Money
		expectedResult *Money
	}{
		{*NewMoney("USD", 9820), *NewMoney("USD", 5523), NewMoney("USD", 4297)},
		{*NewMoney("USD", 100), *NewMoney("USD", 1), NewMoney("USD", 99)},
		{*NewMoney("USD", 9820), *NewMoney("USD", 5500), NewMoney("USD", 4320)},
		{*NewMoney("USD", 9820), *NewMoney("EUR", 5500), nil},
	}
	for _, testData := range testDataTable {
		t.Run("Testing SumMoney", func(t *testing.T) {
			actualResult, error := SubtractMoney(testData.m1, testData.m2)

			if error != nil && actualResult != nil {
				t.Error("the result should be nil in case of error")
			}

			if actualResult != nil && actualResult.amountInCents != testData.expectedResult.amountInCents {
				t.Errorf("amount does not match (%d != %d)", actualResult.amountInCents, testData.expectedResult.amountInCents)
			}
		})
	}
}
