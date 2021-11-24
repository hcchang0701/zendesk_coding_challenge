install:
	go build -o tv cmd/ticketViewer/main.go && mv tv $$GOPATH/bin

clean:
	rm $$GOPATH/bin/tv