usage:
	@echo
	@echo "Command    : Description"
	@echo "---------- : ----------------"
	@echo "make venv  : Create virtual environment"
	@echo "make tools : Generate related tools"
	@echo "make clean : Clean the environment, generated tools and decrypted file"
	@echo

venv:
	@virtualenv -p python3 venv
	@source venv/bin/activate && python --version
	@source venv/bin/activate && pip install -r requirements.txt

tools: clean-tools
	@bash ./scripts/amiitool.sh install

clean-tools:
	@bash ./scripts/amiitool.sh uninstall

clean: clean-tools
	@rm -rf venv
	@rm -rf results/*.bin
