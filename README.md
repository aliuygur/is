# is

[![Build Status](https://travis-ci.org/alioygur/is.svg?branch=master)](https://travis-ci.org/alioygur/is)
[![GoDoc](https://godoc.org/github.com/alioygur/is?status.svg)](https://godoc.org/github.com/alioygur/is)
[![Go Report Card](https://goreportcard.com/badge/github.com/alioygur/is)](https://goreportcard.com/report/github.com/alioygur/is)

Micro check library in Golang. 

## installation

`go get gopkg.in/alioygur/is.v1`

## usage

```go
package main

import "gopkg.in/alioygur/is.v1"
import "log"

func main()  {
    is.Email("jhon@example.com") // true
    is.Numeric("Ⅸ") // false
    is.UTFNumeric("Ⅸ") // true
}
```


for more documentation [godoc](https://godoc.org/github.com/alioygur/is)

## Contribute

**we are waiting your contribution**

- Report problems
- Add/Suggest new features/recipes
- Improve/fix documentation

Many thanks to our contributors: [contributors](https://github.com/alioygur/is/graphs/contributors)

## Thanks & Authors

I use code/got inspiration from these excellent libraries:

- [arasatasaygin/is.js](https://github.com/arasatasaygin/is.js) Micro check library
- [asaskevich/govalidator](https://github.com/asaskevich/govalidator) [Go] Package of validators and sanitizers for strings, numerics, slices and structs