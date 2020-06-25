#!/usr/bin/env bash
printf "\nRegenerating gqlgen files\n"
rm -f graph/generated/generated.go \
    graph/models/models.go \
    graph/resolvers/resolver.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"