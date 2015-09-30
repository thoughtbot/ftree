Contributing
============

We love pull requests from everyone.

Fork the repo.

Get a working [Go installation],
and clone the project into your [Go work environment]
\(that is, `$GOPATH/src/github.com/calebthompson/ftree`).
If your Go environment is already set up, you can get the code instead with
`go get github.com/calebthompson/ftree`. It will be in the same location as the
previous path.

[Go installation]: http://golang.org/doc/install
[Go work environment]: http://golang.org/doc/code.html

## Dependencies

We use `godep` to manage dependencies. You can install that with
`go get github.com/tools/godep`. You don't need to unless you're making changes
to the dependencies.

If you add or update a dependency,
run `godep save ./...` to vendor the changes.

## Developing

You can compile and run the current code with `go run main.go`.

## Testing

To test the `ftree` package, run `go test ./...`.

## Installing

To install the `ftree` binary, run `go install ./...`.

## Building

To build the `ftree` binary, run `go build`. `go test` and `go install` do
this for you, but if you need the artifact this is how you can get that.

## Changes

Make your change, with new passing tests.

Push to your fork. Write a [good commit message][commit]. Submit a pull request.

[commit]: https://robots.thoughtbot.com/5-useful-tips-for-a-better-commit-messag://robots.thoughtbot.com/5-useful-tips-for-a-better-commit-message

Others will give constructive feedback.
This is a time for discussion and improvements,
and making the necessary changes will be required before we can
merge the contribution.
