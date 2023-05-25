package dpfm_api_input_reader

import (
	"data-platform-api-product-tag-reads-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToProductTag() *requests.ProductTag {
	data := sdc.ProductTag
	return &requests.ProductTag{
		ProductTag: data.ProductTag,
	}
}
