package seedwork_test

import (
	"testing"

	"github.com/opencars/seedwork"
	"github.com/stretchr/testify/assert"
)

func Test_ToSnakeCase(t *testing.T) {
	tt := []struct {
		str      string
		expected string
	}{
		{
			str:      "VINs",
			expected: "vins",
		},
		{
			str:      "UserID",
			expected: "user_id",
		},
		{
			str:      "ID",
			expected: "id",
		},
		{
			str:      "VINDecoded",
			expected: "vin_decoded",
		},
	}

	for _, x := range tt {
		t.Run(x.str, func(t *testing.T) {
			actual := seedwork.ToSnakeCase(x.str)
			assert.Equal(t, x.expected, actual)
		})
	}
}
