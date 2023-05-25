package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-tag-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-tag-reads-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-product-tag-reads-rmq-kube/config"
	"sync"

	database "github.com/latonaio/golang-mongodb-network-connector"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type DPFMAPICaller struct {
	ctx     context.Context
	conf    *config.Conf
	rmq     *rabbitmq.RabbitmqClient
	mongodb *database.MongoDB
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient, mongodb *database.MongoDB,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:     context.Background(),
		conf:    conf,
		rmq:     rmq,
		mongodb: mongodb,
	}
}

func (c *DPFMAPICaller) AsyncProductTagReads(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	log *logger.Logger,
) (interface{}, []error) {
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	var response interface{}

	response = c.readProcess(nil, &mtx, input, output, accepter, &errs, log)

	return response, nil
}

func checkResult(msg rabbitmq.RabbitmqMessage) bool {
	data := msg.Data()
	d, ok := data["result"]
	if !ok {
		return false
	}
	result, ok := d.(string)
	if !ok {
		return false
	}
	return result == "success"
}

func getBoolPtr(b bool) *bool {
	return &b
}
