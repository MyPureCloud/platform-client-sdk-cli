package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementobservationqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementobservationqueryfilterDud struct { 
    

}

// Taskmanagementobservationqueryfilter
type Taskmanagementobservationqueryfilter struct { 
    // Predicates - List of predicates that define the filter conditions. Each predicate specifies a dimension and value to filter by. A single queueId predicate is always required.
    Predicates []Taskmanagementobservationpredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementobservationqueryfilter) String() string {
     o.Predicates = []Taskmanagementobservationpredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementobservationqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementobservationqueryfilter

    if TaskmanagementobservationqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementobservationqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        Predicates []Taskmanagementobservationpredicate `json:"predicates"`
        *Alias
    }{

        
        Predicates: []Taskmanagementobservationpredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

