package xx

import (
	"context"
	"reflect"
	"strings"
)

type injectKey string

func InjectKey[T any](v T, keys ...string) any {
	t := reflect.TypeOf(&v)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if len(keys) == 0 {
		return injectKey(t.PkgPath() + "::" + t.Name())
	} else {
		return injectKey(t.PkgPath() + "::" + t.Name() + "::" + strings.Join(keys, "::"))
	}
}

func Inject[T any](ctx context.Context, v T, keys ...string) context.Context {
	return context.WithValue(ctx, InjectKey(&v, keys...), v)
}

func Extract[T any](ctx context.Context, keys ...string) (out T, ok bool) {
	out, ok = ctx.Value(InjectKey(&out, keys...)).(T)
	return
}
