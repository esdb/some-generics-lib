package max

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/google/gofuzz"
)

func Test_int(t *testing.T) {
	should := require.New(t)
	should.Equal(3, DoMax([]interface{}{1, 3, 2}))
}

func max_int_typed(collection []int) int {
	currentMax := collection[0]
	for _, elem := range collection[1:] {
		if elem > currentMax {
			currentMax = elem
		}
	}
	return currentMax
}

func Benchmark_int(b *testing.B) {
	fuzzer := fuzz.New()
	datasets := make([][]interface{}, 32)
	typedDatasets := make([][]int, 32)
	for i := 0; i < len(datasets); i++ {
		dataset := make([]interface{}, 100)
		typedDataset := make([]int, 100)
		for j := 0; j < len(dataset); j++ {
			val := int(0)
			fuzzer.Fuzz(&val)
			dataset[j] = val
			typedDataset[j] = val
		}
		datasets[i] = dataset
		typedDatasets[i] = typedDataset
	}
	DoMax(datasets[0])
	b.Run("plz", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			DoMax(datasets[i%32])
		}
	})
	b.Run("typed", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			max_int_typed(typedDatasets[i%32])
		}
	})
}