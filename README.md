# plop

> Writes stdin to a new temporary file

## Installation

1. Download the latest package for your platform from the [Releases page](https://github.com/justincampbell/plop/releases/latest).
2. Untar the package with `tar -zxvf plop*.tar.gz`.
3. Move the extracted `plop` file to a directory in your `$PATH` (for most systems, this will be `/usr/local/bin/`).

Or, if you have a [Go development environment](https://golang.org/doc/install):

```
go get github.com/justincampbell/plop
```

## Usage

```
$ grep error /var/log/error.log | plop
/tmp/plop-6UCoXscu
$ grep error /var/log/error.log | plop -t errors
/tmp/errors-9WrDDfbX
$ grep error /var/log/error.log | plop -d /var/log/
/var/log/plop-UnmVhLi9
```

## Other Implementations

* **Bash**: [samrocketman/home](https://github.com/samrocketman/home/blob/master/bin/plop)
