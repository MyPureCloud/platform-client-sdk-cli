package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingappointmentstatusrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingappointmentstatusrequestDud struct { 
    

}

// Coachingappointmentstatusrequest
type Coachingappointmentstatusrequest struct { 
    // Status - The status of the coaching appointment
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Coachingappointmentstatusrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingappointmentstatusrequest) MarshalJSON() ([]byte, error) {
    type Alias Coachingappointmentstatusrequest

    if CoachingappointmentstatusrequestMarshalled {
        return []byte("{}"), nil
    }
    CoachingappointmentstatusrequestMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

