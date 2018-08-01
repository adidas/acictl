# AdiCTL
[![Build Status](https://travis-ci.com/jorgechato/acictl.svg?token=x3vLcsQVEzf1kfJyx1Uv&branch=master)](https://travis-ci.com/jorgechato/acictl)
[![Build Status](https://sonarcloud.io/api/project_badges/measure?project=com.adidas.acictl&metric=coverage)](https://sonarcloud.io/dashboard?id=com.adidas.acictl)
[![Build Status](https://sonarcloud.io/api/project_badges/measure?project=com.adidas.acictl&metric=alert_status)](https://sonarcloud.io/dashboard?id=com.adidas.acictl)
[![Go Report Card](https://goreportcard.com/badge/github.com/adidas/adictl)](https://goreportcard.com/report/github.com/adidas/adictl)
[![Godoc](https://img.shields.io/badge/go-documentation-blue.svg)](https://godoc.org/github.com/adidas/adictl)

**TODO: Change all reference of jorgechato repo to adidas**

## Getting Started & Documentation

If you're new to Acid and want to get started creating your own Jenkins as a service with kubernetes, please checkout our Getting Started guide.

Acictl must first be installed on your machine. Acictl is distributed as a binary package for all supported platforms and architectures. You can download it from [here](https://github.com/jorgechato/acictl/releases).

### Installing acictl

To install Acictl, find the appropriate package for your system and download it.

After downloading Acictl, rename the binary to acictl. Acictl runs as a single binary named acictl-[OS]-[arch].

The final step is to make sure that the acictl binary is available on the PATH. See [this page](https://stackoverflow.com/questions/14637979/how-to-permanently-set-path-on-linux-unix) for instructions on setting the PATH on Linux and Mac. [This page](https://stackoverflow.com/questions/1618280/where-can-i-set-path-to-make-exe-on-windows) contains instructions for setting the PATH on Windows.

### Verifying the Installation

After installing Acictl, verify the installation worked by opening a new terminal session and checking that acictl is available. By executing acictl you should see help output similar to this:

```bash
$ acictl
acictl v1.0 is a CLI to control the acid project.
This application is a tool to generate the needed files
to quickly create a Jenkins as a service environment in K8S.

Usage:
  acictl [flags]
  acictl [command]

Available Commands:
# ...
```

If you get an error that acictl could not be found, your PATH environment variable was not set up properly. Please go back and ensure that your PATH variable contains the directory where Acictl was installed.

### Developing Acictl

If you wish to work on Acictl itself, you'll first need Go installed on your machine (version 1.10+ is required). Alternatively, you can use the Vagrantfile in the root of this repo to stand up a virtual machine with the appropriate dev tooling already set up for you.

This repository contains only Acictl core, which includes the command line interface.

For local development of Acictl core, first make sure Go is properly installed and that a
[GOPATH](http://golang.org/doc/code.html#GOPATH) has been set. You will also need to add `$GOPATH/bin` to your `$PATH`.

Next, using [Git](https://git-scm.com/), clone this repository into `$GOPATH/src/github.com/jorgechato/acictl`.

You'll need to run `make tools` to install some required tools, then `make`.  This will compile the code and then run the tests. If this exits with exit status 0, then everything is working!
You only need to run `make tools` once (or when the tools change).

```sh
$ cd "$GOPATH/src/github.com/adidas/acictl"
$ make tools
$ make
```

If you're developing a specific package, you can run tests for just that package by specifying the `TEST` variable. For example below, only `generator` package tests will be run.

```sh
$ make test TEST=./generator
...
```
