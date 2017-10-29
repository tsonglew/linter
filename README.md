# linter

linter website source code

## Support

- [x] Python(pylint)
- [x] Cpp(cpplint)

## Deploy

### Raw

Install Python, Pip, Golang First, and then

```bash
$ git clone https://github.com/kasheemlew/linter.git
```

Install pylint and cpplint(powered by Google Coding Style)

```bash
```

### with Docker(halted)

```bash
$ git clone https://github.com/kasheemlew/linter.git
$ docker build -t linter .
$ docker run --publish 6060:8080 --name linter --rm linter -d
```

## Maintainers

Front-end: Elegenthus
Back-end: kasheemlew