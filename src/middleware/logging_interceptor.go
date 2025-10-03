package middleware

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/fatih/color"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func LoggingInterceptor() connect.UnaryInterceptorFunc {
	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}

	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {

			color.New(color.FgYellow).Printf("[RPC] Procedure: %s\n", req.Spec().Procedure)

			if reqMsg, ok := req.Any().(proto.Message); ok {
				if reqJSON, err := marshaler.Marshal(reqMsg); err == nil {
					color.New(color.FgCyan).Printf("[RPC] Request: %s\n", reqJSON)
				}
			}

			res, err := next(ctx, req)
			if err != nil {
				color.New(color.FgRed).Printf("[RPC] Error: %v\n", err)
				return nil, err
			}

			if resMsg, ok := res.Any().(proto.Message); ok {
				if resJSON, err := marshaler.Marshal(resMsg); err == nil {
					color.New(color.FgGreen).Printf("[RPC] Response: %s\n", resJSON)
				}
			}

			return res, nil
		}
	}
}
