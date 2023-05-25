package dpfm_api_caller

import (
	"bytes"
	"context"
	"data-platform-api-product-tag-reads-rmq-kube/DPFM_API_Caller/requests"
	dpfm_api_input_reader "data-platform-api-product-tag-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-tag-reads-rmq-kube/DPFM_API_Output_Formatter"
	"encoding/json"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
	"io/ioutil"
	"net/http"
	"sync"
)

func (c *DPFMAPICaller) readProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var productTag *[]dpfm_api_output_formatter.ProductTag
	for _, fn := range accepter {
		switch fn {
		case "ProductTag":
			func() {
				productTag = c.ProductTag(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		ProductTag: productTag,
	}

	return data
}

func (c *DPFMAPICaller) ProductTag(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductTag {
	product := input.ProductTag.Product
	size := 10

	query := requests.Query{
		Match: requests.Match{
			Product: product,
		},
	}

	aggs := requests.Aggs{
		DuplicateAggs: requests.DuplicateAggs{
			Terms: requests.Terms{
				Field: "ProductTag.keyword",
				Size:  size,
			},
		},
	}

	descendingOrderQuery := &requests.DescendingOrderQuery{
		Query: query,
		Aggs:  aggs,
	}

	body, _ := json.Marshal(descendingOrderQuery)

	resp, err := http.Post(
		fmt.Sprintf(
			"%s/dataplatformmongodbkube.data_platform_product_tag_data/_search?pretty",
			c.conf.ES.URL,
		),
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		*errs = append(*errs, xerrors.New(string(body)))
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToProductTag(input, resp)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
