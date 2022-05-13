package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgerebootparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgerebootparametersDud struct { 
    

}

// Edgerebootparameters
type Edgerebootparameters struct { 
    // CallDrainingWaitTimeSeconds - The number of seconds to wait for call draining to complete before initiating the reboot. A value of 0 will prevent call draining and all calls will disconnect immediately.
    CallDrainingWaitTimeSeconds int `json:"callDrainingWaitTimeSeconds"`

}

// String returns a JSON representation of the model
func (o *Edgerebootparameters) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgerebootparameters) MarshalJSON() ([]byte, error) {
    type Alias Edgerebootparameters

    if EdgerebootparametersMarshalled {
        return []byte("{}"), nil
    }
    EdgerebootparametersMarshalled = true

    return json.Marshal(&struct {
        
        CallDrainingWaitTimeSeconds int `json:"callDrainingWaitTimeSeconds"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

