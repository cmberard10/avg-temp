GO      				= go
TIMEOUT 				= 300
TESTPKGS 				= $(shell env GO111MODULE=on $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))
TEST_TARGETS = test-default test-bench test-short test-verbose test-race

.PHONY: lint
lint:
	go get -u golang.org/x/lint/golint
	$Q golint -set_exit_status ./...




.PHONY: $(TEST_TARGETS) test-xml check test tests
test-bench:   ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-short:   ARGS=-short        ## Run only short tests
test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Run tests with race detector
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
check test tests: fmt gomock lint ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	$Q $(GO) test ./... -timeout $(TIMEOUT)s $(ARGS) $(TESTPKGS)

#fmt
.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	$Q $(GO) fmt ./...

.PHONY: lint
	$Q $(GO) lint ./...

.PHONY: gomock
gomock: gomocksetup gomockhttpdao gomockscriptdao

.PHONY: gomocksetup
gomocksetup:
	  go get -u github.com/golang/mock/gomock
	  go get -u github.com/golang/mock/mockgen
	  rm -rfv mock
	  mkdir -p mock

.PHONY: gomockhttpdao
gomockhttpdao:
	mockgen -source=internal/services/http/dao.go -destination=mock/mockHTTPDAO/mock_http_dao.go -package=mockHTTPDao

.PHONY: gomockscriptdao
gomockscriptdao:
	mockgen -source=internal/services/dao.go -destination=mock/mockSCRIPTDAO/mock_script_dao.go -package=mockSCRIPTDao
