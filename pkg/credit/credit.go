package credit

import (
	"math"
)

func Calculate(totalSumInPenny int64, termInMonths int64, percent int64) (monthlyPayment int64, overpayment int64, totalPayment int64) {
	I := float64(percent) / 12 / 100
	K := (I * math.Pow(1+I, float64(termInMonths))) / (math.Pow(1+I, float64(termInMonths)) - 1)
	monthlyPayment = int64(K * float64(totalSumInPenny))
	totalPayment = monthlyPayment * termInMonths
	overpayment = totalPayment - totalSumInPenny
	return
}
