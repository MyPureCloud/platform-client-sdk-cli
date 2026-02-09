package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementobservationpredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementobservationpredicateDud struct { 
    


    

}

// Taskmanagementobservationpredicate
type Taskmanagementobservationpredicate struct { 
    // Dimension - The dimension to filter on
    Dimension string `json:"dimension"`


    // Value - The value to filter by for the specified dimension
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementobservationpredicate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementobservationpredicate) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementobservationpredicate

    if TaskmanagementobservationpredicateMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementobservationpredicateMarshalled = true

    return json.Marshal(&struct {
        
        Dimension string `json:"dimension"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

