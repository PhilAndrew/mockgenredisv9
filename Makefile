SHELL = /bin/bash
.SHELLFLAGS = -o pipefail -c
golang_version="1.20"

## GITHUB_USER defines the GitHub
## username to use when downloading
## private dependencies
## GITHUB_USER=""

## GITHUB_TOKEN defines the GitHUb
## user's password or API token to use when
## downloading private dependencies
##
## GITHUB_TOKEN=""

## test runs the Golang unit tests
## for the package ensuring it is ok
## to build and package
test:
	@echo -e "\033[32m===# Preparing docker image for unit tests\033[0m"
	@docker build --build-arg GITHUB_USER="${GITHUB_USER}" --build-arg GITHUB_TOKEN="${GITHUB_TOKEN}" \
				  --build-arg GOLANG_VERSION="${golang_version}" \
    	          -t golang-unit-test-image -f ./.ci/Dockerfile-UnitTests .

	@echo -e "\033[32m===# Executing unit tests inside docker container\033[0m"
	@docker run --rm golang-unit-test-image

	@echo -e "\033[32m===# Deleting unit test image\033[0m"
	@docker rmi golang-unit-test-image
