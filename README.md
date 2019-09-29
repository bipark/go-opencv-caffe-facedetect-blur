# go-opencv-caffe-facedetect-blur
Go &amp; OpenCV &amp; Caffe 기반 얼굴인식 &amp; 얼굴 블러 처리하기

# 이게 뭔가요?
OpenCV 공부중에 주말 프로젝트로 만들어본 Go 언어 기반 얼굴인식 기능 및 블러 처리 프로그램입니다.

이미 학습되어 있는 Caffe 기반의 딥러닝 얼굴인식 모델을 활용하여 주어진 이미지에서 얼굴을 인식하고 인식된 얼굴 부분을 블러 처리하여 웹요청에 응답합니다.

Go Gin 기반으로 웹 프론트에 용청에 응답을 처리 하고 프론트에서는 Vue.js & Axios 등으로 Go 서버에 이미지를 전송 하고 결과를 받아서 화면에 뿌려 줍니다.

이 소스는 MIT 라이센스로 자유롭게 활용할 수 있습니다.

# 데모페이지
아래 링크에서 실행 데모를 볼 수 있습니다.
[http://face.practical.kr](http://face.practical.kr)

# 의존성 
* Go 언어 - [https://golang.org](https://golang.org/)
	* gocv.io/x/gocv
	* github.com/gin-gonic/gin
	* github.com/rs/xid
	* github.com/gin-contrib/static
* Caffe Model - [https://github.com/opencv/opencv_3rdparty/raw/dnn_samples_face_detector_20170830/res10_300x300_ssd_iter_140000.caffemodel](https://github.com/opencv/opencv_3rdparty/raw/dnn_samples_face_detector_20170830/res10_300x300_ssd_iter_140000.caffemodel)
* Prototext Config - [https://raw.githubusercontent.com/opencv/opencv/master/samples/dnn/face_detector/deploy.prototxt](https://raw.githubusercontent.com/opencv/opencv/master/samples/dnn/face_detector/deploy.prototxt)
* Vue.js - [https://vuejs.org/](https://vuejs.org/)

# 실행방법

* Git Clone - [https://github.com/bipark/go-opencv-caffe-facedetect-blur.git](https://github.com/bipark/go-opencv-caffe-facedetect-blur.git)
* GoCV 설치 - [https://gocv.io/getting-started](https://gocv.io/getting-started)
* Caffe 모델 다운로드
	* Caffe Model - [https://github.com/opencv/opencv_3rdparty/raw/dnn_samples_face_detector_20170830/res10_300x300_ssd_iter_140000.caffemodel](https://github.com/opencv/opencv_3rdparty/raw/dnn_samples_face_detector_20170830/res10_300x300_ssd_iter_140000.caffemodel)
	* Prototext Config - [https://raw.githubusercontent.com/opencv/opencv/master/samples/dnn/face_detector/deploy.prototxt](https://raw.githubusercontent.com/opencv/opencv/master/samples/dnn/face_detector/deploy.prototxt)
