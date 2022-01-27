package leslie_api_repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/golanshy/load_tester/model"
	"github.com/golanshy/plime_core-go/logger"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
)

var (
	client *resty.Client
	apiUrl string
)

func init() {
	client = resty.New()
	apiUrl = "https://europe-west2-leslie-tech.cloudfunctions.net"
}

type RestLeslieApiRepository interface {
	GetEntireMenus(*data_models.Location)
	Order(*data_models.Order)
	post(path string, data interface{}) (*resty.Response, error)
}

type leslieApiRepository struct {
}

func NewRestLeslieApiRepository() RestLeslieApiRepository {
	return &leslieApiRepository{}
}

func (r *leslieApiRepository) GetEntireMenus(data *data_models.Location) {

	restResponse, err := r.post("getEntireMenu", data)

	if err != nil {
		logger.Error("error in api", err)
		return
	}

	if restResponse.StatusCode() > 299 {
		logger.Error("error getting menus from Leslie api", errors.New(restResponse.String()))
		return
	}

	logger.Info(fmt.Sprintf("StatusCode: %d, data: %s", restResponse.StatusCode(), restResponse.String()))
}

func (r *leslieApiRepository) Order(data *data_models.Order) {

	restResponse, err := r.post("order", data)

	if err != nil {
		logger.Error("error in api", err)
		return
	}

	if restResponse.StatusCode() > 299 {
		logger.Error("error creating order in Leslie api", errors.New(restResponse.String()))
		return
	}

	var results data_models.OrderResult
	if err := json.Unmarshal(restResponse.Body(), &results); err != nil {
		logger.Error("error unmarshaling json response when trying to get search search_results from companies house api", err)
		return nil, rest_errors.NewInternalServerError("error unmarshaling json response when trying to get search search_results from companies house api", err)
	}

	logger.Info(fmt.Sprintf("StatusCode: %d, data: %s", restResponse.StatusCode(), restResponse.String()))
}

func (r *leslieApiRepository) post(path string, data interface{}) (*resty.Response, error) {

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	params := make(map[string]string)

	return client.R().
		SetHeaders(headers).
		SetQueryParams(params).
		SetBody(data).
		Post(fmt.Sprintf("%s/%s", apiUrl, path))
}
