## About

The gdal.go package provides a go wrapper for GDAL, the Geospatial Data Abstraction Library. More information about GDAL can be found at http://www.gdal.org

This repo was originally forked from `github.com/lukeroth/gdal` and also `github.com/Rob-Fletcher` and will be developed further to closer mirror idiomatic Go.

It was also updated to support gdal 3.x, while the older version wouldn't work on later than 2.4. 
                                     
## Installation

1) go get github.com/illuminatly/godal
2) install gdal 3.x+ (including headers)
3) go build 

## Compatibility

This software has been tested most recently on Ubuntu 20.04 with gdal 3.1

## Examples

See the examples directory for some examples of low level usage.  The macro package has higher level utilities
