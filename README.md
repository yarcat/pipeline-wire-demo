# pipeline wire demo

This is a demonstration of compile-time type-safe pipeline implementation using google/wire.

## Build & Run

```bash
go generate
go run .

# Output:
# G(DEF(D(C(AB(AB)))E(C(AB(AB)))F(C(AB(AB))))) <nil>
```
