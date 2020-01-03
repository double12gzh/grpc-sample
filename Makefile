
.PHONY: gen-proto
gen-proto:
	rm -rf gen
	mkdir gen
	protoc -I proto/ example.proto --go_out=plugins=grpc:gen

