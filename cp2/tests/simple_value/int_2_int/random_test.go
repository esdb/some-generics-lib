package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/google/gofuzz"
	"github.com/v2pro/plz"
	"encoding/json"
	_ "github.com/v2pro/wombat/cp2"
)

func Test_random(t *testing.T) {
	should := require.New(t)
	fz := fuzz.New().MaxDepth(10).NilChance(0.3)
	for i := 0; i < 100; i++ {
		var src SrcType
		fz.Fuzz(&src)
		srcJson := toJson(src)
		var dst1 DstType
		should.Nil(plz.Copy(&dst1, src))
		var dst2 DstType
		fromJson(&dst2, srcJson)
		dst1Json := toJson(dst1)
		dst2Json := toJson(dst2)
		should.Equal(dst1Json, dst2Json)
	}
}

func toJson(obj interface{}) string {
	output, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	return string(output)
}

func fromJson(obj interface{}, encodedAsJson string) {
	err := json.Unmarshal([]byte(encodedAsJson), obj)
	if err != nil {
		panic(err.Error())
	}
}
