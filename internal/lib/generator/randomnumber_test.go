package generator

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_GenerateRandomNumbers(t *testing.T) {
	testCases := []struct {
		totalDigits int
	}{
		{totalDigits: 4},
		{totalDigits: 6},
		{totalDigits: 8},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Total digits: %d", tc.totalDigits), func(t *testing.T) {
			generatedNum, err := randomNumbers(tc.totalDigits)

			if err != nil {
				t.Errorf("Error encountered: %v", err)
			}

			generatedNumStr := strconv.Itoa(generatedNum)
			fmt.Println(generatedNumStr)
			if len(generatedNumStr) != tc.totalDigits {
				t.Errorf("Generated number %s does not have %d digits", generatedNumStr, tc.totalDigits)
			}
		})
	}
}
