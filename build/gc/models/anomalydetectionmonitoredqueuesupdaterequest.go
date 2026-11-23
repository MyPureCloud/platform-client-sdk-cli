package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalydetectionmonitoredqueuesupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalydetectionmonitoredqueuesupdaterequestDud struct { 
    

}

// Anomalydetectionmonitoredqueuesupdaterequest
type Anomalydetectionmonitoredqueuesupdaterequest struct { 
    // Operations - Operations to perform on queue monitoring configuration. Each operation can only be specified a maximum of once
    Operations []Anomalydetectionmonitoringqueueoperation `json:"operations"`

}

// String returns a JSON representation of the model
func (o *Anomalydetectionmonitoredqueuesupdaterequest) String() string {
     o.Operations = []Anomalydetectionmonitoringqueueoperation{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalydetectionmonitoredqueuesupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Anomalydetectionmonitoredqueuesupdaterequest

    if AnomalydetectionmonitoredqueuesupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    AnomalydetectionmonitoredqueuesupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Operations []Anomalydetectionmonitoringqueueoperation `json:"operations"`
        *Alias
    }{

        
        Operations: []Anomalydetectionmonitoringqueueoperation{{}},
        

        Alias: (*Alias)(u),
    })
}

