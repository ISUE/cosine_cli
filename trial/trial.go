package trial

import (
	"cosine_cli/phone"
	"cosine_cli/vicon"
	"fmt"
	"log"
	"math"
	"time"
)

type Trial struct {
	PhoneRecording phone.Recording
	ViconRecording vicon.Recording
}

func (t *Trial) String() string {
	return fmt.Sprintf("{PhoneRecording=%s, ViconRecording=%s}", t.PhoneRecording, t.ViconRecording)
}

func NewTrial(phoneRecording phone.Recording, viconRecording vicon.Recording) *Trial {
	trial := &Trial{PhoneRecording: phoneRecording, ViconRecording: viconRecording}

	return trial
}

type JointRecordings struct {
	PhoneRecordings []*phone.Recording
	ViconRecordings []*vicon.Recording
}

func (r *JointRecordings) String() string {
	return fmt.Sprintf("{PhoneRecordings=%s, ViconRecordings=%s}", r.PhoneRecordings, r.ViconRecordings)
}

func NewJointRecordings() *JointRecordings {
	return &JointRecordings{PhoneRecordings: nil, ViconRecordings: nil}
}

func (r *JointRecordings) AddViconRecording(recording *vicon.Recording) {
	if r.ViconRecordings == nil {
		r.ViconRecordings = make([]*vicon.Recording, 0)
	}
	r.ViconRecordings = append(r.ViconRecordings, recording)
}

func (r *JointRecordings) AddPhoneRecording(recording *phone.Recording) {
	if r.PhoneRecordings == nil {
		r.PhoneRecordings = make([]*phone.Recording, 0)
	}
	r.PhoneRecordings = append(r.PhoneRecordings, recording)
}

const RECORDING_TIME_THRESHOLD = time.Minute * 4
const RECORDING_CLOSENESS = time.Second * 10

func MatchRecordings(phoneRecordings []*phone.Recording, viconRecordings []*vicon.Recording) ([]*Trial, error) {

	fmt.Printf("phoneRecordings=%s, viconRecordings=%s\n", phoneRecordings, viconRecordings)

	jointRecordingsMap := make(map[int]*JointRecordings, 0)

	for _, recording := range phoneRecordings {
		jointRecordings, ok := jointRecordingsMap[recording.StartTime.YearDay()]
		if !ok {
			jointRecordings = NewJointRecordings()
			jointRecordingsMap[recording.StartTime.YearDay()] = jointRecordings
		}

		jointRecordings.AddPhoneRecording(recording)
	}

	for _, recording := range viconRecordings {
		jointRecordings, ok := jointRecordingsMap[recording.StartTime.YearDay()]
		if !ok {
			jointRecordings = NewJointRecordings()
			jointRecordingsMap[recording.StartTime.YearDay()] = jointRecordings
		}

		jointRecordings.AddViconRecording(recording)
	}

	trials := make([]*Trial, 0)

	for _, jointRecordings := range jointRecordingsMap {
		for _, viconRecording := range jointRecordings.ViconRecordings {
			var closest *phone.Recording = nil
			smallestDifference := math.MaxInt64
			for _, phoneRecording := range jointRecordings.PhoneRecordings {
				if closest == nil {
					closest = phoneRecording
					continue
				}

				viconPhoneDifference := int(math.Abs(float64(phoneRecording.StartTime.Unix() - viconRecording.StartTime.Unix())))
				if viconPhoneDifference < smallestDifference {
					closest = phoneRecording
					smallestDifference = viconPhoneDifference
				}
			}

			if smallestDifference > 120 {
				log.Printf("unable to find matching recording for vicon recording: %s, within 60 seconds of nearest phone recording. Closest phone recording is at %d", viconRecording, smallestDifference)
				continue
			}

			trials = append(trials, &Trial{
				PhoneRecording: *closest,
				ViconRecording: *viconRecording,
			})
		}
	}

	return trials, nil
}
