EXE=wallpape

build:
	go build -o $(EXE)

run:
	go build -o $(EXE)
	./$(EXE)
