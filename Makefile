.PHONY: build lint clean package push

default: build

build:
	go build -a -o ./bin/rubusidaeus *.go

install:
	go install .

build_arm6:
	GOOS=linux GOARM=6 GOARCH=arm go build -a -o ./bin/rubusidaeus *.go
	
package: build_arm6
	docker build -t dhogborg/rubusidaeus:latest .
    
push: package
	docker push dhogborg/rubusidaeus:latest
	
lint:
	golint .

clean:
	- rm -r bin