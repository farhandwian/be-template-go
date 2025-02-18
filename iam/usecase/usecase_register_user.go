package usecase

import (
	"context"
	"errors"
	"fmt"
	"iam/gateway"
	"iam/model"
	"shared/constant"
	"shared/core"
	sharedGateway "shared/gateway"
	sharedModel "shared/model"
	"strings"
	"time"

	"github.com/google/uuid"
)

type RegisterUserReq struct {
	Email       model.Email
	PhoneNumber model.PhoneNumber
	Name        string
	Now         time.Time
}

type RegisterUserRes struct {
	UserID model.UserID `json:"user_id"`
}

type RegisterUser = core.ActionHandler[RegisterUserReq, RegisterUserRes]

func ImplRegisterUser(

	generateId gateway.GenerateId,
	saveUser gateway.UserSave,
	userGetAll gateway.UserGetAll,
	createActivityMonitoring sharedGateway.CreateActivityMonitoringGateway,
) RegisterUser {
	return func(ctx context.Context, request RegisterUserReq) (*RegisterUserRes, error) {

		if err := request.Validate(); err != nil {
			return nil, err
		}

		generatedIdObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		users, err := userGetAll(ctx, gateway.UserGetAllReq{Email: request.Email})
		if err != nil {
			return nil, err
		}

		if users.Count > 0 {
			return nil, fmt.Errorf("user dengan email %s sudah ada", request.Email)
		}

		userId := model.UserID(generatedIdObj.RandomId)

		newUser := model.NewUser(
			userId,
			request.Email,
			request.PhoneNumber,
			request.Name,
			request.Now,
		)

		if _, err := saveUser(ctx, gateway.UserSaveReq{User: newUser}); err != nil {
			return nil, err
		}

		//store logging
		_, err = createActivityMonitoring(ctx, sharedGateway.CreateActivityMonitoringReq{
			ActivityMonitor: sharedModel.ActivityMonitor{
				ID:           uuid.NewString(),
				UserName:     request.Name,
				Category:     constant.MONITORING_TYPE_IAM,
				ActivityTime: time.Now(),
				Description:  fmt.Sprintf("%s telah terdaftar ke dalam sistem", request.Name),
			},
		})
		if err != nil {
			return nil, err
		}

		return &RegisterUserRes{UserID: newUser.ID}, nil
	}
}

func (r RegisterUserReq) Validate() error {

	err := r.Email.Validate()
	if err != nil {
		return err
	}

	err = r.PhoneNumber.Validate()
	if err != nil {
		return err
	}

	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name must not empty")
	}

	return nil
}
