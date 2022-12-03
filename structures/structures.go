package structures

import (
	. "finalWork/pkg/email"
	. "finalWork/pkg/incident"
	. "finalWork/pkg/mms"
	. "finalWork/pkg/sms"
	. "finalWork/pkg/support"
	. "finalWork/pkg/voiceCall"
)

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}

func GetResultData() ResultSetT {

	result := ResultSetT{
		SMS:       GetResultSms(),
		MMS:       GetResultMMS(),
		VoiceCall: GetVoiceCallData(),
		Email:     GetResultEmail(),
		Support:   GetResultSupport(),
		Incidents: GetResultSIncident(),
	}
	return result

}
