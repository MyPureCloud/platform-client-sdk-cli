package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowloglevelrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowloglevelrequestDud struct { 
    

}

// Flowloglevelrequest - Used to set the log level of a particular flow
type Flowloglevelrequest struct { 
    // LogLevelCharacteristics - The log level characteristics currently set for this flow
    LogLevelCharacteristics Flowloglevel `json:"logLevelCharacteristics"`

}

// String returns a JSON representation of the model
func (o *Flowloglevelrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowloglevelrequest) MarshalJSON() ([]byte, error) {
    type Alias Flowloglevelrequest

    if FlowloglevelrequestMarshalled {
        return []byte("{}"), nil
    }
    FlowloglevelrequestMarshalled = true

    return json.Marshal(&struct {
        
        LogLevelCharacteristics Flowloglevel `json:"logLevelCharacteristics"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

