package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatepredictorrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatepredictorrequestDud struct { 
    


    


    


    


    

}

// Createpredictorrequest
type Createpredictorrequest struct { 
    // QueueIds - The queue IDs associated with the predictor.
    QueueIds []string `json:"queueIds"`


    // Kpi - The KPI that the predictor attempts to maximize/minimize.
    Kpi string `json:"kpi"`


    // RoutingTimeoutSeconds - Number of seconds allocated to predictive routing before attempting a different routing method. This is a value between 12 and 900 seconds.
    RoutingTimeoutSeconds int `json:"routingTimeoutSeconds"`


    // Schedule - The predictor schedule that determines when the predictor is used for routing interactions.
    Schedule Predictorschedule `json:"schedule"`


    // WorkloadBalancingConfig - The predictor balancing configuration to enable workload balancing
    WorkloadBalancingConfig Predictorworkloadbalancing `json:"workloadBalancingConfig"`

}

// String returns a JSON representation of the model
func (o *Createpredictorrequest) String() string {
     o.QueueIds = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createpredictorrequest) MarshalJSON() ([]byte, error) {
    type Alias Createpredictorrequest

    if CreatepredictorrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatepredictorrequestMarshalled = true

    return json.Marshal(&struct {
        
        QueueIds []string `json:"queueIds"`
        
        Kpi string `json:"kpi"`
        
        RoutingTimeoutSeconds int `json:"routingTimeoutSeconds"`
        
        Schedule Predictorschedule `json:"schedule"`
        
        WorkloadBalancingConfig Predictorworkloadbalancing `json:"workloadBalancingConfig"`
        *Alias
    }{

        
        QueueIds: []string{""},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

