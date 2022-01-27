package services

import (
	"app/model"
	"app/repository/rest/leslie_api_repo"
	"fmt"
	"github.com/golanshy/plime_core-go/logger"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"time"
)

type Service interface {
	LoadTest(data_models.LoadTestRequest) *rest_errors.RestErr
	performTest(string, *data_models.Location, *data_models.Order) *rest_errors.RestErr
}

type service struct {
	repository leslie_api_repo.RestLeslieApiRepository
}

func NewService(repo leslie_api_repo.RestLeslieApiRepository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) LoadTest(request data_models.LoadTestRequest) *rest_errors.RestErr {
	if request.Repetitions <= 0 {
		return rest_errors.NewBadRequestError("Unknown Repetitions")
	}

	if request.DelayInMilliseconds < 0 {
		return rest_errors.NewBadRequestError("Unknown delay in milliseconds")
	}

	location := data_models.NewLocation()

	for i := 0; i < request.Repetitions; i++ {
		logger.Info(fmt.Sprintf("Perform Test number %d", i+1))

		order := data_models.NewOrder(i + 1)

		s.performTest(request.TestType, location, order)

		if request.DelayInMilliseconds > 0 {
			time.Sleep(time.Duration(request.DelayInMilliseconds) * time.Millisecond)
		}
	}

	return nil
}

func (s *service) performTest(testType string, location *data_models.Location, order *data_models.Order) *rest_errors.RestErr {

	switch testType {

	case "GetEntireMenu":
		fmt.Println("Get Entire Menu")
		go s.repository.GetEntireMenus(location)
		return nil

	case "Order":
		fmt.Println("Order")
		go s.repository.Order(order)
		return nil

	case "GetEntireMenuAndOrder":
		fmt.Println("Get Entire Menu")
		go s.repository.GetEntireMenus(location)
		fmt.Println("Order")
		go s.repository.Order(order)
		return nil
	}

	return rest_errors.NewBadRequestError("unknown test type")
}
