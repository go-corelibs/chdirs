[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/chdirs)
[![codecov](https://codecov.io/gh/go-corelibs/chdirs/graph/badge.svg?token=8I3llBeXkL)](https://codecov.io/gh/go-corelibs/chdirs)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/chdirs)](https://goreportcard.com/report/github.com/go-corelibs/chdirs)

# chdirs - push/pop current working directory

chdirs is a package for managing the current working directory in a similar
fashion to the unix shell `pushd` and `popd` functions.

# Installation

``` shell
> go get github.com/go-corelibs/chdirs@latest
```

# Examples

## Push, Pop

``` go
func main() {
    // need to do something in another directory
    // and ensure the program is returned to the
    // current working directory
    if err := chdirs.Push(".."); err != nil {
        panic(err)
    }
    defer chdirs.Pop()
    // do more stuff in the pushed directory...
    // the above defer will pop the path to the
    // original working directory when all is
    // done
}
```

## MockBadWD, UnMockBadWD

``` go
// need to test important code when os.Getwd or filepath.Abs may fail,
// normally these tests are acceptably not written, however for command
// line programs where the current working directory could be anything,
// this testing is in fact quite necessary...

func Test(t *testing.T) {
    Convey("handling of os.Getwd/filepath.Abs errors", t, func() {
        err := MockBadWD()
        So(err, ShouldEqual, nil)
        /*
          ... test code handling os.Getwd and/or filepath.Abs errors ...
        */
        err = UnMockBadWD()
        So(err, ShouldEqual, nil)
    })
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2023 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
