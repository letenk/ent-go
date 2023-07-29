run:
	go run main.go

schema:
	@echo "Creating new schema..."
	go run -mod=mod entgo.io/ent/cmd/ent new $(arg)
	@echo "Finish"


generate:
	@echo "Generating ent schema..."
	go generate ./ent
	@echo "Finish"