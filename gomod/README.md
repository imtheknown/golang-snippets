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

On a more general note, the purpose of ```shell $ go mod tidy ```is to also add any dependencies needed for other combinations of OS, architecture, and build tags. Make sure to run this before every release.

