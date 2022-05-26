.PHONY: build-stable push-stable build-stable-ng push-stable-ng build push all

include Makefile.env

build-stable:
	docker build $$(echo ${TAGS_STABLE} | sed 's/[^ ]* */-t &/g') ${BUILD_ARGS_STABLE} .

build-stable-ng:
	docker build $$(echo ${TAGS_STABLE_NG} | sed 's/[^ ]* */-t &/g') ${BUILD_ARGS_STABLE_NG} .

build: build-stable build-stable-ng

push-stable:
	for TAG in $(TAGS_STABLE); do docker push $$TAG; done

push-stable-ng:
	for TAG in $(TAGS_STABLE_NG); do docker push $$TAG; done

push: push-stable push-stable-ng

all: build push
