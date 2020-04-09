# allow optional per-repo overrides
-include Makefile.overrides.mk

RUN = echo +

%:
	@$(RUN) + $@

default:
	@$(RUN)

.PHONY: default

test:
	@$(RUN)
