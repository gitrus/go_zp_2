package multiplyr

import (
	"errors"
)

type MultiplyArgs struct {
	A, B float64
}

type MultiplyService struct{}

func NewMultiplyService() *MultiplyService {
	return &MultiplyService{}
}

func (m *MultiplyService) Multiply(args *MultiplyArgs, result *float64) error {
	if args == nil {
		return errors.New("nil arguments provided")
	}

	*result = args.A * args.B
	return nil
}
