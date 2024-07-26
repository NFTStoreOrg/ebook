package models

import (
	"context"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/result"
)

func CreateESIndex(collection string) (message *esapi.Response, err error) {
	mapping := `{
        "settings": {
            "analysis": {
                "analyzer": {
                    "ik_max_word": {
                        "type": "ik_max_word"
                    }
                }
            }
        },
        "mappings": {
            "properties": {
                "title": {
                    "type": "text",
                    "analyzer": "ik_max_word"
                }
            }
        }
    }`
	req := esapi.IndicesCreateRequest{
		Index: collection,
		Body:  strings.NewReader(mapping),
	}
	res, err := req.Do(context.Background(), EsClient)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreateESDocument(info Book) (result.Result, error) {
	resp, err := EsClient.Index(info.Class.ClassName).Id(strconv.FormatInt(info.TokenId, 10)).
		Document(info).Do(context.Background())
	if err != nil {
		return result.Result{}, err
	}

	result := resp.Result
	return result, nil
}

func UpdateESDocument(info Book) error {
	_, err := EsClient.Update(info.Class.ClassName, "1").Doc(info).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
