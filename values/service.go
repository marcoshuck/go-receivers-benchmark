package values

import (
	"github.com/marcoshuck/go-receivers-benchmark/logger"
)

type Service interface {
	Do()
}

type service struct {
	Logger logger.Logger
}

func (s service) Do() {
	s.Logger.Log("Done.")
}

func NewService(logger logger.Logger) Service {
	return &service{
		Logger: logger,
	}
}
