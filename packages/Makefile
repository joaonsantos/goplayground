# go option
PKGS       := ./pkg/...
TESTFLAGS  :=
GOFLAGS    :=

Q=@

# ------------------------------------------------------------------------------
#  run
.PHONY: run
run:
	$Qgo run ./cmd/main.go


# ------------------------------------------------------------------------------
#  test
.PHONY: test
test: test-unit

.PHONY: test-unit
test-unit:
	@echo
	@echo "==> Running unit tests <=="
	@echo
	$Qgo test $(GOFLAGS) $(PKGS) $(TESTFLAGS)
