![Build](https://github.com/joonas-fi/pwdgen/workflows/Build/badge.svg)
[![DockerHub](https://img.shields.io/docker/pulls/joonas/pwdgen.svg?style=for-the-badge)](https://hub.docker.com/r/joonas/pwdgen/)

Super-simple password generator.


Why
---

I didn't quickly find one that I could trust (made in a memory safe language).


Usage
-----

```console
$ pwdgen
n6uLyBpj7-9E2o=r

$ pwdgen --help
Password generator

Usage:
  pwdgen [flags]

Flags:
  -h, --help         help for pwdgen
  -l, --length int   Length of password (default 16)
  -v, --version      version for pwdgen
```
