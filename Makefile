.PHONY: build

build:
	@echo "Building a docker image of fachrin/dns-resolver"
	docker build . -t fachrin/dns-resolver -f ./build/Dockerfile --compress

run: build
	@echo "Running a docker image of fachrin/dns-resolver on http://localhost:9999"
	docker run -p 9999:80 -d fachrin/dns-resolver
	@echo "Successfully running the docker image of fachrin/dns-resolver on http://localhost:9999"
	@echo "Open http://localhost:9999 on your browser"
