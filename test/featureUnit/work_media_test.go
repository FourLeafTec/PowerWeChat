package featureUnit

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/media"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/media/response"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func Test_Media_Upload_Image(t *testing.T) {

	var wxResponse response2.ResponseUploadImage
	homePath, _ := os.UserHomeDir()
	response, _ := Work.Media.UploadImage(homePath+"/Desktop/123.jpg", nil)

	if response == nil {
		t.Error("response nil")
	} else if wxResponse.ErrCode != 0 {
		t.Error("response error uniformMessage as :", wxResponse.ErrMSG)
	}

	fmt.Dump(wxResponse, response)

}

func Test_Media_Upload_Temp_Image(t *testing.T) {

	var wxResponse *response2.ResponseUploadMedia
	homePath, _ := os.UserHomeDir()
	response, _ := Work.Media.UploadTempImage(homePath+"/Desktop/123.jpg", nil)

	if response == nil {
		t.Error("response nil")
	} else if wxResponse.ErrCode != 0 {
		t.Error("response error uniformMessage as :", wxResponse.ErrMSG)
	}

	fmt.Dump(wxResponse, response)

}

func Test_Media_Get(t *testing.T) {
	//respContract, _ := Work.Media.Get("3i12vSikgRGsE9Xbs3e1CZH4b_X6cREZtuHe2nh8VZ4w")
	respContract, _ := Work.Media.Get("3w91f_bN-iVUv1dbR4pLSBzFyhrVaF_G8lYEq-dk02ffr8F2lbYKeWuMVs1VeDTVB")

	//if response == nil {
	//	t.Error("response nil")
	//}
	resp := respContract.(*http.Response)
	defer resp.Body.Close()

	imageBuffer, err := ioutil.ReadAll(resp.GetBody())

	homePath, _ := os.UserHomeDir()
	imgPath := homePath + "/Desktop/getMediaImage.jpg"

	var opts jpeg.Options
	opts.Quality = 50

	media.SaveImage(imageBuffer, imgPath, opts)

	fmt.Dump("Test:", resp.GetBody(), err)
}
