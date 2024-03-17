# OpenSlides Go

This repo contains go packages, that are potentially used by more then one OpenSlides repo.


## Update models.yml

To use a new models.yml update the meta repository in `meta`.
Afterwards call `go generate ./...` to update the generated files.
