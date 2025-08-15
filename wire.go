//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
)

func InitG(ctx context.Context) (G, error) {
	panic(wire.Build(
		ComputeAB,
		ComputeC,
		ComputeDEF,
		ComputeG,
	))
}
