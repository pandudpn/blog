package hash_test

import (
	"errors"
	"os"
	"testing"
	
	"github.com/pandudpn/blog/utils/hash"
	"github.com/stretchr/testify/assert"
)

func TestHash_Decode(t *testing.T) {
	testCases := []struct {
		name           string
		val            string
		expectedResult int
		expectedError  error
	}{
		{
			name:           "Test Case #1 Success",
			val:            "KEjBzKxkvO2",
			expectedResult: 1,
			expectedError:  nil,
		},
		{
			name:           "Test Case #2 Error hash wrong",
			val:            "KEjBzKxkvO",
			expectedResult: 0,
			expectedError:  errors.New("mismatch between encode and decode: KEjBzKxkvO start KEjBzKxkvO2 re-encoded. result: [1]"),
		},
		{
			name:           "Test Case # Error hash wrong",
			val:            "KEjBzKxkvO",
			expectedResult: 0,
			expectedError:  hash.ErrData,
		},
	}
	
	for k, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if k == len(testCases)-1 {
				os.Setenv("HASH_IDS_ALPHABET", "abc")
			} else {
				os.Setenv("HASH_IDS_ALPHABET", "")
			}
			os.Setenv("SALT_IDS", "unit-test")
			h := hash.NewHash()
			result, err := h.Decode(tc.val)
			
			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
