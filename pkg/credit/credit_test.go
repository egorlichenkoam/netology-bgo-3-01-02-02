package credit_test

import (
	"01/pkg/credit"
	"fmt"
)

func ExampleCalculate() {
	fmt.Println(credit.Calculate(1_000_000_00, 36, 99))
	fmt.Println(credit.Calculate(4_000_000_00, 435, 46))
	fmt.Println(credit.Calculate(6_000_000_00, 45, 4))
	// Output: 8754457 215160452 315160452
	// 15333334 6270000290 6670000290
	// 14380492 47122140 647122140
}
