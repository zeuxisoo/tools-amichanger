usage:
	@echo
	@echo "Command            : Description"
	@echo "------------------ : -----------------"
	@echo "make ami           : Create and build the ami files"
	@echo "make build         : Build the main program"
	@echo "------------------ : -----------------"
	@echo "make clean-results : Clean the generated *.bin files in results directory"
	@echo

ami: clean-ami
	@cd changer && git clone https://github.com/socram8888/amiitool.git
	@cd changer/amiitool/mbedtls && git submodule init && git submodule update --recursive
	@cd changer/amiitool && make
	@cd changer/amiitool && \
		ar q libamiibo.a amiibo.o && \
		ar q libkeygen.a keygen.o && \
		ar q libdrbg.a drbg.o

clean-ami:
	@rm -rf changer/amiitool

clean-results:
	@rm -rf results/*.bin

build:
	@go build *.go
