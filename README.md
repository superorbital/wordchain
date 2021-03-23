# Word Chain

## Overview

`wordchain` is an application that can generate readable chains of customizable words for naming things like containers, clusters, and other objects.

## Development

```go
go get github.com/markbates/pkger/cmd/pkger
pkger -include /data/words.json
go build .
```

## Releasing

* Update the Version in `cmd/root.go`

```shell
go get github.com/mitchellh/gox
gox -osarch='!darwin/386' -output './builds/wordchain_{{.OS}}_{{.Arch}}'
```

* Create a release in Github with the resulting binaries.

### Docker Release

```shell
$ docker build -t superorbital/wordchain:${VERISON} .
$ docker push superorbital/wordchain:${VERISON}
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

./wordchain random -s "someone-elses-unique-git-branch-name"
minty-sleet

./wordchain random -s "someone-elses-unique-git-branch-name"
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
$ ./wordchain export > internal-word-list.json
```

### Docker

```shell
$ docker run superorbital/wordchain:latest
alpha-drink

$ docker run superorbital/wordchain:latest random -l 3
cut-oak
```

### Library

* In you project you can do something like this to use this as a library:

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

## Ideas

* Add viper config file w/ env support
* Allow for a range of word length (like 3-5 characters)
* extend list (versus completely replace)
* Blacklist words (list words you never want in your results)
* Advanced templating of output
