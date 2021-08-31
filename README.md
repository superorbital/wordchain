# Word Chain

![wordchain](./images/wordchain-black_on_white_cropped.png)

## Overview

`wordchain` is an application that can generate readable chains of customizable words for naming things like containers, clusters, and other objects.

## Development

```shell
$ go get github.com/markbates/pkger/cmd/pkger
$ pkger -include /data/words.json
# Install go-swagger
# https://goswagger.io/install.html
$ swagger generate server -f ./swagger.yaml --exclude-main
$ go mod vendor
$ go build .
```

## Releasing

* Update the Version in `cmd/root.go`

```shell
$ swagger generate server -f ./swagger.yaml --exclude-main
$ go install github.com/mitchellh/gox@latest
$ gox -osarch='darwin/amd64 darwin/arm64 freebsd/386 freebsd/amd64 freebsd/arm linux/386 linux/amd64 linux/arm linux/arm64 linux/mips linux/mips64 linux/mips64le linux/mipsle linux/s390x netbsd/386 netbsd/amd64 netbsd/arm openbsd/386 openbsd/amd64 windows/386 windows/amd64' -output './builds/wordchain_{{.OS}}_{{.Arch}}'
```

* Create a release in Github with the resulting binaries.

### Docker Release

* See: https://www.docker.com/blog/multi-arch-build-and-images-the-simple-way/

```shell
$ swagger generate server -f ./swagger.yaml --exclude-main
$ docker buildx build \
    --push \
    --platform linux/arm/v7,linux/arm64/v8,linux/amd64 \
    --tag superorbital/wordchain:${VERISON} .
```

## Usage

### CLI

* Use the built-in list to get a 5 letter adj/noun chain

```shell
$ wordchain random
sudsy-cloak
```

* Use the built-in list to get a 4 letter adj/noun chain

```shell
$ wordchain random -l 4
holy-wart
```

* Adjust the words types in the chain

```shell
$ wordchain random -t adjective,adjective,noun
fried-picky-month
```

* Get deterministic results from your list by providing a seed string

```shell
$ wordchain random -s "my-unique-git-branch-name"
mangy-berry

$ wordchain random -s "my-unique-git-branch-name"
mangy-berry

$ wordchain random -s "someone-elses-unique-git-branch-name"
minty-sleet

$ wordchain random -s "someone-elses-unique-git-branch-name"
minty-sleet
```

* Provide a custom list to get a 3 letter adj/noun chain

```shell
$ wordchain random -j ./data/tests/words.json -l 3
tan-nap
```

* Get a 3 letter adj/noun chain with a custom divider and pre-pended and post-pended word.

```shell
$ wordchain random -l 3 -d + -r hello -o goodbye
hello+odd+pad+goodbye
```

* Get a copy of the internal word list in the valid json format

```shell
$ wordchain export > internal-word-list.json
```

### Microservice

```shell
$ wordchain listen --port 8080
2021/03/30 11:32:02 Serving word chains at http://[::]:8080

$ curl -X POST -d '{}' -H 'Content-Type: application/json' http://127.0.0.1:8080/v1/random
"{\"chain\":\"quack-bayou\"}"

$ curl -X POST -d '{"length": 3}' -H 'Content-Type: application/json' http://127.0.0.1:8080/v1/random
"{\"chain\":\"odd-toy\"}"

$ curl -X POST -d '{"divider": "_", "length": 3, "prepend": "hello", "postpend": "adios", "seed": "deterministic" }' -H 'Content-Type: application/json' http://127.0.0.1:8080/v1/random
"{\"chain\":\"hello_bad_ace_adios\"}"
```

### Docker

```shell
$ docker run superorbital/wordchain:latest
2021/03/30 20:35:07 Serving word chains at http://[::]:8080

$ docker run superorbital/wordchain:latest random
alpha-drink

$ docker run superorbital/wordchain:latest random -l 3
cut-oak
```

### Library

* In your project you can do something like this to use this as a library:

```go
package main

import (
	types "github.com/superorbital/wordchain/types"
	words "github.com/superorbital/wordchain/words"
)

func main() {
	prefs := types.Preferences{
		WordFile: "",
		Length:   5,
		Divider:  "-",
		Prepend:  "",
		Postpend: "",
		Seed:     "",
		Type:     []string{"adjective", "noun"},
	}
	words.Random(prefs)
}
```

It will expect a data file of words to exist. You can either create a valid JSON data file in your project at `data/words.json` or you can copy `wordchain/pkged.go` into your project to use the same word list that is embedded in the `wordchain` binary by default.

## Word List Format

```json
{
  "lists":
  [
    {
      "type": "adjective",
      "length": 4,
      "words": [
        "arid",
        "blue"
      ]
    },
    {
      "type": "noun",
      "length": 4,
      "words": [
        "tent",
        "tree"
      ]
    }
  ]
}
```

## TODO & Ideas

* Add tests
* Make listener use defaults passed in at the command line
* Add swagger into Dockerfile (w/ multi-arch support)
* Create script for Github and Docker versioned releases
* Add a homebrew/linuxbrew tap
* Add basic Kubernetes Deployment Manifest
* Add basic /health endpoint
* Create a single page Github site for tool
* Add a small demo video and maybe even a one page tool github page at `wordchain.superorbital.io`
* Add HTTPS support to listener mode.
* Add viper config file w/ env support
* extend list (versus completely replace)
* Blacklist words (list words you never want in your results)
* Advanced templating of output

## Acknowledgements

* Original `wordchain` logo by [@fredyates](https://github.com/fredyates)
