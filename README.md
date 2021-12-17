# go118-any-migration-example

For Go1.17 or older, you prefer to add following code if you want to use `any` type.

```go
package main

// go:build !go1.18

type any = interface{}
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
