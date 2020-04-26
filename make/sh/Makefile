args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

%:
	@:

test:
	@echo $(call args,defaultstring)
	@echo $(call args,defaultstring)