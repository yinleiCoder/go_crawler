package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"goCrawler/model"
	"testing"
)

func TestSaveToElkByDocker(t *testing.T) {
	expected := model.Post{
		Name: "《STEEL  GANGSTER》",
		Home: "https://www.zcool.com.cn/work/ZNTg5NDI0MDg=.html",
		Imgs: []string{"https://img.zco.cn/community/01u6vsh8xgcaz7fs8a8ibj3734.jpg?x-oss-process=image/auto-orient,1/resize,m_lfit,w_1280,limit_1/sharpen,100/format,webp/quality,Q_100",
				"https://img.zcool.cn/community/01w83dpybcaec5v8mi08wn3832.jpg?x-oss-process=image/auto-orient,1/resize,m_lfit,w_1280,limit_1/sharpen,100/format,webp/quality,Q_100",
				"https://img.zcool.cn/community/01kfzf0msmu83rxrovw4kp3132.jpg?x-oss-process=image/auto-orient,1/resize,m_lfit,w_1280,limit_1/sharpen,100/format,webp/quality,Q_100",
				"https://img.zcool.cn/community/018yt2hpb4guant6my7nos3936.jpg?x-oss-process=image/auto-orient,1/resize,m_lfit,w_1280,limit_1/sharpen,100/format,webp/quality,Q_100",
				"https://img.zcool.cn/community/01gqqehvn7vnwomwkqpiv13133.jpg?x-oss-process=image/auto-orient,1/resize,m_lfit,w_1280,limit_1/sharpen,100/format,webp/quality,Q_100"}}

	id, err := SaveToElkByDocker(expected)
	if err != nil {
		//
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {

	}
	resp, err :=client.Get().Index("golang_spa").
		Type("zcool").
		Id(id).Do(context.Background())
	if err != nil {

	}
	t.Logf("%s", resp.Source)
	var actual model.Post
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {

	}
	//if expected != actual {
	//	t.Errorf("c")
	//}
}