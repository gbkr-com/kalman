package kalman

import (
	"fmt"
	"math"
	"testing"
)

func TestWithPrices(t *testing.T) {
	prices := []float64{
		83.55,
		83.45,
		83.45,
		83.4,
		83.45,
		83.45,
		83.45,
		83.45,
		83.4,
		83.35,
		83.35,
		83.4,
		83.4,
		83.45,
		83.4,
		83.45,
		83.45,
		83.4,
		83.4,
		83.4,
		83.3,
		83.25,
		83.2,
		83.15,
		83.1,
		83.2,
		83.2,
		83.2,
		83.2,
		83.2,
	}
	var k *Filter
	for i, p := range prices {
		if i == 0 {
			k = New(p)
		} else {
			k.Observe(p)
			fmt.Printf("%f %f\n", p, k.X)
			if i == 29 {
				//
				// The estimate should trend to the same value predicted by the
				// 'q' implementation on gbkr.com.
				//
				if math.Abs(k.X-83.19944) > 0.00001 {
					t.Error()
				}
			}
		}
	}
}
