package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AftercallworkupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AftercallworkupdateDud struct { 
    

}

// Aftercallworkupdate
type Aftercallworkupdate struct { 
    // AfterCallWorkRequired - Indicates whether or not after-call work must be completed for the communication. Can only be updated for connected communications.
    AfterCallWorkRequired bool `json:"afterCallWorkRequired"`

}

// String returns a JSON representation of the model
func (o *Aftercallworkupdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aftercallworkupdate) MarshalJSON() ([]byte, error) {
    type Alias Aftercallworkupdate

    if AftercallworkupdateMarshalled {
        return []byte("{}"), nil
    }
    AftercallworkupdateMarshalled = true

    return json.Marshal(&struct {
        
        AfterCallWorkRequired bool `json:"afterCallWorkRequired"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

