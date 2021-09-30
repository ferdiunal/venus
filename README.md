# Venus

It allows us to use Venus hashids more easily.

```go
package main

func main() {

  salt := "CbF3v1V33H6mYeh9ygyy8rLKB4MmVn"
  len := 8
  v := venus.New(salt, len)

  encoded := v.Encode(8)
  
  fmt.Println(encoded)

  fmt.Println(v.Decode(encoded))

}

```

Third-party library licenses

[go-hashids](https://github.com/speps/go-hashids/blob/master/LICENSE)
