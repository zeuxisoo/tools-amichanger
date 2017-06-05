usage:
	@echo
	@echo "Command             : Description"
	@echo "------------------- : -----------------"
	@echo "make amiitool       : Create and build the amiitool files"
	@echo "make build          : Build the main program"
	@echo "------------------- : -----------------"
	@echo "make clean-amiitool : Clean the amiitool"
	@echo "make clean-results  : Clean the generated *.bin files in results directory"
	@echo "make clean          : Clean all generated, downloaded files"
	@echo

amiitool: clean-amiitool
	@cd binding && git clone https://github.com/socram8888/amiitool.git
	@cd binding/amiitool && git reset --hard adda02952e77abdc59cb3f1cc0318d172efbe043
	@cd binding/amiitool/mbedtls && git submodule init && git submodule update --recursive
	@cd binding/amiitool && make
	@cd binding/amiitool && \
		ar q libamiibo.a amiibo.o && \
		ar q libkeygen.a keygen.o && \
		ar q libdrbg.a drbg.o

clean-amiitool:
	@rm -rf binding/amiitool

clean-results:
	@rm -rf results/*.bin

build:
	@go build *.go

clean: clean-amiitool clean-results
	@rm -rf main
