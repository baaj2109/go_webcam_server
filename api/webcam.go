package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

var (
	Fps          float32 = 0
	DeviceID     int     = -1
	Cam          *gocv.VideoCapture
	closeChannel = make(chan bool)
	Window       *gocv.Window
	DeviceCount  int = 0
)

func init() {
	for {
		cam, err := gocv.OpenVideoCapture(DeviceCount)
		if nil != err {
			break
		}
		DeviceCount++
		cam.Close()
	}
	fmt.Println("[webcam] gocv got " + strconv.Itoa(DeviceCount) + " devices.")
}

func ListAllCamera(c *gin.Context) {
	currentDeviceId := 0
	for {
		cam, err := gocv.OpenVideoCapture(currentDeviceId)
		if nil != err {
			break
		}
		currentDeviceId++
		cam.Close()
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "device count : " + strconv.Itoa(currentDeviceId),
	})
}

func SelectCamera(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("webcam"))
	if id < 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "webcam id should greater than or eqaul to zero, or " + err.Error(),
		})
	}
	DeviceID = id
	c.JSON(http.StatusOK, gin.H{
		"message": "set webcam source : " + strconv.Itoa(DeviceID),
	})
}

func StopWebCam(c *gin.Context) {
	closeChannel <- true
}

func StartWebCam(c *gin.Context) {

	// deviceID := 1
	if DeviceID == -1 {
		if DeviceCount == 1 {
			DeviceID = 0
		} else if DeviceCount == 2 {
			DeviceID = 1
		}
	}
	cam, err := gocv.OpenVideoCapture(DeviceID)
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer cam.Close()

	Window := gocv.NewWindow("gocv window")
	defer Window.Close()

	frame := gocv.NewMat()
	defer frame.Close()
	timer := time.NewTicker(time.Second)

	for {
		select {
		case <-closeChannel:
			return
		case <-timer.C:
			fmt.Println("fps :", Fps)
			Fps = 0
		default:
			if ok := cam.Read(&frame); !ok {
				// fmt.Printf("cannot read device %v\n", deviceID)
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}
			if frame.Empty() {
				continue
			}
			Window.IMShow(frame)
			Window.WaitKey(1)
			Fps++
		}
	}
}
