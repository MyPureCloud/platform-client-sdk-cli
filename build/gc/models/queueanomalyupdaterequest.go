package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueueanomalyupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueueanomalyupdaterequestDud struct { 
    

}

// Queueanomalyupdaterequest
type Queueanomalyupdaterequest struct { 
    // FalsePositive - Indicates whether the anomaly has been flagged as a false positive
    FalsePositive bool `json:"falsePositive"`

}

// String returns a JSON representation of the model
func (o *Queueanomalyupdaterequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queueanomalyupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Queueanomalyupdaterequest

    if QueueanomalyupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    QueueanomalyupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        FalsePositive bool `json:"falsePositive"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

