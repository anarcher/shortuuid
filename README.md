# shortuuid.go

``shortuuid`` is a simple go library that generates concise, unambiguous,
+URL-safe UUIDs. 

It's inspired by [shortuuid python library](https://github.com/stochastic-technologies/shortuuid)


## Usage

```
package main

import (
    "fmt"
    "github.com/anarcher/shortuuid"
)

func main() {
    id := shortuuid.New()
   fmt.Println(id)
}

$go run main.go
wusrj4VHumEALr4mQjArnj
$go run main.go
YTPp9cZDECzweSyywSespm

```

