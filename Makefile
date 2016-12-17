all:
	@echo Hello

fmt:
	govendor fmt +local

test:
	govendor test +local