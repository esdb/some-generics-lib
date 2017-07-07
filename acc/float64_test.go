package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
)

func Test_float64_kind(t *testing.T) {
	should := require.New(t)
	v := float64(1)
	should.Equal(lang.Float64, objAcc(v).Kind())
}

func Test_float64_gostring(t *testing.T) {
	should := require.New(t)
	v := float64(1)
	should.Equal("float64", objAcc(v).GoString())
}

func Test_float64_get_float64(t *testing.T) {
	should := require.New(t)
	v := float64(1)
	should.Equal(float64(1), objAcc(v).Float64(objPtr(v)))
}

func Test_float64_set_float64(t *testing.T) {
	should := require.New(t)
	v := float64(1)
	should.Panics(func() {
		objAcc(v).SetInt(objPtr(v), 2)
	})
}
