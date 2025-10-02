package middleware

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bufbuild/connect-go"
)

// LoggingInterceptor は Connect v1 用
func LoggingInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {

			reqJSON, _ := json.Marshal(req.Any())
			fmt.Printf("[RPC] Request: %s\n", reqJSON)

			res, err := next(ctx, req)
			if err != nil {
				fmt.Printf("[RPC] Error: %v\n", err)
				return nil, err
			}

			resJSON, _ := json.Marshal(res.Any())
			fmt.Printf("[RPC] Response: %s\n", resJSON)

			return res, nil
		}
	}
}
