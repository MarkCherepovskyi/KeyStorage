package handlers

import (
	"context"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/data"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	containersCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}
func CtxContainerQ(entry data.ContainerQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, containersCtxKey, entry)
	}
}
func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func ContainerQ(r *http.Request) data.ContainerQ {
	return r.Context().Value(containersCtxKey).(data.ContainerQ)
}
