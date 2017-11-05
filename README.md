# linter

linter website source code

## Support

Now available on MacOS and Linux

- [x] Python(pylint)
- [x] C++(cpplint)
- [x] C(cpplint)
- [ ] Java
- [ ] JavaScript
- [ ] Go
- [ ] C#
- [ ] Ruby
- [ ] Swift
- [ ] Perl
- [ ] Shell

## Deploy

### Raw

Install Golang First, and then get the code from GitHub

```sh
$ git clone https://github.com/kasheemlew/linter.git
```

Build to present working directory

```sh
$ make build
```

Install binary to go-path/bin

```sh
$ make install
```

### with Docker(halted)

```sh
$ git clone https://github.com/kasheemlew/linter.git
$ docker build -t linter .
$ docker run --publish 6060:8080 --name linter --rm linter -d
```

## Maintainers

Front-end: Elegenthus

Back-end: kasheemlew