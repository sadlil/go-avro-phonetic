# go-avro-phonetic
Go port of the popular Bengali phonetic-typing software `Avro Phonetic`. 
[website](http://omicronlab.com)


## Overview
`go-avro-phonetic` provides a `go` package which includes a text parser
that converts Bangla written in Roman script to its phonetic
equivalent in Bangla. It implements the **Avro Phonetic Dictionary
Search Library** by [Mehdi Hasan Khan](https://github.com/mugli).

In *addition* this library provides a [cli](#cli) to parse text file or 
text via console.


### Inspirations
This package is inspired from [Rifat Nabi](https://github.com/torifat)'s [jsAvroPhonetic](https://github.com/torifat/jsAvroPhonetic). 
So far, the code is a go port of `jsAvroPhonetic`.

And acknowledges [Kaustav Das Modak](https://github.com/kaustavdm) for [pyAvroPhonetic](https://github.com/kaustavdm/pyAvroPhonetic). 


## Installation

```bash

$ go get -u -v github.com/sadlil/go-avro-phonetic/...

```

## Usage
Import the library as
```go

import (
    avro "github.com/sadlil/go-avro-phonetic"
)
```

With Built in Rules:
```
// Parse() tries to parse the given string
// In case of failure it returns an erros
text, err := avro.Parse("ami banglay gan gai")
if err != nil {
    // Handle error
}

fmt.Println(test) // আমি বাংলায় গান গাই


// MustParse() tries to parse the given string
// In case of failure it panics
text := avro.MustParse("ami banglay gan gai")
fmt.Println(text) // আমি বাংলায় গান গাই

```

With Custom Rules:
`avro.ParseWith()` receives an plugable `Dictionary` `interface{}` to support
custom parsing.

`data.LoadJSON()` provides support for overloading custom dictionary from JSON,
which can be used for custom parsing.

```go
customDictonary := []byte("custom dictonary json")

dict, err := data.LoadJSON(customDictonary)
if err == nil {
    text := avro.ParseWith(dict, "ami banglay gan gai")
    fmt.Println(text) // custom parsed text
}

```

## CLI
Avro command line interface to parse data using command line.

#### Install

```bash
$ go get -u -v github.com/sadlil/go-avro-phonetic/avrocli
$ go install github.com/sadlil/go-avro-phonetic/avrocli
```

#### Usages
Parse a test:
```bash
$ avrocli parse ami banglay gan gai
আমি বাংলায় গান গাই  # this could need font support for cli.
```

Parse a file:
```bash
$ avrocli parse -f my_bangla.txt # this will create a my_bangla.bn.txt file
                                 # in the same directory as the given file
                                 # with parsed bangla texts.
```


## Acknowledgements
 - Mehdi Hasan Khan for originally developing and maintaining Avro Phonetic,
 - Rifat Nabi for porting it to Javascript,
 - [Sarim Khan](https://github.com/sarim) for writing ibus-avro.


## License
Licensed under [MIT Licence](LICENSE).