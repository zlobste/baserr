# `baserr` - Go Package Documentation

## Overview

The `baserr` package in Go provides a simple and effective way to create inherited errors, similar to the inheritance of
exceptions in object-oriented programming languages. This approach allows for more specific error handling in Go
applications.

## Installation

```
go get github.com/zlobste/baserr
```

## Usage

#### Import the `baserr` package into your Go project:

```go
import (
    "github.com/zlobste/baserr"
)
```
#### Define your own error types:

```go
// OSError -> IOError -> WriteError
type (
    // OSError represents a first level error. It is used as parent for other errors.
    OSError struct {
        baserr.Base[baserr.BaseError]
    }
    // IOError represents an error inherited from OSError.
    IOError struct {
        baserr.Base[OSError]
    }
    // WriteError represents an error inherited from IOError.
    WriteError struct {
        baserr.Base[IOError]
    }
)
```


#### Create typed errors:

```go
func execute() error {
    return baserr.NewError[WriteError]("EOF")
}
```

#### Check if an error is inherited from a specific error type:

```go
if err := execute(); err != nil {
    switch {
    case baserr.InheritedFrom[CustomError](err):
        fmt.Println("custom")
    case baserr.InheritedFrom[OSError](err):
        fmt.Println("write error")
    }
} else {
    fmt.Println("no error occurred")
}
```

Output:

```
write error
```