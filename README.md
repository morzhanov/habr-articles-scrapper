# habr-articles-scrapper

Golang console application to retrieve top-x habr articles.

<img src="https://i.ibb.co/ZTwCQbn/1632212887949.png" alt="1632212887949">

## Usage

To run app with default params:
```bash
go run ./cmd/main.go
```

### Params
- `top` - top results filter (available options are: 0, 10, 25, 50, 100)
- `pages` - number of pages to parse

Example:
```bash
go run ./cmd/main.go -top=25 -pages=10
```
