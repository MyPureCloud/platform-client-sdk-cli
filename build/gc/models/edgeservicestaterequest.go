package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgeservicestaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgeservicestaterequestDud struct { 
    


    

}

// Edgeservicestaterequest
type Edgeservicestaterequest struct { 
    // InService - A boolean that sets the Edge in-service or out-of-service.
    InService bool `json:"inService"`


    // CallDrainingWaitTimeSeconds - The number of seconds to wait for call draining to complete before initiating the reboot. A value of 0 will prevent call draining and all calls will disconnect immediately.
    CallDrainingWaitTimeSeconds int `json:"callDrainingWaitTimeSeconds"`

}

// String returns a JSON representation of the model
func (o *Edgeservicestaterequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgeservicestaterequest) MarshalJSON() ([]byte, error) {
    type Alias Edgeservicestaterequest

    if EdgeservicestaterequestMarshalled {
        return []byte("{}"), nil
    }
    EdgeservicestaterequestMarshalled = true

    return json.Marshal(&struct {
        
        InService bool `json:"inService"`
        
        CallDrainingWaitTimeSeconds int `json:"callDrainingWaitTimeSeconds"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

