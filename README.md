<p style="width: 100%; padding: 100px; text-align: center;">
  <img src="./.github/images/logo.png" width="250" align="right" alt="HTML Templar"/>
</p>

# AciCTL
[![Build Status](https://travis-ci.com/adidas/acictl.svg?token=x3vLcsQVEzf1kfJyx1Uv&branch=master)](https://travis-ci.com/adidas/acictl)
[![Build Status](https://sonarcloud.io/api/project_badges/measure?project=com.adidas.acictl&metric=coverage)](https://sonarcloud.io/dashboard?id=com.adidas.acictl)
[![Build Status](https://sonarcloud.io/api/project_badges/measure?project=com.adidas.acictl&metric=alert_status)](https://sonarcloud.io/dashboard?id=com.adidas.acictl)
[![Go Report Card](https://goreportcard.com/badge/github.com/adidas/acictl)](https://goreportcard.com/report/github.com/adidas/acictl)
[![Godoc](https://img.shields.io/badge/go-documentation-blue.svg)](https://godoc.org/github.com/adidas/acictl)


## Getting Started & Documentation

If you're new to Acid and want to get started creating your own Jenkins as a service with kubernetes, please checkout our Getting Started guide.

Acictl must first be installed on your machine. Acictl is distributed as a binary package for all supported platforms and architectures. You can download it from [here](https://github.com/adidas/acictl/releases).

### Installing acictl

To install Acictl, find the appropriate package for your system and download it.

After downloading Acictl, rename the binary to `acictl` and if you are in UNIX systems and it is needed add the required permissions. Acictl runs as a single binary named acictl-[OS]-[arch].

```bash
$ chmod +x acictl
```

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

You can verify the SHA256 hashes easily by placing the files in the same folder as the download artifact and then running `shasum`, which is included in most UNIX-like systems:

```bash
$ ls
acictl-darwin-386        acictl-darwin-386.sha256
$ shasum -a 256 -c acictl-darwin-386.sha256
acictl-darwin-386: OK
```

### Developing Acictl

If you wish to work on Acictl itself, you'll first need Go installed on your machine (version 1.10+ is required). Alternatively, you can use the Vagrantfile in the root of this repo to stand up a virtual machine with the appropriate dev tooling already set up for you.

This repository contains only Acictl core, which includes the command line interface.

For local development of Acictl core, first make sure Go is properly installed and that a
[GOPATH](http://golang.org/doc/code.html#GOPATH) has been set. You will also need to add `$GOPATH/bin` to your `$PATH`.

Next, using [Git](https://git-scm.com/), clone this repository into `$GOPATH/src/github.com/adidas/acictl`.

You'll need to run `make setup` to install some required tools, then `make dev`.  This will compile the code and then run the tests. If this exits with exit status 0, then everything is working!
You only need to run `make server` once (or when the tools change).

```sh
$ cd "$GOPATH/src/github.com/adidas/acictl"
$ make setup
$ make dep
$ make server
```

If you're developing a specific package, you can run tests for just that package by specifying the `TEST` variable. For example below, only `generator` package tests will be run.

```sh
$ make unit TEST=./generator
...
```

### Dependencies

Acictl stores its dependencies under `vendor/`, which [Go 1.6+ will automatically recognize and load](https://golang.org/cmd/go/#hdr-Vendor_Directories). We use [`dep`](https://github.com/golang/dep) to manage the dep dependencies.

If you're developing Acictl, there are a few tasks you might need to perform.

#### Adding a dependency

If you're adding a dependency, you'll need to vendor it in the same Pull Request as the code that depends on it.

To add a dependency:

Assuming your work is on a branch called `my-feature-branch`, the steps look like this:

1. Add the new package to your `vendor/` directory:

    ```bash
    $ dep ensure -add github.com/foo/bar
    ```

2. Review the changes in git and commit them.

#### Updating a dependency

To update a dependency:

1. Fetch the dependency:

    ```bash
    $ dep ensure -update github.com/foo/bar
    ```

2. Review the changes in git and commit them.

## FAQ

### Maintainers

Check the contributor list and you will be welcome if you want to contribute.

### Contributing

Check out the [CONTRIBUTING.md](.github/CONTRIBUTING.md) file to know how to contribute to this project.

## License and Software Information

© adidas AG

adidas AG publishes this software and accompanied documentation (if any) subject to the terms of the MIT license with the aim of helping the community with our tools and libraries which we think can be also useful for other people. You will find a copy of the MIT license in the root folder of this package. All rights not explicitly granted to you under the MIT license remain the sole and exclusive property of adidas AG.

If you want to contact adidas regarding the software, you can mail us at _software.engineering@adidas.com_.

For further information open the [adidas terms and conditions](https://github.com/adidas/adidas-contribution-guidelines/wiki/Terms-and-conditions) page.

### License

[MIT](LICENSE)
