# Requests

## Features

### Get

```go
package main

import (
	"fmt"
    "github.com/gzeroner/requests"
    "log"
)

func main() {
	res, err :=requests.Get("https://bing.com")
	if err != nil {
        log.Fatal(err)
	}
	fmt.Println(res.String())
}
```

### Post

```go
package main

import (
	"fmt"
	"github.com/gzeroner/requests"
	"log"
)

func main() {
	res, err :=requests.Post("https://bing.com", nil)
	if err != nil {
        log.Fatal(err)
	}
	fmt.Println(res.String())
}
```

### Delete

```go
package main

import (
	"fmt"
	"github.com/gzeroner/requests"
	"log"
)

func main() {
	res, err :=requests.Delete("https://bing.com")
	if err != nil {
        log.Fatal(err)
	}
	fmt.Println(res.String())
}
```

### Put

```go
package main

import (
	"fmt"
	"github.com/gzeroner/requests"
	"log"
)

func main() {
	res, err :=requests.Put("https://bing.com", nil)
	if err != nil {
        log.Fatal(err)
	}
	fmt.Println(res.String())
}
```

### Head

```go
package main

import (
	"fmt"
	"github.com/gzeroner/requests"
	"log"
)

func main() {
	res, err :=requests.Head("https://bing.com")
	if err != nil {
        log.Fatal(err)
	}
	fmt.Println(res.String())
}
```

### Options

```go
package main

import (
	"fmt"
	"github.com/gzeroner/requests"
	"log"
)

func main() {
	res, err :=requests.Options("https://bing.com")
	if err != nil {
        log.Fatal(err)
	}
	fmt.Println(res.String())
}
```

## Todo

- [ ] Support file upload
- [ ] Support form-data
- [ ] Support x-www-form-urlencoded
- [x] Support application/json
- [ ] Support cookies
- [x] Support authentication