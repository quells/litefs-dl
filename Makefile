all: litefs-dl

.PHONY: clean
clean:
	-rm litefs-dl

litefs-dl: main.go
	go build -o litefs-dl .
