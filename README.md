# OpenSlides Go

[![ci](https://github.com/OpenSlides/openslides-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/OpenSlides/openslides-go/actions/workflows/test.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/OpenSlides/openslides-go)](https://pkg.go.dev/badge/github.com/OpenSlides/openslides-go)

This repository contains chaired packages for all openslides-go-services.


## Update meta

This repository contains all data from the meta repository. It should not be
necessary, to add the meta-subrepo to any other go repo. If meta changes, only
openslides-go has to be updated.

To Update openslides-go, update the sub repo and then call:

```
go generate ./...
```
