package api

import "context"

type API struct {
	ctx context.Context
}

func New(ctx context.Context) *API {
	return &API{ctx}
}
