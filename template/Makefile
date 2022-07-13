TOP = {{ .Top }}
SBT = sbt --client

SRC = $(shell find ./src/main/scala -name "*.scala")
TEST = $(shell find ./src/test/scala -name "*.scala")

build: generated/$(TOP).v

check: $(SRC) $(TEST)
	$(SBT) compile

generated:
	mkdir -p generated

generated/$(TOP).v: $(SRC) generated
	$(SBT) run

test: $(SRC) $(TEST)
	$(SBT) "Test / test"

clean:
	rm -rf generated

.PHONY: build check test clean
