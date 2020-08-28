# GO mod

For a quick start, if your project is under a VCS,

```shell
$ go mod init
```

To Supply the path manually,

```shell
$ go mod github.com/imtheknown/example
```

This command will create go.mod file which both defines projects requirements and locks dependencies to their correct versions like package.json you may have already used.
The primary motivation for Go modules was to improve the experience of using (that is, adding a dependency on) code written by other developers.

On a more general note, the purpose of ``` go mod tidy ```is to also add any dependencies needed for other combinations of OS, architecture, and build tags. Make sure to run this before every release.

Commands like go build or go test will automatically download all the missing dependencies though you can do this explicitly with go mod download to pre-fill local caches which may prove useful in CI.

Visit [This Page](https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31).