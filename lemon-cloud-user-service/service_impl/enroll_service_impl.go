package service_impl

import (
	"context"
	"lemon-cloud-user-common/dto"
)

type EnrollServiceImpl struct {
}

func (e EnrollServiceImpl) SignInByNumber(context.Context, *dto.SignInByNumberReq) (*dto.SignInRsp, error) {
	return nil, nil
}
