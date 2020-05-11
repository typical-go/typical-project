package typcfg_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typcfg"
)

func TestWrite(t *testing.T) {
	testcases := []struct {
		typcfg.Config
		before      string
		expected    string
		expectedErr string
	}{
		{
			Config: &typcfg.Configuration{
				Name: "TEST",
				Spec: &someSpec{},
			},
			expected: "TEST_FIELD1=defaulValue1\nTEST_FIELD2=defaulValue2\n",
		},
		{
			Config: &typcfg.Configuration{
				Name: "TEST",
				Spec: &someSpec{},
			},
			before:   "XXXX=XXXX",
			expected: "XXXX=XXXX\nTEST_FIELD1=defaulValue1\nTEST_FIELD2=defaulValue2\n",
		},
	}

	for i, tt := range testcases {
		dest := fmt.Sprintf("write%d.env", i)
		defer os.Remove(dest)

		if tt.before != "" {
			ioutil.WriteFile(dest, []byte(tt.before), 0777)
		}

		err := typcfg.Write(dest, tt.Config)
		if tt.expectedErr != "" {
			require.EqualError(t, err, tt.expectedErr)
		} else {
			require.NoError(t, err)
		}

		b, _ := ioutil.ReadFile(dest)
		require.Equal(t, tt.expected, string(b))
	}
}

type someSpec struct {
	Field1 string `default:"defaulValue1"`
	Field2 string `default:"defaulValue2"`
}
