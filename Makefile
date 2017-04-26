usage:
	@echo
	@echo "Command   : Description"
	@echo "--------- : ----------------"
	@echo "make venv : Create virtual environment"
	@echo

venv:
	@virtualenv -p python3 venv
	@source venv/bin/activate && python --version
	@source venv/bin/activate && pip install -r requirements.txt
