package generator

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func randomNumbers(totalDigit int) (int, error) {
	if totalDigit <= 0 {
		return 0, fmt.Errorf("total digit should be a positive integer")
	}

	rand.Seed(time.Now().UnixNano())

	minVal := int(math.Pow10(totalDigit - 1))
	maxVal := int(math.Pow10(totalDigit)) - 1

	return minVal + rand.Intn(maxVal-minVal+1), nil
}
