gomock
======

GoMock is a mocking framework for the [Go programming language](https://golang.org). It
integrates well with Go's built-in `testing` package, but can be used in other
contexts too. This is a fork that allows the gomock library to print its output 
as a diff. This is especially useful when the mock receives a map or a slice
and errors. Instead of showing a gigantic blob, wrapping in the terminal several
times, you get just the parts of the argument that are different.


Installation
------------

Install and gomock/mockgen as normal, then:

    mkdir vendor
    git clone git@github.com:fishybell/mock vendor/mock
    echo "replace github.com/golang/mock => ./vendor/mock" >> go.mod

This is necessary as the mock package uses types rather than interfaces.
