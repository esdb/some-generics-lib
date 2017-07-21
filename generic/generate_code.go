package generic

import (
	"bytes"
	"io/ioutil"
)

var expandedFuncs = map[string]interface{}{}

func GenerateCode(gopath string, pkgPath string) {
	state.out = bytes.NewBuffer(nil)
	state.importPackages = map[string]bool{
		"github.com/v2pro/wombat/generic": true,
	}
	state.declarations = map[string]bool{}
	state.expandedFuncNames = map[string]bool{}
	state.pkgPath = pkgPath
	prelog := []byte(`
package model
	`)
	for importPackage := range state.importPackages {
		prelog = append(prelog, '\n')
		prelog = append(prelog, `import "`...)
		prelog = append(prelog, importPackage...)
		prelog = append(prelog, '"')
	}
	prelog = append(prelog, "\nfunc init() {"...)
	for _, funcDeclaration := range funcDeclarations {
		expandedFuncName, err := funcDeclaration.funcTemplate.expand(funcDeclaration.templateArgs)
		if err != nil {
			panic(err.Error())
		}
		prelog = append(prelog, '\n')
		prelog = append(prelog, `generic.RegisterExpandedFunc("`...)
		prelog = append(prelog, expandedFuncName...)
		prelog = append(prelog, `",`...)
		prelog = append(prelog, expandedFuncName...)
		prelog = append(prelog, ')')
	}
	prelog = append(prelog, '}')
	for declaration := range state.declarations {
		prelog = append(prelog, '\n')
		prelog = append(prelog, declaration...)
	}
	source := append([]byte(prelog), state.out.Bytes()...)
	err := ioutil.WriteFile(gopath+"/src/"+pkgPath+"/generated.go", source, 0666)
	if err != nil {
		panic(err.Error())
	}
}

func RegisterExpandedFunc(expandedFuncName string, expandedFunc interface{}) {
	expandedFuncs[expandedFuncName] = expandedFunc
}