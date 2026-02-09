package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementobservationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementobservationresultDud struct { 
    


    

}

// Taskmanagementobservationresult
type Taskmanagementobservationresult struct { 
    // Group - The group dimensions and their values for this result. Represents the combination of groupBy dimensions that define this result set.
    Group Taskmanagementobservationgroupresult `json:"group"`


    // Data - The metric data for this group. Contains the actual observation values for each requested metric.
    Data []Taskmanagementobservationdatacontainer `json:"data"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementobservationresult) String() string {
    
     o.Data = []Taskmanagementobservationdatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementobservationresult) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementobservationresult

    if TaskmanagementobservationresultMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementobservationresultMarshalled = true

    return json.Marshal(&struct {
        
        Group Taskmanagementobservationgroupresult `json:"group"`
        
        Data []Taskmanagementobservationdatacontainer `json:"data"`
        *Alias
    }{

        


        
        Data: []Taskmanagementobservationdatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

