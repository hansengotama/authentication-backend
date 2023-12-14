package generator

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

func RandomNumbers(totalDigit int) (int, error) {
	if totalDigit <= 0 {
		return 0, errors.New("total digit should be a positive integer")
	}

	rand.Seed(time.Now().UnixNano())

	minVal := int(math.Pow10(totalDigit - 1))
	maxVal := int(math.Pow10(totalDigit)) - 1

	return minVal + rand.Intn(maxVal-minVal+1), nil
}
