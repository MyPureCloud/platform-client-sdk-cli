package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ParkingstaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ParkingstaterequestDud struct { 
    


    

}

// Parkingstaterequest
type Parkingstaterequest struct { 
    // State - State to set the participant.
    State string `json:"state"`


    // ResumeTime - Timestamp for resume parked conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ResumeTime time.Time `json:"resumeTime"`

}

// String returns a JSON representation of the model
func (o *Parkingstaterequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Parkingstaterequest) MarshalJSON() ([]byte, error) {
    type Alias Parkingstaterequest

    if ParkingstaterequestMarshalled {
        return []byte("{}"), nil
    }
    ParkingstaterequestMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        ResumeTime time.Time `json:"resumeTime"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

