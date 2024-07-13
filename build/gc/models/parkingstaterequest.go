package models
import (
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
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

