package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalydetectionmonitoredqueuelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalydetectionmonitoredqueuelistingDud struct { 
    


    

}

// Anomalydetectionmonitoredqueuelisting
type Anomalydetectionmonitoredqueuelisting struct { 
    // Entities
    Entities []Anomalydetectionmonitoredqueue `json:"entities"`


    // TotalQueuesMonitored - Total count of queues monitored across all divisions in the organization
    TotalQueuesMonitored int `json:"totalQueuesMonitored"`

}

// String returns a JSON representation of the model
func (o *Anomalydetectionmonitoredqueuelisting) String() string {
     o.Entities = []Anomalydetectionmonitoredqueue{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalydetectionmonitoredqueuelisting) MarshalJSON() ([]byte, error) {
    type Alias Anomalydetectionmonitoredqueuelisting

    if AnomalydetectionmonitoredqueuelistingMarshalled {
        return []byte("{}"), nil
    }
    AnomalydetectionmonitoredqueuelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Anomalydetectionmonitoredqueue `json:"entities"`
        
        TotalQueuesMonitored int `json:"totalQueuesMonitored"`
        *Alias
    }{

        
        Entities: []Anomalydetectionmonitoredqueue{{}},
        


        

        Alias: (*Alias)(u),
    })
}

