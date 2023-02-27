# GOpenAI Command line tool

This is a Command line tool built with GOpenAI, it allows you to call OpenAI api from terminal without any code.

It's very convenient for testing.

## Build

```shell
go build
```

## Install

Run following command in tool directory.

```shell
go install
```

It will install gopenai to GOPATH/bin.

## Usage

### Create Completion
```shell
OPENAI_API_KEY=yourkey gopenai completion create --model=text-davinci-003 --prompt="Say this is a test" --max-tokens=200
```
### List Models
```
OPENAI_API_KEY=yourkey gopenai model list
```

### More

use 
```shell
gopenai --help
```
for more details

## Development

* [x] Model
* [x] Completion
* [x] Edit
* [x] Images
* [x] Embeddings
* [x] Files
* [x] Fine-tunes
* [x] Moderations
