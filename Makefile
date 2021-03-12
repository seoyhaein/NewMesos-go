# Note ?= 할당되지 않으면, 할당한다.
TEST_LDFLAGS ?=
TEST_FLAGS ?= -race
CAP_RESERVATION_REFINEMENT ?= true

# $(MAKE) _test CAP_RESERVATION_REFINEMENT=false
.PHONY: test
test:
    $(MAKE) _test CAP_RESERVATION_REFINEMENT=true

.PHONY: _test
_test:
	go test $(TEST_FLAGS) $(TEST_DIRS)

.PHONY: start
start:
	docker-compose build
	docker-compose up

.PHONY: stop
stop:
	docker-compose stop -t 0
	docker-compose rm -f