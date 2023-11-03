CURRENT_DIR = $(dir $(abspath $(firstword $(MAKEFILE_LIST))))

all: generate-crds

CONTROLLER_TOOLS_VERSION ?= v0.13.0
CONTROLLER_GEN ?= bin/controller-gen-${CONTROLLER_TOOLS_VERSION}
${CONTROLLER_GEN}:
	GOBIN=$(CURRENT_DIR)/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)
	mv ./bin/controller-gen ${CONTROLLER_GEN}

.PHONY: generate-crds
generate-crds: ${CONTROLLER_GEN}
	$(CONTROLLER_GEN) paths=./... crd:crdVersions=v1 output:artifacts:config=./crds
