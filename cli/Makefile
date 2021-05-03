all: clean builds/coa-nft-cli_darwin-amd64 builds/coa-nft-cli_darwin-arm64 builds/coa-nft-cli_linux-amd64 builds/coa-nft-cli_linux-arm64 builds/checksum.sha256 builds/validate.sh builds/release.zip
clean:
	rm -rf builds
.PHONY: build clean

builds:
	mkdir -p builds

builds/coa-nft-cli_darwin-amd64: builds
	GOOS=darwin GOARCH=amd64 go build -o $@ .

builds/coa-nft-cli_darwin-arm64: builds
	GOOS=darwin GOARCH=arm64 go build -o $@ .

builds/coa-nft-cli_linux-amd64: builds
	GOOS=linux GOARCH=amd64 go build -o $@ .

builds/coa-nft-cli_linux-arm64: builds
	GOOS=linux GOARCH=arm64 go build -o $@ .

builds/checksum.sha256: builds
	cd builds && shasum -a 256 * > checksum.sha256

builds/validate.sh:
	echo "#!/bin/bash" > $@
	echo "shasum -c checksum.sha256" >> $@
	chmod +x $@

builds/release.zip:
	zip -j -9 $@ builds/*
