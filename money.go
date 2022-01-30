package money

import (
	"errors"
	"math"
)

type Money struct {
	currencyCode  string
	amountInCents int64
}

func NewMoney(currencyCode string, amountInCents int64) *Money {
	return &Money{
		currencyCode:  currencyCode,
		amountInCents: amountInCents,
	}
}

func GetIntegralPart(m Money) int64 {
	return m.amountInCents / getPowerOfTen(2)
}

func GetFractionalPart(m Money) int64 {
	return m.amountInCents % getPowerOfTen(2)
}

func SumMoney(m1 Money, m2 Money) (*Money, error) {
	if m1.currencyCode != m2.currencyCode {
		return nil, errors.New("currencyCode does not match")
	}

	return &Money{
		currencyCode:  m1.currencyCode,
		amountInCents: m1.amountInCents + m2.amountInCents,
	}, nil
}

func SubtractMoney(m1 Money, m2 Money) (*Money, error) {
	if m1.currencyCode != m2.currencyCode {
		return nil, errors.New("currencyCode does not match")
	}

	return &Money{
		currencyCode:  m1.currencyCode,
		amountInCents: m1.amountInCents - m2.amountInCents,
	}, nil
}

func getPowerOfTen(precision uint8) int64 {
	return int64((math.Pow(10, float64(precision))))
}
