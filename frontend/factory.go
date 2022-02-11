package frontend

import (
	"fmt"
)

func NewFrontEnd(frontEndType string) (FrontEnd, error) {
	switch frontEndType {
	case "zero":
		return zeroFrontEnd{}, nil

	case "rest":
		return &restFrontEnd{}, nil

	case "grpc":
		return &grpcFrontEnd{}, nil

	default:
		return nil, fmt.Errorf("no such frontend %s", frontEndType)
	}
}
