package delivery

import (
	"assistantor/repository"
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type DispatchService struct {
}

func init() {
	hystrix.ConfigureCommand("query_delivery_info", hystrix.CommandConfig{
		Timeout:                int(3 * time.Second),
		MaxConcurrentRequests:  10,
		SleepWindow:            5000,
		RequestVolumeThreshold: 20,
		ErrorPercentThreshold:  30,
	})
}

func (service DispatchService) QueryDeliveryInfo(ctx context.Context, parameter *QueryDeliverParameter) (*DeliveryResult, error) {
	_ = hystrix.Do("query_delivery_info", func() error {
		repository.GetDeliveryInfoByOrderId(parameter.OrderId)
		return nil
	}, func(err error) error {
		return nil
	})
	res := &DeliveryResult{
		OrderId:         "",
		UserId:          "",
		UserPhoneNumber: "",
		DeliveryAddress: "",
		DeliveryList:    nil,
	}
	return res, nil
}

func StartDeliverServer(port int) {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Error().Msgf("start grpc listener error: %s", err.Error())
		return
	}
	s := grpc.NewServer()
	d := DispatchService{}
	RegisterQueryServer(s, &d)
	reflection.Register(s)
	log.Info().Msgf("delivery system serve at %v", addr)
	err = s.Serve(listener)
	if err != nil {
		log.Error().Msgf("start grpc server error: %s", err.Error())
		return
	}
	return
}
