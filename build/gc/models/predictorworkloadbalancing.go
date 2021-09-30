package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictorworkloadbalancingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictorworkloadbalancingDud struct { 
    

}

// Predictorworkloadbalancing
type Predictorworkloadbalancing struct { 
    // Enabled - Flag to activate and deactivate workload balancing.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Predictorworkloadbalancing) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictorworkloadbalancing) MarshalJSON() ([]byte, error) {
    type Alias Predictorworkloadbalancing

    if PredictorworkloadbalancingMarshalled {
        return []byte("{}"), nil
    }
    PredictorworkloadbalancingMarshalled = true

    return json.Marshal(&struct { 
        Enabled bool `json:"enabled"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

