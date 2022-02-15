An unsafe version that assumes enumerations always succeed.

DotNet LINQ may throws when it involves databases, multi-threading, await etc. In Go, database integration is not available yet. In most cases we just use `FromSlice`. So returning a `error` makes code unnecessarily complicated.