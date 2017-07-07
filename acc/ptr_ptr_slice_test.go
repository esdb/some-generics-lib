package acc

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/lang"
	"testing"
	"unsafe"
)

func Test_ptr_ptr_slice_kind(t *testing.T) {
	should := require.New(t)
	v1 := []int{}
	v2 := &v1
	v := &v2
	should.Equal(lang.Array, objAcc(v).Kind())
}

func Test_ptr_ptr_slice_gostring(t *testing.T) {
	should := require.New(t)
	v1 := []int{}
	v2 := &v1
	v := &v2
	should.Equal("**[]int", objAcc(v).GoString())
}

func Test_ptr_ptr_slice_elem(t *testing.T) {
	should := require.New(t)
	v1 := []int{}
	v2 := &v1
	v := &v2
	should.Equal(lang.Int, objAcc(v).Elem().Kind())
}

func Test_ptr_ptr_slice_random_accessible(t *testing.T) {
	should := require.New(t)
	v1 := []int{}
	v2 := &v1
	v := &v2
	should.True(objAcc(v).RandomAccessible())
}

func Test_ptr_ptr_slice_get_by_array_index(t *testing.T) {
	should := require.New(t)
	v1 := []int{1, 2, 3}
	v2 := &v1
	v := &v2
	elem := objAcc(v).ArrayIndex(objPtr(v), 1)
	should.Equal(2, objAcc(v).Elem().Int(elem))
}

func Test_ptr_ptr_slice_set_by_array_index(t *testing.T) {
	should := require.New(t)
	v1 := []int{1, 2, 3}
	v2 := &v1
	v := &v2
	elem := objAcc(v).ArrayIndex(objPtr(v), 1)
	objAcc(v).Elem().SetInt(elem, 4)
	should.Equal(4, v1[1])
}

func Test_ptr_ptr_slice_iterate_array(t *testing.T) {
	should := require.New(t)
	v1 := []int{1, 2, 3}
	v2 := &v1
	v := &v2
	elems := []int{}
	objAcc(v).IterateArray(objPtr(v), func(index int, elem unsafe.Pointer) bool {
		elems = append(elems, objAcc(v).Elem().Int(elem))
		return true
	})
	should.Equal(**v, elems)
}

func Test_ptr_ptr_slice_fill_array(t *testing.T) {
	should := require.New(t)
	v1 := []int{}
	v2 := &v1
	v := &v2
	objAcc(v).FillArray(objPtr(v), func(filler lang.ArrayFiller) {
		index, elem := filler.Next()
		should.Equal(0, index)
		objAcc(v).Elem().SetInt(elem, 1)
		filler.Fill()
		index, elem = filler.Next()
		should.Equal(1, index)
		objAcc(v).Elem().SetInt(elem, 2)
		filler.Fill()
	})
	should.Equal([]int{1, 2}, v1)
}
