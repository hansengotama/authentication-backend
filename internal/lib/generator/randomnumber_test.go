package generator

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
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
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedErr.Error())
				return
			}

			generatedNumStr := strconv.Itoa(generatedNum)
			assert.Equal(t, tc.totalDigits, len(generatedNumStr))
			assert.NoError(t, err)
		})
	}
}
