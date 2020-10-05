package main

import (
	"fmt"

	"github.com/illuminatly/godal"
)

func main() {

	//var ds gdal.Dataset
	imageList := []string{
		"/vsis3/jaic-dev-ra-rasterstore/jaicdev_gis_NOAA_Hurricane_Dorian_lkbpjiel/fef57e1fc4030a0c049e0c0bdff87415bd13d95fa6e0317fc9728719f4764e19.tif",
		"/vsis3/jaic-dev-ra-rasterstore/jaicdev_gis_NOAA_Hurricane_Dorian_lkbpjiel/ff684a4099155d03b26fccc2edf32371aa30efc624a9704123bb25a25a5e0840.tif",
		"/vsis3/jaic-dev-ra-rasterstore/jaicdev_gis_NOAA_Hurricane_Dorian_lkbpjiel/ff730ad627b3e7361ac610048e2fc5b4ea0731a127c1daeb4d2a20600040e1a0.tif",
		"/vsis3/jaic-dev-ra-rasterstore/jaicdev_gis_NOAA_Hurricane_Dorian_lkbpjiel/fff1281b2e63039339ce32188c71568be237d4a386e491268a9cb1bbcf1b7b79.tif",
		"/vsis3/jaic-dev-ra-rasterstore/jaicdev_gis_NOAA_Hurricane_Dorian_lkbpjiel/fffdaefa24ebaa16eafe9633f0007642dcd94f6a5b03b87c5bb9115be7e84a4a.tif",
		"/vsis3/jaic-dev-ra-rasterstore/jaicdev_gis_NOAA_Hurricane_Dorian_lkbpjiel/ffff76ea1eeb30ccbe89fc13ee1a2a9a5f1a322d93d883d2cd5c6caf59fb7283.tif",
	}

	options := []string{}

	outputFile := "/vsis3/jaic-dev-ra-rasterstore/jaicdev_gis_NOAA_Hurricane_Dorian_lkbpjiel/test.vrt"

	outputDs := gdal.BuildVRT(outputFile, imageList, options)

	fmt.Println(outputDs)

	outputDs.Close()

	//outputDs := gdal.GDALWarp(outputFile, gdal.Dataset{}, []gdal.Dataset{ds}, options)
	fmt.Printf("End program\n")
}
