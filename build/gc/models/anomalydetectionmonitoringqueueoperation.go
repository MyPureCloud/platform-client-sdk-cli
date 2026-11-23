package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalydetectionmonitoringqueueoperationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalydetectionmonitoringqueueoperationDud struct { 
    


    

}

// Anomalydetectionmonitoringqueueoperation
type Anomalydetectionmonitoringqueueoperation struct { 
    // Operation - Operation to perform on the specified queues
    Operation string `json:"operation"`


    // QueueIds - The IDs of the queues to which the operation will be applied
    QueueIds []string `json:"queueIds"`

}

// String returns a JSON representation of the model
func (o *Anomalydetectionmonitoringqueueoperation) String() string {
    
     o.QueueIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalydetectionmonitoringqueueoperation) MarshalJSON() ([]byte, error) {
    type Alias Anomalydetectionmonitoringqueueoperation

    if AnomalydetectionmonitoringqueueoperationMarshalled {
        return []byte("{}"), nil
    }
    AnomalydetectionmonitoringqueueoperationMarshalled = true

    return json.Marshal(&struct {
        
        Operation string `json:"operation"`
        
        QueueIds []string `json:"queueIds"`
        *Alias
    }{

        


        
        QueueIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

