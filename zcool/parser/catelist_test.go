package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCateList(t *testing.T) {
	contents, err := ioutil.ReadFile("catelist_test_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", contents)
	result := ParseCateList(contents)
	const resultSize = 10
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d Items; but had %d", resultSize, len(result.Items))
	}
}
