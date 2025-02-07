package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"shared/core"
)

func Logging[R any, S any](actionHandler core.ActionHandler[R, S], indentation int) core.ActionHandler[R, S] {

	return func(ctx context.Context, request R) (*S, error) {

		{
			bytes, err := json.Marshal(request)
			if err != nil {
				return nil, err
			}

			PrintWithIndentation(fmt.Sprintf(">>> REQUEST          %s\n", string(bytes)), indentation)
		}

		response, err := actionHandler(ctx, request)
		if err != nil {
			PrintWithIndentation(fmt.Sprintf(">>> RESPONSE ERROR  %s\n\n", err.Error()), indentation)
			PrintLine(indentation)
			return nil, err
		}

		{
			bytes, err := json.Marshal(response)
			if err != nil {
				return nil, err
			}

			PrintWithIndentation(fmt.Sprintf(">>> RESPONSE SUCCESS %s\n\n", string(bytes)), indentation)
		}

		PrintLine(indentation)

		return response, nil
	}
}

func PrintLine(indentation int) {
	if indentation == 0 {
		// fmt.Printf("----------------------------------------------------------------------------------------------------------\n")
	}
}

func PrintWithIndentation(message string, indentation int) {
	// indentStr := strings.Repeat(" ", indentation)
	// fmt.Printf("%s%s", indentStr, message)
}
