package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementobservationdatacontainerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementobservationdatacontainerDud struct { 
    


    


    

}

// Taskmanagementobservationdatacontainer
type Taskmanagementobservationdatacontainer struct { 
    // Metric - The metric for this data point
    Metric string `json:"metric"`


    // Stats - The observed statistics for this metric
    Stats Taskmanagementobservationmetricstats `json:"stats"`


    // Qualifier - Qualifier for duration based metrics.
    Qualifier string `json:"qualifier"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementobservationdatacontainer) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementobservationdatacontainer) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementobservationdatacontainer

    if TaskmanagementobservationdatacontainerMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementobservationdatacontainerMarshalled = true

    return json.Marshal(&struct {
        
        Metric string `json:"metric"`
        
        Stats Taskmanagementobservationmetricstats `json:"stats"`
        
        Qualifier string `json:"qualifier"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

