## About


The gdal.go package provides a go wrapper for GDAL, the Geospatial Data Abstraction Library. More information about GDAL can be found at http://www.gdal.org
This repo was originally forked from `github.com/lukeroth/gdal` however it appears that most development has stopped. This
pulls in several unmerged PRs from that repo (including importantly gdal's VSI functionality).
                                     

## Installation

1) go get github.com/boundless/go-gdal
2) `brew install gdal` (on OSX)
3) go build 

## Compatibility

This software has been tested most recently on OSX and Linux using GDAL 2.4.0

## Examples

See the examples directory for some examples of low level usage.  The macro package has higher level utilities