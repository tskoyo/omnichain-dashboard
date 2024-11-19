package services

import (
	"github.com/go-resty/resty/v2"
)

func FetchLayerZeroData(endpoint string) (interface{}, error) {
	client := resty.New()
	response, err := client.R().Get(endpoint)
	if err != nil {
		return nil, err
	}
	return response.Body(), nil
}
