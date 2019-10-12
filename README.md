<p align="center"><img src="images/pushpanel-logo.png" alt="pushpanel"></p>

# pp - PushPanel CLI

[![GoDoc](https://godoc.org/github.com/spagu/pp?status.svg)](https://godoc.org/github.com/spagu/pp)&nbsp; 
[![Build Status](https://travis-ci.org/spagu/pp.svg?branch=master)](https://travis-ci.org/spagu/pp)&nbsp;
[![Go Report Card](https://goreportcard.com/badge/github.com/spagu/pp)](https://goreportcard.com/report/github.com/spagu/pp)&nbsp;
[![Coverage Status](https://coveralls.io/repos/github/spagu/pp/badge.svg?branch=master)](https://coveralls.io/github/spagu/pp?branch=master)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3267/badge)](https://bestpractices.coreinfrastructure.org/projects/3267)&nbsp;
<a href="https://github.com/spagu/pp/releases"><img src="https://img.shields.io/github/downloads/spagu/pp/total.svg" alt="Downloads"></a>&nbsp;
<a href="https://github.com/spagu/pp/releases"><img src="https://img.shields.io/github/release/spagu/pp.svg?label=Release" alt="Release"></a>


`pp` is a command line interface for [PushPanel.io](https://pushpanel.io) services. 

---

## Requirements
- Local Linux/BSD shell access
- Account on [PushPanel.io](https://pushpanel.io)
- Unblock Firewall on port 443 so your console is able to communicate with dash.pushpanel.io:443
- [GO](https://golang.org/) language installed 

## Install
```sh
go get github.com/spagu/pp
#or
git clone https://github.com/spagu/pp.git
cd pp && make install
```

## Configuration
Create file `.pushpanel` in your home directory with a Token from Panel

## Usage
Commands are controlled via server.

### Command-Line 

```
pp project list
pp project create projectname repotype projecttype --initialversion=
pp project release projectname
pp project revision projectname
pp project unpublish projectname
pp project archive projectname
pp website create domainname
pp website deploy development to production domainname
```

### Project management
Add file called `pushpanel.yaml` inside your repository with a controlled content:

```yaml
projects:

    name: supername

website:
    production:
        domainname:domanname.co.uk
    staging:
        domainname:stg.domanname.co.uk
    development:
        domainname:dev.domanname.co.uk
```
### How to help

Did you find any bugs or have some suggestions?
- Feel free to open [new issue](https://github.com/spagu/pp/issues/new).

Do you want to contribute to the project?
- Fork the repository and open a pull request. [Here](https://github.com/spagu/pp/projects/1) you can find TODO features.

---
## Author
`pp` was developed by [PushPanel](https://pushpanel.io).

## License
`pp` is available under the [BSD](./LICENSE) license.


