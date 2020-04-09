

PHONYS = echo OV
.PHONY: $(PHONYS)

$(PHONYS):
	@$(PHONYS) $@

.PHONY: init
init:
	@echo INIT