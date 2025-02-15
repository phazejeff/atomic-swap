package coins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimalCtx(t *testing.T) {
	c := DecimalCtx()
	assert.Equal(t, c.Precision, uint32(MaxCoinPrecision))
	// verify that package variable decimalCtx is unmodified, because c is a copy
	c.Precision = 3
	assert.Equal(t, decimalCtx.Precision, uint32(MaxCoinPrecision))
}
