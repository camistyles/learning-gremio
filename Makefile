.PHONY: build clean deploy

FUNCTION =
FUNCTION += register



clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose
test:
	echo "$(@) esto es lo que me llega"

npm-cle:


