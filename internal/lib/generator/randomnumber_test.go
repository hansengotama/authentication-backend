package generator

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func Test_GenerateRandomNumbers(t *testing.T) {
	testCases := []struct {
		totalDigits int
		expectedErr error
	}{
		{totalDigits: -1, expectedErr: errors.New("total digit should be a positive integer")},
		{totalDigits: 4},
		{totalDigits: 6},
		{totalDigits: 8},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Total digits: %d", tc.totalDigits), func(t *testing.T) {
			generatedNum, err := RandomNumbers(tc.totalDigits)
			if tc.expectedErr != nil {
				if err.Error() != tc.expectedErr.Error() {
					t.Errorf("Unmatch error: %v with :%v", err, tc.expectedErr)
				}

				return
			}

			if err != nil {
				t.Errorf("Error encountered: %v", err)
			}

			generatedNumStr := strconv.Itoa(generatedNum)
			if len(generatedNumStr) != tc.totalDigits {
				t.Errorf("Generated number %s does not have %d digits", generatedNumStr, tc.totalDigits)
			}
		})
	}
}
