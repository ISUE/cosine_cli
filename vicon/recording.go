package vicon

import (
	"cosine_cli/utils"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type (
	DeviceName string
	Action     string
	Modality   string
)

var (
	DeviceNameNoteTenPlus1 DeviceName = "n10pl-1"
	DeviceNameNoteTenPlus2 DeviceName = "n10pl-2"

	AllDeviceNames = map[DeviceName]any{
		DeviceNameNoteTenPlus1: nil,
		DeviceNameNoteTenPlus2: nil,
	}
)

var (
	ActionFigure8 Action = "figure8"
	ActionRaise   Action = "raise"
	ActionSpin    Action = "spin"
	ActionWalking Action = "walking"
	ActionJumping Action = "jumping"

	ActionCoordinatedStep Action = "coordinatedstep"
	ActionLunge           Action = "lunge"
	ActionTwirl           Action = "twirl"
	ActionCrouch          Action = "crouch"

	StickActions = map[Action]any{
		// shared actions start
		ActionWalking:         nil,
		ActionJumping:         nil,
		ActionCoordinatedStep: nil,
		// shared actions end
		ActionFigure8: nil,
		ActionRaise:   nil,
		ActionSpin:    nil,
	}

	BodyActions = map[Action]any{
		// shared actions start
		ActionWalking:         nil,
		ActionJumping:         nil,
		ActionCoordinatedStep: nil,
		// shared actions end
		ActionCrouch: nil,
		ActionTwirl:  nil,
		ActionLunge:  nil,
	}

	AllActions = utils.MergeMaps(BodyActions, StickActions)
)

var (
	ModalityBody  Modality = "A"
	ModalityStick Modality = "B"

	AllModalities = map[Modality]any{ModalityBody: nil, ModalityStick: nil}
)

// Recording
// n10pl-1_1A_coordinated_step_user_32-1_2-28-23_output
// A = Body
// B = Stick
type Recording struct {
	StartTime      time.Time
	EndTime        time.Time
	DeviceName     DeviceName
	Action         Action
	Modality       Modality
	JostleStrength int
	UserID         int
}

func (t *Recording) String() string {
	return fmt.Sprintf(
		"{DeviceName=%s, Action=%s, Modality=%s, JostleStrength=%d, UserID=%d, StartTime=%s, EndTime=%s}",
		t.DeviceName,
		t.Action,
		t.Modality,
		t.JostleStrength,
		t.UserID,
		t.StartTime,
		t.EndTime,
	)
}

func NewRecording(xcpPath string) (*Recording, error) {
	absPath, err := filepath.Abs(xcpPath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cameras Cameras
	if err = xml.NewDecoder(file).Decode(&cameras); err != nil {
		return nil, err
	}

	startTime, err := ParseViconDataTime(cameras.Camera[0].Capture.STARTTIME)
	if err != nil {
		return nil, fmt.Errorf("unable to parse vicon start time in xcp file: %v", err)
	}

	endTime, err := ParseViconDataTime(cameras.Camera[0].Capture.ENDTIME)
	if err != nil {
		return nil, fmt.Errorf("unable to parse vicon end time in xcp file: %v", err)
	}

	t := &Recording{StartTime: startTime, EndTime: endTime}
	if err != nil {
		return nil, err

	}

	csvFileName := strings.Replace(filepath.Base(xcpPath), ".xcp", "_output.csv", 1)
	err = t.ScrapeFileName(csvFileName)
	if err != nil {
		return nil, fmt.Errorf("unable to scrape file name %s: %v", csvFileName, err)
	}

	return t, nil
}

// ScrapeFileName
// n10pl-1_1A_coordinated_step_user_32-1_2-28-23_output.csv
func (t *Recording) ScrapeFileName(csvFileName string) (err error) {
	parts := strings.Split(csvFileName, "_")
	if t.DeviceName, err = GetDeviceName(strings.Replace(parts[0], "-", "", 1)); err != nil {
		return err
	}

	modalityCombo := parts[1]

	jostleStrength, err := strconv.ParseInt(string(modalityCombo[0]), 10, 64)
	if err != nil {
		return err
	}
	t.JostleStrength = int(jostleStrength)

	if t.Modality, err = GetModality(string(modalityCombo[1])); err != nil {
		return err
	}

	if err != nil {
		return err
	}

	if t.Action, err = GetAction(parts[2]); err != nil {
		return err
	}

	userIdPart := parts[4]

	// remove the -1 from the user id
	userIdPartRemovedNum := userIdPart[0 : len(userIdPart)-2]
	userId, err := strconv.ParseInt(userIdPartRemovedNum, 10, 64)
	if err != nil {
		return err
	}

	t.UserID = int(userId)

	return nil
}

func GetDeviceName(name string) (DeviceName, error) {
	_, ok := AllDeviceNames[DeviceName(name)]

	if !ok {
		return "", fmt.Errorf("unable to find device name with name %s", name)
	}

	return DeviceName(name), nil
}

func GetAction(action string) (Action, error) {
	_, ok := AllActions[Action(action)]

	if !ok {
		return "", fmt.Errorf("unable to find an action with name %s", action)
	}

	return Action(action), nil
}

func GetModality(modality string) (Modality, error) {
	_, ok := AllModalities[Modality(modality)]

	if !ok {
		return "", fmt.Errorf("unable to find a modality with name %s", modality)
	}

	return Modality(modality), nil
}
