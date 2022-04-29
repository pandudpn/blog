package hash_test

import (
	"errors"
	"os"
	"strconv"
	"testing"
	
	"github.com/pandudpn/blog/utils/hash"
	"github.com/stretchr/testify/assert"
)

func TestHash_Encode(t *testing.T) {
	var errs = errors.New("")
	testCases := []struct {
		name           string
		val            interface{}
		expectedResult string
		expectedError  error
	}{
		{
			name:           "Test Case #1 Success",
			val:            1,
			expectedResult: "KEjBzKxkvO2",
			expectedError:  errs,
		},
		{
			name:           "Test Case #2 Success with Int32",
			val:            int32(1),
			expectedResult: "KEjBzKxkvO2",
			expectedError:  errs,
		},
		{
			name:           "Test Case #3 Success with Int64",
			val:            int64(1),
			expectedResult: "KEjBzKxkvO2",
			expectedError:  errs,
		},
		{
			name:           "Test Case #4 Success with String",
			val:            "1",
			expectedResult: "KEjBzKxkvO2",
			expectedError:  errs,
		},
		{
			name:           "Test Case #5 error must be number",
			val:            "1abc",
			expectedResult: "",
			expectedError:  strconv.ErrSyntax,
		},
		{
			name:           "Test Case #6 error type unsupported",
			val:            float64(1),
			expectedResult: "",
			expectedError:  hash.ErrUnSupportedType,
		},
		{
			name:           "Test Case #7 error negative number",
			val:            -1,
			expectedResult: "",
			expectedError:  errors.New("negative number not supported"),
		},
		{
			name:           "Test Case #8 minimum length alphabet",
			val:            -1,
			expectedResult: "",
			expectedError:  errors.New("alphabet must contain at least 16 characters"),
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
			result, err := h.Encode(tc.val)
			
			assert.Equal(t, tc.expectedResult, result)
			assert.Error(t, tc.expectedError, err)
		})
	}
}
