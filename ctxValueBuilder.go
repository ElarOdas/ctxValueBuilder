package ctxValueBuilder

import (
	"context"
)

func FromContext[K comparable, R any](err error) func(ctx context.Context, key K) (R, error) {
	return func(ctx context.Context, key K) (R, error) {

		value := ctx.Value(key)
		switch typedValue := value.(type) {
		case R:
			return typedValue, nil
		default:
			var emptyVal R
			return emptyVal, err
		}
	}
}

func ToContext[K comparable, V any]() func(ctx context.Context, key K, value V) context.Context {
	return func(ctx context.Context, key K, value V) context.Context {
		return context.WithValue(ctx, key, value)
	}
}
