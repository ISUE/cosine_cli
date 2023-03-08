package vicon

import (
	"encoding/xml"
	"time"
)

type Cameras struct {
	XMLName xml.Name `xml:"Cameras"`
	Text    string   `xml:",chardata"`
	SOURCE  string   `xml:"SOURCE,attr"`
	VERSION string   `xml:"VERSION,attr"`
	Camera  []struct {
		Text             string `xml:",chardata"`
		ACTIVETHRESHOLD  string `xml:"ACTIVE_THRESHOLD,attr"`
		DEVICEID         string `xml:"DEVICEID,attr"`
		DISPLAYTYPE      string `xml:"DISPLAY_TYPE,attr"`
		ISDV             string `xml:"ISDV,attr"`
		NAME             string `xml:"NAME,attr"`
		PIXELASPECTRATIO string `xml:"PIXEL_ASPECT_RATIO,attr"`
		SENSOR           string `xml:"SENSOR,attr"`
		SENSORSIZE       string `xml:"SENSOR_SIZE,attr"`
		SYSTEM           string `xml:"SYSTEM,attr"`
		TYPE             string `xml:"TYPE,attr"`
		USERID           string `xml:"USERID,attr"`
		ThresholdGrid    struct {
			Text     string `xml:",chardata"`
			BITDEPTH string `xml:"BIT_DEPTH,attr"`
			DATA     string `xml:"DATA,attr"`
			GRIDSIZE string `xml:"GRID_SIZE,attr"`
			TILESIZE string `xml:"TILE_SIZE,attr"`
		} `xml:"ThresholdGrid"`
		Calibration struct {
			Text      string `xml:",chardata"`
			ENDTEMP   string `xml:"END_TEMP,attr"`
			ENDTIME   string `xml:"END_TIME,attr"`
			ID        string `xml:"ID,attr"`
			STARTTEMP string `xml:"START_TEMP,attr"`
			STARTTIME string `xml:"START_TIME,attr"`
			TYPE      string `xml:"TYPE,attr"`
		} `xml:"Calibration"`
		IntrinsicsCalibration struct {
			Text      string `xml:",chardata"`
			ENDTEMP   string `xml:"END_TEMP,attr"`
			ENDTIME   string `xml:"END_TIME,attr"`
			ID        string `xml:"ID,attr"`
			STARTTEMP string `xml:"START_TEMP,attr"`
			STARTTIME string `xml:"START_TIME,attr"`
			TYPE      string `xml:"TYPE,attr"`
		} `xml:"IntrinsicsCalibration"`
		Capture struct {
			Text      string `xml:",chardata"`
			ENDTEMP   string `xml:"END_TEMP,attr"`
			ENDTIME   string `xml:"END_TIME,attr"`
			ID        string `xml:"ID,attr"`
			STARTTEMP string `xml:"START_TEMP,attr"`
			STARTTIME string `xml:"START_TIME,attr"`
		} `xml:"Capture"`
		ControlFrames string `xml:"ControlFrames"`
		KeyFrames     struct {
			Text     string `xml:",chardata"`
			KeyFrame struct {
				Text           string `xml:",chardata"`
				FOCALLENGTH    string `xml:"FOCAL_LENGTH,attr"`
				FRAME          string `xml:"FRAME,attr"`
				IMAGEERROR     string `xml:"IMAGE_ERROR,attr"`
				ORIENTATION    string `xml:"ORIENTATION,attr"`
				POSITION       string `xml:"POSITION,attr"`
				PRINCIPALPOINT string `xml:"PRINCIPAL_POINT,attr"`
				VICONRADIAL2   string `xml:"VICON_RADIAL2,attr"`
				WORLDERROR     string `xml:"WORLD_ERROR,attr"`
			} `xml:"KeyFrame"`
		} `xml:"KeyFrames"`
	} `xml:"Camera"`
}

const VICON_TIME = "2006-01-02 15:04:05"

// ParseViconDataTime
// Example Time: 2023-02-28 15:52:45.126
func ParseViconDataTime(timeString string) (time.Time, error) {
	timeString = timeString[0 : len(timeString)-3]
	return time.Parse(VICON_TIME, timeString)
}
