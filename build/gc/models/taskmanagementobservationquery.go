package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementobservationqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementobservationqueryDud struct { 
    


    


    


    

}

// Taskmanagementobservationquery
type Taskmanagementobservationquery struct { 
    // GroupBy - Dimension(s) to group by. Determines how the results will be grouped in the response.
    GroupBy []string `json:"groupBy"`


    // Metrics - List of metrics to be retrieved. Specifies which observational metrics should be included in the response.
    Metrics []Taskmanagementquerymetric `json:"metrics"`


    // Filter - Filter to return a subset of observations.
    Filter Taskmanagementobservationqueryfilter `json:"filter"`


    // Expands - List of properties to expand. Additional details about the objects returned in the results will be included in the response if supplied.
    Expands []string `json:"expands"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementobservationquery) String() string {
     o.GroupBy = []string{""} 
     o.Metrics = []Taskmanagementquerymetric{{}} 
    
     o.Expands = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementobservationquery) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementobservationquery

    if TaskmanagementobservationqueryMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementobservationqueryMarshalled = true

    return json.Marshal(&struct {
        
        GroupBy []string `json:"groupBy"`
        
        Metrics []Taskmanagementquerymetric `json:"metrics"`
        
        Filter Taskmanagementobservationqueryfilter `json:"filter"`
        
        Expands []string `json:"expands"`
        *Alias
    }{

        
        GroupBy: []string{""},
        


        
        Metrics: []Taskmanagementquerymetric{{}},
        


        


        
        Expands: []string{""},
        

        Alias: (*Alias)(u),
    })
}

