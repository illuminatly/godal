package gdal

/*
#include "go_gdal.h"
#include "gdal_version.h"
#include "gdal_utils.h"

#cgo linux  pkg-config: gdal
#cgo darwin pkg-config: gdal
#cgo windows LDFLAGS: -Lc:/gdal/release-1600-x64/lib -lgdal_i
#cgo windows CFLAGS: -IC:/gdal/release-1600-x64/include
*/
import "C"
import (
	"fmt"
	"unsafe"
)

var _ = fmt.Println

/* --------------------------------------------- */
/* GDAL utilities                                */
/* --------------------------------------------- */

//GDALTranslateOptions holds options to be passed to gdal translated
type GDALTranslateOptions struct {
	cval *C.GDALTranslateOptions
}

//GDALWarpAppOptions holds options to be passed to gdal translated
type GDALWarpAppOptions struct {
	cval *C.GDALWarpAppOptions
}

//GDALTranslate is a utility to convert images into different formats
func GDALTranslate(
	destName string,
	srcDS Dataset,
	options []string,
) Dataset {

	var err C.int

	length := len(options)
	cOptions := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	gdalTranslateOptions := GDALTranslateOptions{C.GDALTranslateOptionsNew((**C.char)(unsafe.Pointer(&cOptions[0])), nil)}

	outputDs := C.GDALTranslate(
		C.CString(destName),
		srcDS.cval,
		gdalTranslateOptions.cval,
		&err,
	)

	return Dataset{outputDs}

}

//GDALWarp is a utility to warp images into different projections
func GDALWarp(
	destName string,
	dstDs Dataset,
	srcDs []Dataset,
	options []string,
) (Dataset, error) {

	var err C.int

	length := len(options)
	cOptions := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	gdalWarpOptions := GDALWarpAppOptions{C.GDALWarpAppOptionsNew((**C.char)(unsafe.Pointer(&cOptions[0])), nil)}
	if gdalWarpOptions.cval == nil {
		fmt.Println("Warp options value is nil")
	}

	pahSrcDs := make([]C.GDALDatasetH, len(srcDs)+1)
	for i := 0; i < len(srcDs); i++ {
		pahSrcDs[i] = srcDs[i].cval
	}
	pahSrcDs[len(srcDs)] = (C.GDALDatasetH)(unsafe.Pointer(nil))

	outputDs := C.GDALWarp(
		C.CString(destName),
		dstDs.cval,
		C.int(len(srcDs)),
		(*C.GDALDatasetH)(unsafe.Pointer(&pahSrcDs[0])),
		gdalWarpOptions.cval,
		&err,
	)
	fmt.Println("Error code from warp is:")
	fmt.Println(err)

	if err != 0 {
		return Dataset{outputDs}, ErrFailure
	}

	return Dataset{outputDs}, nil

}
