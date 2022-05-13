package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingappointmentstatusresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingappointmentstatusresponseDud struct { 
    Appointment Coachingappointmentreference `json:"appointment"`


    CreatedBy Userreference `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    Status string `json:"status"`

}

// Coachingappointmentstatusresponse
type Coachingappointmentstatusresponse struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Coachingappointmentstatusresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingappointmentstatusresponse) MarshalJSON() ([]byte, error) {
    type Alias Coachingappointmentstatusresponse

    if CoachingappointmentstatusresponseMarshalled {
        return []byte("{}"), nil
    }
    CoachingappointmentstatusresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

