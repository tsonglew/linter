# linter

linter website source code

## Support

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

Install Python, Pip, Golang First, and then

```sh
$ git clone https://github.com/kasheemlew/linter.git
```

Install pylint and cpplint(both in Google Coding Style)

```sh
$ pip install pylint
$ pip install cpplint
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