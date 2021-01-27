EXE=wallpaper-changer

build:
	go build -o $(EXE)

run:
	go build -o $(EXE)
	./$(EXE)
