export GOPATH=$(shell pwd):$(shell pwd)/vendors

all: alidns

alidns:
	cd src && go build -gcflags "-N -l" -o ../$@

clean:
	rm -fr alidns

-include .deps

dep:
	echo -n "alidns:" > .deps
	find src/ -name '*.go' | awk '{print $$0 " \\"}' >> .deps
	echo "" >> .deps
