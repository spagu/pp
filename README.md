# pp - PushPanel CLI
[![GoDoc](https://godoc.org/github.com/spagu/pp?status.svg)](https://godoc.org/github.com/spagu/pp)&nbsp; 
[![Build Status](https://travis-ci.org/spagu/pp.svg?branch=master)](https://travis-ci.org/spagu/pp)&nbsp;
[![Go Report Card](https://goreportcard.com/badge/github.com/spagu/pp)](https://goreportcard.com/report/github.com/spagu/pp)&nbsp;
[![Coverage Status](https://coveralls.io/repos/github/spagu/pp/badge.svg?branch=master)](https://coveralls.io/github/spagu/pp?branch=master)

`pp` is a command line interface for PushPanel services. 

## Requirments
1. Local Linux/FreeBSD shell access
2. Account on PushPanel.io
3. Firewall able to comuniacte with dash.pushpanel.io on port 443
4. GO language installed 

## Install
```sh
go get github.com/spagu/pp
```

## Configuration
Create file .pushpanel in your home directory with a Token from Panel

## Usage
Commands are controlled via server.

### Command-Line 

pp project list
pp project create projectname repotype projecttype --initialversion=
pp project release projectname
pp project revision projectname
pp project unpublish projectname
pp project archive projectname
pp website create domainname
pp website deploy development to production domainname

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

## Author
`pp` was developed by [PushPanel](https://pushpanel.io).

## License
`pp` is available under the [BSD](./LICENSE) license.


