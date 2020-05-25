package macro

import (
	"testing"

	"github.com/Rob-Fletcher/go-gdal"
)

const S3_BUCKET_URL = "/vsis3/imagerytest"
const S3_FILE = "/vsis3/imagerytest/m_3707501_ne_18_1_20150801.tif"

func testMosaic(t *testing.T) {
	type args struct {
		urls []string
		epsg string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "everything", args: args{urls: gdal.VSIReadDirRecursiveAbs(S3_BUCKET_URL), epsg: "EPSG:4326"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mosaic(tt.args.urls, tt.args.epsg)
		})
	}
}
