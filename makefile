APP_NAME=draw

build:
	go build -o $(APP_NAME) main.go

run:
	./$(APP_NAME) $(ARGS)

clean:
	rm -f $(APP_NAME)

