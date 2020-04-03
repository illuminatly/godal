package main

import (
	"fmt"
	"github.com/Rob-Fletcher/go-gdal"
	"log"
)

func main() {

    options := []string{
        "-te",
        "-80.82471",
        "28.616317",
        "-80.80688",
        "28.632901",
        "-tr",
        "5e-6",
        "5e-6",
        "-r",
        "near",
        "-overwrite",
    }

    fmt.Println(options)

    inputDatasets := make([]gdal.Dataset, 4)

    //var ds gdal.Dataset
    imageList := []string{
        "/data/projects/jaic/20190905aC0804930w283730n.tif",
        "/data/projects/jaic/20190905aC0804930w283900n.tif",
        "/data/projects/jaic/20190905bC0804930w283730n.tif",
        "/data/projects/jaic/20190905bC0804930w283900n.tif",
    }

    for i := 0; i < len(imageList); i++ {
        ds, err := gdal.Open(imageList[i], gdal.ReadOnly)
        if err != nil {
		    log.Fatal(err)
	    }

        inputDatasets[i] = ds
	    defer ds.Close()
    }

    outputFile := "/data/projects/jaic/test_out.tif"

	outputDs := gdal.GDALWarp(outputFile, gdal.Dataset{}, inputDatasets, options)
	defer outputDs.Close()

	fmt.Printf("End program\n")
}
