package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementquerymetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementquerymetricDud struct { 
    


    

}

// Taskmanagementquerymetric
type Taskmanagementquerymetric struct { 
    // Name - The requested metric name
    Name string `json:"name"`


    // Qualifier - Qualifier for duration based metrics. Required when requesting oWorkitemsDue
    Qualifier string `json:"qualifier"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementquerymetric) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementquerymetric) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementquerymetric

    if TaskmanagementquerymetricMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementquerymetricMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Qualifier string `json:"qualifier"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

