# GOpenAI Command line tool

This is a Command line tool built with GOpenAI, it allows you to call OpenAI api from terminal without any code.

It's very convenient for testing.

## Usage

### Create Completion
```shell
OPENAI_API_KEY=yourkey gopenai completion create --model=...
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
* [ ] Edit
* [ ] Images
* [ ] Embeddings
* [ ] Files
* [ ] Fine-tunes
* [ ] Moderations
