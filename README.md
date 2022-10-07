# go-enumerable

This is a project mostly for learning to implement a enumerable library in Go using generics.
It is inspired by [Enumerable](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable) in .NET though it does not attempt to replicate the linq API.
There has not been analysis on performance but efforts have been made to follow best practices where applicable.
I am sure that there have been many similar libraries made like this but I avoided searching them out in order to see what a naive implementation could look like.

## Future Features

- [x] Lazy evaluation
- [x] Parallel evaluation
- [ ] More functions
- [ ] Performance improvements and analysis

## Stretch Goals

- [ ] DSL
- [ ] Query builder for SQL and other databases
- [ ] Large data set support
