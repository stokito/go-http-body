# go-http-body
Read Golang http.Request.Body quickly and in effective way.
The default way is to use `io.ReadAll(r.Body)` but it makes a lot of memory re-allocations and copies.
Instead we can allocate a slice with `Content-Lenght` from beginning.


## Usage

```go
import (
	http_body "github.com/stokito/go-http-body"
)

func handleUpdateConf(w http.ResponseWriter, r *http.Request) {
	defer http_body.CloseBody(&r.Body) // safely close body even if it's nil
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	reqBody, err := http_body.ReadHttpBody(r.Body, r.Header)
	println(string(reqBody))
	w.WriteHeader(http.StatusNoContent)
}
```

## Install

    go get -u github.com/stokito/go-http-body


## License

[0BSD](https://opensource.org/licenses/0BSD) (similar to Public Domain)
