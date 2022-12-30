
MONGODB_URI = "mongodb://admin:admin@localhost:27017/?connect=direct"

build:
	mkdir build releases
	GOOS=darwin GOARCH=arm64 go build -o ./build/bongo_darwin_arm64/bongo && \
		tar --strip-components=1 -czf releases/bongo_darwin_arm64.tar.gz build/bongo_darwin_arm64/
	GOOS=linux  GOARCH=amd64 go build -o ./build/bongo_linux_x86_64/bongo && \
		tar --strip-components=1 -czf releases/bongo_linux_x86_64.tar.gz build/bongo_linux_x86_64/

clean:
	rm -rf build releases 2>/dev/null

run:
	go run . --mongodb-uri $(MONGODB_URI)