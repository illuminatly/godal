package main

import (
	"fmt"
	"github.com/Rob-Fletcher/go-gdal"
)

func main() {

    extent := [4]float64{
        -80.82471,
        28.616317,
        -80.80688,
        28.632901,
    }

    vrtOptions := gdal.GoBuildVRTOptions{
        Extent: extent,
        XRes: 5e-6,
        YRes: 5e-6,
        Resample: gdal.GRA_NearestNeighbour,
    }
    fmt.Println(vrtOptions)

    //var ds gdal.Dataset
    imageList := []string{
        "/data/projects/jaic/20190905aC0804930w283730n.tif",
        "/data/projects/jaic/20190905aC0804930w283900n.tif",
        "/data/projects/jaic/20190905bC0804930w283730n.tif",
        "/data/projects/jaic/20190905bC0804930w283900n.tif",
    }

    outputFile := "/data/projects/jaic/test_out.tif"

    outputDs := gdal.BuildVRT(outputFile, vrtOptions, imageList)


    fmt.Println(outputDs)

    outputDs.Close()

	//outputDs := gdal.GDALWarp(outputFile, gdal.Dataset{}, []gdal.Dataset{ds}, options)
	fmt.Printf("End program\n")
}
