package main

import (
	"context"
	"fmt"
	"reflect"

	"golang.org/x/sync/errgroup"
)

type fetchFunc[T any] func(context.Context) (T, error)

func fetch[T any, PV *T](ctx context.Context, x PV, f fetchFunc[T]) func() error {
	return func() (err error) { *x, err = f(ctx); return err }
}

// NodeAB  -> NodeC -> NodeDEF -> NodeG
// A, B    -> C     -> D, E, F -> G
type (
	A   string
	B   string
	C   string
	D   string
	E   string
	F   string
	G   string
	AB  string
	DEF string
)

func computeFor[OutT, InT ~string](in InT) (o OutT, err error) {
	return OutT(fmt.Sprintf("%s(%s)", reflect.TypeOf(o).Name(), in)), nil
}

func ComputeA(context.Context) (A, error)           { return "A", nil }
func ComputeB(context.Context) (B, error)           { return "B", nil }
func ComputeC(_ context.Context, in AB) (C, error)  { return computeFor[C](in) }
func ComputeD(_ context.Context, in C) (D, error)   { return computeFor[D](in) }
func ComputeE(_ context.Context, in C) (E, error)   { return computeFor[E](in) }
func ComputeF(_ context.Context, in C) (F, error)   { return computeFor[F](in) }
func ComputeG(_ context.Context, in DEF) (G, error) { return computeFor[G](in) }

func ComputeAB(ctx context.Context) (AB, error) {
	var (
		a A
		b B
	)
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(fetch(ctx, &a, ComputeA))
	eg.Go(fetch(ctx, &b, ComputeB))
	if err := eg.Wait(); err != nil {
		return "", err
	}
	return computeFor[AB](string(a) + string(b))
}

func ComputeDEF(ctx context.Context, in C) (DEF, error) {
	var (
		d D
		e E
		f F
	)
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(fetch(ctx, &d, func(ctx context.Context) (D, error) { return ComputeD(ctx, in) }))
	eg.Go(fetch(ctx, &e, func(ctx context.Context) (E, error) { return ComputeE(ctx, in) }))
	eg.Go(fetch(ctx, &f, func(ctx context.Context) (F, error) { return ComputeF(ctx, in) }))
	if err := eg.Wait(); err != nil {
		return "", err
	}
	return computeFor[DEF](string(d) + string(e) + string(f))
}

func main() {
	ctx := context.Background()
	fmt.Println(InitG(ctx))
}
