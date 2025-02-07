package gateway

import (
	"context"
	"shared/core"
	"time"

	"math/rand"
)

type GenerateRandomReq struct {
	N int
}

type GenerateRandomRes struct {
	Random string
}

type GenerateRandom = core.ActionHandler[GenerateRandomReq, GenerateRandomRes]

func ImplGenerateRandom() GenerateRandom {
	return func(ctx context.Context, req GenerateRandomReq) (*GenerateRandomRes, error) {

		// Seed the random number generator
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		const digits = "0123456789"
		result := make([]byte, req.N)
		for i := 0; i < req.N; i++ {
			if i == 0 {
				result[i] = digits[1+r.Intn(9)] // Ensures first digit is not 0
			} else {
				result[i] = digits[r.Intn(10)]
			}
		}

		return &GenerateRandomRes{Random: string(result)}, nil

	}
}

func ImplGenerateRandomButStaticForMock(value string) GenerateRandom {
	return func(ctx context.Context, request GenerateRandomReq) (*GenerateRandomRes, error) {
		return &GenerateRandomRes{Random: value}, nil
	}
}
