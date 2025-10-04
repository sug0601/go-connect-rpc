package middleware

import (
	"context"
	"fmt"
	"log"

	"buf.build/go/protovalidate"
	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/proto"
)

func ValidationInterceptor() connect.UnaryInterceptorFunc {
	validator, err := protovalidate.New()
	if err != nil {
		log.Fatalf("failed to create validator: %v", err)
	}

	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if msg, ok := req.Any().(proto.Message); ok {
				if err := validator.Validate(msg); err != nil {
					log.Printf("Validation failed for %s: %v", req.Spec().Procedure, err)
					return nil, connect.NewError(connect.CodeInvalidArgument,
						fmt.Errorf("validation failed: %w", err))
				}
			}

			return next(ctx, req)
		}
	}
}
