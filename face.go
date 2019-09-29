package main

import (
	"fmt"
	"image"
	_ "image/color"
	"os"
	"path/filepath"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"

	"gocv.io/x/gocv"
)

var (
	net    gocv.Net
	fcount int
)

// 얼굴인증 결과 처리
func performDetection(frame *gocv.Mat, results gocv.Mat) {
	for i := 0; i < results.Total(); i += 7 {
		confidence := results.GetFloatAt(0, i+2)
		if confidence > 0.4 {
			left := int(results.GetFloatAt(0, i+3) * float32(frame.Cols()))
			top := int(results.GetFloatAt(0, i+4) * float32(frame.Rows()))
			right := int(results.GetFloatAt(0, i+5) * float32(frame.Cols()))
			bottom := int(results.GetFloatAt(0, i+6) * float32(frame.Rows()))

			imgFace := frame.Region(image.Rect(left, top, right, bottom))
			fcount = fcount + 1
			gocv.GaussianBlur(imgFace, &imgFace, image.Pt(75, 75), 0, 0, gocv.BorderDefault)
			imgFace.Close()
		}
	}
}

func faceDetection(uid string) string {

	// 얼굴인식 초기값
	ratio := 1.0
	mean := gocv.NewScalar(104, 177, 123, 0)
	swapRGB := false

	// 업로드된 이미지를 읽어서...
	img := gocv.IMRead("./web/public/uploads/"+uid, gocv.ColorBGRAToBGR)
	if img.Empty() {
		return "Load Error"
	}
	blob := gocv.BlobFromImage(img, ratio, image.Pt(300, 300), mean, swapRGB, false)

	// 얼굴인식
	net.SetInput(blob, "")
	prob := net.Forward("")
	performDetection(&img, prob)

	prob.Close()
	blob.Close()

	// 결과를저장
	savename := "./web/public/results/" + uid + ".jpg"
	gocv.IMWrite(savename, img)

	return uid
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {

	// 폴더 없으면 만든다
	newpath1 := filepath.Join("./web/", "public")
	os.MkdirAll(newpath1, os.ModePerm)
	newpath2 := filepath.Join("./web/public", "uploads")
	os.MkdirAll(newpath2, os.ModePerm)
	newpath3 := filepath.Join("./web/public", "results")
	os.MkdirAll(newpath3, os.ModePerm)

	// 카페 모델 읽어들이기
	if !fileExists("res10_300x300_ssd_iter_140000.caffemodel") {
		fmt.Println("Face Detect Model Not Exist!")
		fmt.Println("Download : wget https://github.com/opencv/opencv_3rdparty/raw/dnn_samples_face_detector_20170830/res10_300x300_ssd_iter_140000.caffemodel")
		return
	}

	if !fileExists("deploy.prototxt") {
		fmt.Println("Face Detect Prototext Config Not Exist!")
		fmt.Println("Download : wget https://raw.githubusercontent.com/opencv/opencv/master/samples/dnn/face_detector/deploy.prototxt")
		return
	}

	// DNN 네트워크 생성
	net = gocv.ReadNet("res10_300x300_ssd_iter_140000.caffemodel", "deploy.prototxt")
	defer net.Close()

	net.SetPreferableBackend(gocv.NetBackendType(gocv.NetBackendDefault))
	net.SetPreferableTarget(gocv.NetTargetType(gocv.NetTargetCPU))

	// 웹 서버 생성
	r := gin.Default()
	r.MaxMultipartMemory = 1
	r.Use(static.Serve("/", static.LocalFile("./web", false)))

	// 이미지 업로드
	r.POST("/api/image/add", func(c *gin.Context) {
		file, _ := c.FormFile("imageFile")

		uid := xid.New().String()
		filename := "./web/public/uploads/" + uid
		c.SaveUploadedFile(file, filename)

		// 얼굴인식
		fcount = 0
		location := faceDetection(uid)
		c.JSON(http.StatusOK, gin.H{
			"location": location,
			"fcount":   fcount,
		})
	})

	r.Run(":8082")
}
