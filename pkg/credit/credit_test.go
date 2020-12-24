package credit_test

import (
	"01/pkg/credit"
	"fmt"
)

func ExampleCalculate() {
	fmt.Println(credit.Calculate(1_000_000_00, 36, 99))
	fmt.Println(credit.Calculate(3_000_000_00, 435, 46))
	fmt.Println(credit.Calculate(5_000_000_00, 45, 4))
	// Output: 8754457 215160452 315160452
	// 11500000 4702500000 5002500000
	// 11983743 39268435 539268435
}
