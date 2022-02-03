snapshot:
	goreleaser build --snapshot -p 10 --rm-dist
clean:
	rm -rf dist

@phony: clean shapshot
