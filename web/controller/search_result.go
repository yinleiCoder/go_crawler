package controller

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"goCrawler/model"
	"goCrawler/web/pagemodel"
	"goCrawler/web/view"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Printf("elastic create client error: %v", err)
	}
	return SearchResultHandler{
		view: view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	fmt.Println(q, from)
	if err != nil {
		from = 0
	}
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (pagemodel.SearchResult, error) {
	var result pagemodel.SearchResult
	result.Query = q
	var resp *elastic.SearchResult
	if q != "" {
		resp, _ = h.client.
			Search("golang_spa").
			Query(elastic.NewQueryStringQuery(q)).
			From(from).
			Do(context.Background())
	} else {
		// query all data
		resp, _ = h.client.Search("golang_spa").From(from).Do(context.Background())
	}

	result.Hits = int(resp.TotalHits())
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(model.Post{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}
