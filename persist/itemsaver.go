package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"goCrawler/model"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func ItemSaver(elkIndex string, elkType string) (chan interface{}, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			log.Printf("(item saver)got item: %v", item)
			//save(item)
			_, err := SaveToElkByDocker(client, elkIndex, elkType, item)
			if err != nil {
				log.Printf("Item saver: error saving tiem %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

/**
save data by ElasticSearch.
there is elasticsearch which created by docker.
 */
func SaveToElkByDocker(client *elastic.Client, elkIndex string, elkType string, item interface{}) (id string, err error) {
	resp, err := client.Index().
		Index(elkIndex).
		Type(elkType).BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}

/**
Abandoned API: save imgs to your local.
 */
func save(item interface{}) {
	post := item.(model.Post)
	for index, imgUrl := range post.Imgs {
		resp, err := http.Get(imgUrl)
		if err != nil {
			log.Printf("(http download error): %s", err)
		}
		defer resp.Body.Close()
		imgBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("(ioutil read error): %s", err)
		}
		err = os.WriteFile("imgs/"+post.Name+strconv.Itoa(index)+".jpg", imgBytes, 0666)
		if err != nil {
			log.Printf("(ioutil write error): %s", err)
		}
	}
}
