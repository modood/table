GO = go
GO_GET = $(GO) get
GO_TEST = $(GO) test -v


all: dep test

dep:
	-$(GO_GET) github.com/smartystreets/goconvey/convey

test:
	$(GO_TEST)

.PHONY: all dep test

