package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalydetectionavailablequeueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalydetectionavailablequeueDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Anomalydetectionavailablequeue
type Anomalydetectionavailablequeue struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // AverageInteractionsPerDay - Average number of interactions offered per day for this queue over the past 90 days, rounded to the nearest whole number
    AverageInteractionsPerDay int `json:"averageInteractionsPerDay"`


    // Allowed - Indicates whether the queue is eligible for anomaly detection monitoring
    Allowed bool `json:"allowed"`


    

}

// String returns a JSON representation of the model
func (o *Anomalydetectionavailablequeue) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalydetectionavailablequeue) MarshalJSON() ([]byte, error) {
    type Alias Anomalydetectionavailablequeue

    if AnomalydetectionavailablequeueMarshalled {
        return []byte("{}"), nil
    }
    AnomalydetectionavailablequeueMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        AverageInteractionsPerDay int `json:"averageInteractionsPerDay"`
        
        Allowed bool `json:"allowed"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

