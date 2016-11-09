DOCKER_RUN := docker run --rm -it -v $$PWD:/go/src/page-elements
SUPPORTED_ASSET_DIRS := templates
BUILDLESS_CONTAINERS := consul traefik

.PHONY: vendor specs

up: up/consul up/traefik wait up/site up/page-home-v1 up/header up/footer \
	up/recommendations up/navigation up/skeleton up/page-sub up/dashboard \
	up/page-home-v2

up/%: build/% stop/%
	docker-compose up -d $*

stop/%:
	docker-compose stop $*

build/%: ASSET_DIRS = $(addsuffix /..., $(filter $(SUPPORTED_ASSET_DIRS), $(subst $*/,, $(wildcard $*/*))))
build/%:
	# run go-bindata
	$(if $(ASSET_DIRS), $(DOCKER_RUN) -w /go/src/page-elements/$* dominikznidar/go-bindata $(ASSET_DIRS))
	# build it
	$(DOCKER_RUN) -w /go/src/page-elements/$* golang:1.7.1-alpine go build -o ../bin/micro-$*

build/consul: ;
build/traefik: ;

logs:
	docker-compose logs -f --tail=0

# update local vendor folder
vendor:
	$(DOCKER_RUN) -w /go/src/page-elements trifs/govendor:latest fetch +missing
	$(DOCKER_RUN) -w /go/src/page-elements trifs/govendor:latest remove +unused

specs:
	@$(MAKE) -C specs

clean:
	rm bin/*
	docker-compose down

wait:
	@echo "Waiting a while for consul to wake up ..."
	sleep 10
