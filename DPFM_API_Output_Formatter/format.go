package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-product-tag-reads-rmq-kube/DPFM_API_Input_Reader"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ConvertToProductTag(sdc *api_input_reader.SDC, resp *http.Response) (*[]ProductTag, error) {
	var productTag []ProductTag

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var esResponse ESResponse

	if err := json.Unmarshal(body, &esResponse); err != nil {
		return nil, err
	}

	for _, data := range esResponse.Aggregations.DuplicateAggs.Buckets {
		productTag = append(productTag, ProductTag{
			Key:      data.Key,
			DocCount: data.DocCount,
		})
	}

	return &productTag, nil
}
