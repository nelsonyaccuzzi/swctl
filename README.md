[![Build Status](https://travis-ci.org/sitewhere/swctl.svg?branch=master)](https://travis-ci.org/sitewhere/swctl) [![Go Report Card](https://goreportcard.com/badge/github.com/sitewhere/swctl)](https://goreportcard.com/report/github.com/sitewhere/swctl) [![GoDoc](https://godoc.org/github.com/sitewhere/swctl?status.svg)](https://godoc.org/github.com/sitewhere/swctl)

![SiteWhere](https://s3.amazonaws.com/sitewhere-branding/SiteWhereLogo.svg)

---

# SiteWhere Control CLI

## Build

For building it requires go 1.11+.

```console
go build
```

## Install swctl

### From source code

```console
go install
```

### Install binary with curl on Linux

```bash
curl -L https://github.com/sitewhere/swctl/releases/latest/download/swctl.linux.amd64 -o swctl && \
chmod +x ./swctl && sudo mv ./swctl /usr/local/bin/swctl
```

### Install binary with curl on macOS

```bash
curl -L https://github.com/sitewhere/swctl/releases/latest/download/swctl.darwin.amd64 -o swctl && \
chmod +x ./swctl && sudo mv ./swctl /usr/local/bin/swctl
```

### Install binary with curl on Windows

```bash
curl -L https://github.com/sitewhere/swctl/releases/latest/download/swctl.windows.amd64.exe -o swctl.exe
```
