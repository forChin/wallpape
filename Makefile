EXE=wallpape
OPERSYSTEMS=windows darwin linux
ARCHS=amd64 386

build:
	go build -o $(EXE)

run: build
	./$(EXE)

all:
	@for os in $(OPERSYSTEMS); \
	do \
		for arch in $(ARCHS); \
		do \
			if [ $$os = "windows" ]; then \
				GOOS=$$os GOARCH=$$arch go build $(LDFLAGS) -o $(DIR)/$(EXE)-$$os-$$arch.exe ; \
			else \
				GOOS=$$os GOARCH=$$arch go build $(LDFLAGS) -o $(DIR)/$(EXE)-$$os-$$arch ; \
			fi \
		done \
	done
