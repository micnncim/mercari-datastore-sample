.PHONY: proto-gen
proto-gen:
	protoc \
		-I=. \
		-I=${GOPATH}/src \
		-I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		--gogoslick_out=. \
		proto/*.proto
