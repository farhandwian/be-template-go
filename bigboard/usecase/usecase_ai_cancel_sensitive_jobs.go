package usecase

import (
	"bigboard/gateway"
	bigboardModel "bigboard/model"
	"context"
	"fmt"
	"shared/core"
	sharedGateway "shared/gateway"
)

type AiCancelSensitiveJobsReq struct {
	IDSensitveJob string `json:"id_sensitive_job"`
}

type AiCancelSensitiveJobsRes struct{}

type AiAiCancelSensitiveJobsUseCase = core.ActionHandler[AiCancelSensitiveJobsReq, AiCancelSensitiveJobsRes]

func ImplAiCancelSensitiveJobs(
	getSensitiveJob gateway.GetListSensitiveJobGateway,
	updateSensitiveJob gateway.SensitiveJobsSave,
	sendSSEGateway sharedGateway.SendSSEMessage,
) AiAiCancelSensitiveJobsUseCase {
	return func(ctx context.Context, req AiCancelSensitiveJobsReq) (*AiCancelSensitiveJobsRes, error) {
		sensitiveJob, err := getSensitiveJob(ctx, gateway.GetListSensitiveJobReq{ID: req.IDSensitveJob})
		if err != nil {
			return nil, err
		}

		if sensitiveJob.Status == bigboardModel.StatusFailed || sensitiveJob.Status == bigboardModel.StatusSuccess {
			return nil, fmt.Errorf("sensitive job with id %v is already in %v status", req.IDSensitveJob, sensitiveJob.Status)
		}
		sensitiveJob.Status = bigboardModel.StatusFailed
		if _, err := updateSensitiveJob(ctx, gateway.SensitiveJobsSaveReq{SensitiveJobs: *sensitiveJob}); err != nil {
			return nil, err
		}

		_, _ = sendSSEGateway(ctx, sharedGateway.SendSSEMessageReq{
			Subject:      "cancel-sensitive-jobs",
			FunctionName: "cancelSensitiveJob",
			Data:         fmt.Sprintf("Sensitive job with id %v has been canceled", req.IDSensitveJob),
		})
		return &AiCancelSensitiveJobsRes{}, nil
	}
}
