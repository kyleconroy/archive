.PHONY: golden

golden: instagram/testdata/golden.bin twitter/testdata/golden.bin

instagram/testdata/golden.bin:
		go run internal/golden/main.go instagram

twitter/testdata/golden.bin:
		go run internal/golden/main.go twitter
