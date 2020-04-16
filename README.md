[![Build Status](https://travis-ci.com/tsunejui/fops.svg?branch=master)](https://travis-ci.com/tsunejui/fops)

# fops
A command-line application `fops` written in Golang

![image](./docs/overview.gif)

## Menu
- [Features](#features)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Test](#test)

<a name="features"/>

## Features
- count line with the specific file
- generate checksum with the specific file

_**Note1:** use `spf13 / cobra` to implement command line utilities._

_**Note2:** use `gabriel-vasile/mimetype` to detect non-text files._
<a name="installation"/>

## Installation
```shell
go build -o fops
```

<a name="getting-started"/>

## Getting Started
```shell
# Print the line count of file, support both short and long options 
./fops linecount -f myfile.txt

# Print the checksum of file, support multiple algorithms: md5, sha1 and sha256 
./fops checksum -f myfile.txt --md5 
./fops checksum -f myfile.txt --sha1 
./fops checksum -f myfile.txt --sha256 

# Show version and help
./fops versio
./fops help
```

<a name="test"/>

## Test
```shell
go test ./...
```
![image](./docs/testing.PNG)
