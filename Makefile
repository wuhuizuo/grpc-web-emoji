PROTO_FILES 		:= $(shell find . -regextype posix-extended -regex '.*\.proto')

.PHONY: default
default: gen_grpc

.PHONY: gen_pre
gen_pre:
	@chmod +x ./.build/*.sh

.PHONY: gen_grpc
gen_grpc: gen_pre
	@echo "start $@"
	@for f in $(PROTO_FILES); do \
		./.build/gen_grpc.sh $${f}; \
	done
	@echo "$@ success"

.PHONY: install_tools
install_tools:
	@chmod +x ./.build/install_tools.sh
	./.build/install_tools.sh

.PHONY: clean_gen
clean_gen:
	# pb go files
	@find . -name *.pb.go   			-exec rm -vf {} \;
	@find . -name *.pb.*.go   			-exec rm -vf {} \;
	# go mock files
	@find . -name *_mock.go   			-exec rm -vf {} \;
	# api doc scheme json
	@find . -name apidocs.swagger.json 	-exec rm -vf {} \;