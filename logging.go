package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next IPriceFetcherService
}

func NewLoggingService(next IPriceFetcherService) IPriceFetcherService {
	return &loggingService{
		next: next,
	}
}

func (l *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"ID":    ctx.Value("requestID"),
			"took":  time.Since(begin),
			"err":   err,
			"price": price,
		}).Info("Fetch price")
	}(time.Now())

	return l.next.FetchPrice(ctx, ticker)
}
