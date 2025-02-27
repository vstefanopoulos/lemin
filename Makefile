.PHONY: build run fmt clean test vis git debug dockerRun dockerRemove
BINARY_NAME=letmein
SRC= ./main.go
ARGS =  badexample00.txt \
		badexample01.txt \
		badexample02.txt \
		example00.txt \
		example01.txt \
		example02.txt \
		example03.txt \
		example04.txt \
		example05.txt \
		example06.txt \
		example07.txt \
		example08.txt

build:
	go build -o ./$(BINARY_NAME) $(SRC)
	
run:
	for arg in $(ARGS); do ./letmein $$arg; done

fmt:
	go fmt ./main.go 
	go fmt ./lemin/moves/*
	go fmt ./lemin/colony/*
	go fmt ./lemin/handle/*
	go fmt ./lemin/paths/*

clean:
	rm -f ./$(BINARY_NAME)
	rm -f ./output.txt

test:
	go test ./lemin/colony/
	go test ./lemin/handle/

vis:
	./letmein $(f) > output.txt
	python3 python/ant.py

git:
	@$(MAKE) clean
	git add .
	git commit -m "$(m)"
	git push

debug:
	@make clean
	@make build
	@make run
