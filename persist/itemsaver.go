package persist

import (
	"goCrawler/model"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			log.Printf("(item saver)got item: %v", item)
			//save(item)
		}
	}()
	return out
}

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
