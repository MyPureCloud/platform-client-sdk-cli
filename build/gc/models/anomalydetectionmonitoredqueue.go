package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalydetectionmonitoredqueueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalydetectionmonitoredqueueDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Anomalydetectionmonitoredqueue
type Anomalydetectionmonitoredqueue struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // OverallAnomalyScore - System-determined numeric value representing how anomalous the queue was when last evaluated. Only present when the queue has been evaluated by the anomaly detection model
    OverallAnomalyScore float64 `json:"overallAnomalyScore"`


    // DateLastEvaluated - The timestamp when the queue was last evaluated for anomalies, in ISO-8601 format. Only present when the queue has been evaluated by the anomaly detection model
    DateLastEvaluated time.Time `json:"dateLastEvaluated"`


    // DateModified - The timestamp when this monitoring configuration was last updated, in ISO-8601 format
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Anomalydetectionmonitoredqueue) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalydetectionmonitoredqueue) MarshalJSON() ([]byte, error) {
    type Alias Anomalydetectionmonitoredqueue

    if AnomalydetectionmonitoredqueueMarshalled {
        return []byte("{}"), nil
    }
    AnomalydetectionmonitoredqueueMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        OverallAnomalyScore float64 `json:"overallAnomalyScore"`
        
        DateLastEvaluated time.Time `json:"dateLastEvaluated"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

