usage:
	@echo
	@echo "Command  : Description"
	@echo "-------- : -----------------"
	@echo "make ami : Create and build the ami files"
	@echo

ami: clean-ami
	git clone https://github.com/socram8888/amiitool.git
	cd amiitool/mbedtls && git submodule init && git submodule update --recursive
	cd amiitool && make
	cd amiitool && ar q libamiibo.a amiibo.o && ar q libkeygen.a keygen.o && ar q libdrbg.a drbg.o

clean-ami:
	rm -rf amiitool
