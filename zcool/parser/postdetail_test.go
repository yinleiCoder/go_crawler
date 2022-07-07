package parser

import (
	"goCrawler/model"
	"io/ioutil"
	"testing"
)

func TestParsePostDetail(t *testing.T) {
	contents, err := ioutil.ReadFile("postdetail_test_data.html")
	if err != nil {
		//panic(err)
	}
	result := ParsePostDetail(contents)
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}
	post := result.Items[0].(model.Post)
	expected := model.Post{
		Name: "儒家与巴豆",
		Home: "https://www.zcool.com.cn/u/17829922",
		Imgs: []string{"https://img.zcool.cn/community/0162e26221e07811013e8cd0e8d491.jpg@1280w_1l_0o_100sh.jpg", "https://img.zcool.cn/community/014bed6221e08111013f01cd4cdedd.png@1280w_1l_2o_100sh.png"},
	}
	t.Logf("expected %v, but was %v", expected, post)
	//if expected != post {
	//	t.Errorf("expected %v; but was %v", expected, post)
	//}
}
