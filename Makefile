all: ssqr compress

ssqr:
	@go build -ldflags '-w -s' -o ssqr main.go

compress:
	@upx ssqr

test:
	@./ssqr shadow.json

clean:
	@rm ssqr
