package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchpredictorrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchpredictorrequestDud struct { 
    


    


    

}

// Patchpredictorrequest
type Patchpredictorrequest struct { 
    // RoutingTimeoutSeconds - Number of seconds allocated to predictive routing before attempting a different routing method. This is a value between 12 and 900 seconds.
    RoutingTimeoutSeconds int `json:"routingTimeoutSeconds"`


    // Schedule - The predictor schedule that determines when the predictor is used for routing interactions.
    Schedule Predictorschedule `json:"schedule"`


    // WorkloadBalancingConfig - The predictor balancing configuration to enable workload balancing
    WorkloadBalancingConfig Predictorworkloadbalancing `json:"workloadBalancingConfig"`

}

// String returns a JSON representation of the model
func (o *Patchpredictorrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchpredictorrequest) MarshalJSON() ([]byte, error) {
    type Alias Patchpredictorrequest

    if PatchpredictorrequestMarshalled {
        return []byte("{}"), nil
    }
    PatchpredictorrequestMarshalled = true

    return json.Marshal(&struct { 
        RoutingTimeoutSeconds int `json:"routingTimeoutSeconds"`
        
        Schedule Predictorschedule `json:"schedule"`
        
        WorkloadBalancingConfig Predictorworkloadbalancing `json:"workloadBalancingConfig"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

