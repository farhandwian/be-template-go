package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
)

type GetPerizinanStatusReq struct {
}

type GetPerizinanStatusResp struct {
	ReportedCount   int `json:"reported_count"`
	UnreportedCount int `json:"unreported_count"`
}

type GetPerizinanStatusUseCase = core.ActionHandler[GetPerizinanStatusReq, GetPerizinanStatusResp]

func ImplGetPerizinanStatusUseCase(
	getPerizinanStatus gateway.GetLaporanPerizinanStatusCountGateway,
) GetPerizinanStatusUseCase {
	return func(ctx context.Context, req GetPerizinanStatusReq) (*GetPerizinanStatusResp, error) {
		perizinanStatus, err := getPerizinanStatus(ctx, gateway.GetLaporanPerizinanStatusCountReq{})
		if err != nil {
			return nil, err
		}
		return &GetPerizinanStatusResp{
			ReportedCount:   perizinanStatus.SubmittedCount,
			UnreportedCount: perizinanStatus.NotSubmittedCount,
		}, nil

	}
}
