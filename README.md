# Go Check

Just a simple re-usable snippet for unit testing boilerplate.

## Usage

You would do best to copy the relevant content of the `gocheck.go` file to your project rather than add as a dependency.

If you do prefer to keep up with this repo (and it will be slow moving so why bother)

```go
// Use it in your go code

import (
    ...
    "github.com/taikedz/gocheck"
    ...
)
```

Then run once one of the following:

```sh
# To pin (preferred):
go get github.com/taikedz/gocheck@v1.0.0

# -- OR --

# To get the current latest
go mod tidy
```

## Usage

See `unittests/gocheck_test.go` for an example.

# Maintenance

This repo is maintained by its author for its author. No guarantees are given. I would rather you copy this snippet and maintain your own copy for your own needs.
