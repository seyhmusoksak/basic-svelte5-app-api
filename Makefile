NAME = apiApp
COMPILER = go

build:
	$(COMPILER) build -o $(NAME) main.go
	@echo "Build complete"

run:
	@echo "Running application"
	$(COMPILER) run main.go

fclean:
	rm -rf $(NAME)
	@echo "Cleaned up build files"


.PHONY: all build run fclean
