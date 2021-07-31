.PHONY: build-stable push-stable build-beta push-beta build push all

include Makefile.env

build-stable:
	docker build $$(echo ${TAGS_STABLE} | sed 's/[^ ]* */-t &/g') ${BUILD_ARGS_STABLE} .

build-beta:
	docker build $$(echo ${TAGS_BETA} | sed 's/[^ ]* */-t &/g') ${BUILD_ARGS_BETA} .

build: build-stable build-beta

push-stable:
	for TAG in $(TAGS_STABLE); do docker push $$TAG; done

push-beta:
	for TAG in $(TAGS_BETA); do docker push $$TAG; done

push: push-stable push-beta

all: build push
