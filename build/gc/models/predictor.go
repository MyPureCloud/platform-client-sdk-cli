package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictorDud struct { 
    Id string `json:"id"`


    


    


    


    


    State string `json:"state"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    ErrorCode string `json:"errorCode"`


    Models []Predictormodelbrief `json:"models"`


    SelfUri string `json:"selfUri"`

}

// Predictor
type Predictor struct { 
    


    // Queues - The queue IDs associated with the predictor.
    Queues []Addressableentityref `json:"queues"`


    // Kpi - The KPI that the predictor attempts to maximize/minimize.
    Kpi string `json:"kpi"`


    // RoutingTimeoutSeconds - Number of seconds allocated to predictive routing before attempting a different routing method. This is a value between 12 and 900 seconds.
    RoutingTimeoutSeconds int `json:"routingTimeoutSeconds"`


    // Schedule - The predictor schedule that determines when the predictor is used for routing interactions.
    Schedule Predictorschedule `json:"schedule"`


    


    


    


    // WorkloadBalancingConfig - The predictor balancing configuration to enable workload balancing.
    WorkloadBalancingConfig Predictorworkloadbalancing `json:"workloadBalancingConfig"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Predictor) String() string {
     o.Queues = []Addressableentityref{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictor) MarshalJSON() ([]byte, error) {
    type Alias Predictor

    if PredictorMarshalled {
        return []byte("{}"), nil
    }
    PredictorMarshalled = true

    return json.Marshal(&struct {
        
        Queues []Addressableentityref `json:"queues"`
        
        Kpi string `json:"kpi"`
        
        RoutingTimeoutSeconds int `json:"routingTimeoutSeconds"`
        
        Schedule Predictorschedule `json:"schedule"`
        
        WorkloadBalancingConfig Predictorworkloadbalancing `json:"workloadBalancingConfig"`
        *Alias
    }{

        


        
        Queues: []Addressableentityref{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

