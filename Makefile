
MONGODB_URI = "mongodb://admin:admin@localhost:27017/?connect=direct"

build:
	go build -o bongo

run:
	go run . --mongodb-uri $(MONGODB_URI)