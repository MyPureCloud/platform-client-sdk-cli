package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementobservationmetricstatsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementobservationmetricstatsDud struct { 
    


    

}

// Taskmanagementobservationmetricstats
type Taskmanagementobservationmetricstats struct { 
    // Count - The observed value for this metric
    Count int `json:"count"`


    // Max - The maximum observed value for this metric. Used for `oWorkitemOldestUnassigned` and  `oWorkitemOldestAssigned`
    Max int `json:"max"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementobservationmetricstats) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementobservationmetricstats) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementobservationmetricstats

    if TaskmanagementobservationmetricstatsMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementobservationmetricstatsMarshalled = true

    return json.Marshal(&struct {
        
        Count int `json:"count"`
        
        Max int `json:"max"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

