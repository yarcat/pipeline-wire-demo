# pipeline wire demo

This is a demonstration of compile-time type-safe pipeline implementation using google/wire.

## Libraries Used

### Core Dependencies

- **[github.com/google/wire v0.6.0](https://github.com/google/wire)** - Compile-time dependency injection framework for Go. Used to generate type-safe wiring code for the pipeline components at build time.

- **[golang.org/x/sync v0.16.0](https://golang.org/x/sync)** - Extended Go synchronization primitives. The `errgroup` package is used to run pipeline stages concurrently while properly handling errors and context cancellation.

### Development Tools

- **[github.com/google/wire/cmd/wire](https://github.com/google/wire)** - Wire code generation tool declared in `go.mod` as a tool dependency. Run with `go generate` to generate the dependency injection code.

## Build & Run

```bash
go generate
go run .

# Output:
# G(DEF(D(C(AB(A,B))),E(C(AB(A,B))),F(C(AB(A,B))))) <nil>
```
