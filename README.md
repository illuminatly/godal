## About


The gdal.go package provides a go wrapper for GDAL, the Geospatial Data Abstraction Library. More information about GDAL can be found at http://www.gdal.org
This repo was originally forked from `github.com/lukeroth/gdal` however it appears that most development has stopped. This
pulls in several unmerged PRs from that repo (including importantly gdal's VSI functionality) and continues to extend functionality.
                                     
## Status
[![CircleCI](https://circleci.com/gh/Rob-Fletcher/go-gdal.svg?style=shield)](https://circleci.com/gh/Rob-Fletcher/go-gdal)
[![Go Report Card](https://goreportcard.com/badge/github.com/Rob-Fletcher/go-gdal)](https://goreportcard.com/report/github.com/Rob-Fletcher/go-gdal)
## Installation

1) go get github.com/boundless/go-gdal
2) `brew install gdal` (on OSX)
3) go build 

## Compatibility

This software has been tested most recently on 1) OSX and 2) Ubuntu 19.04 Linux using GDAL 2.4.0 

## Examples

See the examples directory for some examples of low level usage.  The macro package has higher level utilities
