# GOpenAI Command line tool

This is a Command line tool built with GOpenAI, it allows you to call OpenAI api from terminal without any code.

It's very convenient for testing.

## Build

```shell
make build
```

It will build executable file located in ./bin/gopenai.

## Install

Run following command in tool directory.

```shell
make install
```

It will install gopenai to GOPATH/bin.

## Usage

### Create Completion
```shell
OPENAI_API_KEY=yourkey gopenai completion create --model=text-davinci-003 --prompt="Say this is a test" --max-tokens=200
```
### List Models
```shell
OPENAI_API_KEY=yourkey gopenai model list
```

### Chat
```shell
gopenai chat create -m 'gpt-3.5-turbo' --messages "[{\"role\":\"user\",\"content\":\"Who are you?\"}]"
```

### More

use 
```shell
gopenai --help
```
for more details

## Features

* [x] Model
  * [x] Create
  * [x] Retrieve
  * [x] Delete
* [x] Completion
  * [x] Create
* [x] Chat
  * [x] Create
* [x] Edit
  * [x] Create
* [x] Images
  * [x] Create
  * [x] Edit
  * [x] Variation
* [x] Embeddings
  * [x] Create
* [x] Files
  * [x] Create
  * [x] Retrieve
  * [x] Download
* [x] Fine-tunes
  * [x] Create
  * [x] Retrieve
  * [x] List
  * [x] Cancel
  * [x] Events
* [x] Moderations
  * [x] Create
