package models

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

var EsClient *elasticsearch.TypedClient

func init() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	var err error
	EsClient, err = elasticsearch.NewTypedClient(cfg)
	if err != nil {
		fmt.Println(err)
	}
}
