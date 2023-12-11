package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UtilizationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UtilizationrequestDud struct { 
    

}

// Utilizationrequest
type Utilizationrequest struct { 
    // Utilization - Map of media type to utilization settings.
    Utilization map[string]Mediautilization `json:"utilization"`

}

// String returns a JSON representation of the model
func (o *Utilizationrequest) String() string {
     o.Utilization = map[string]Mediautilization{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Utilizationrequest) MarshalJSON() ([]byte, error) {
    type Alias Utilizationrequest

    if UtilizationrequestMarshalled {
        return []byte("{}"), nil
    }
    UtilizationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Utilization map[string]Mediautilization `json:"utilization"`
        *Alias
    }{

        
        Utilization: map[string]Mediautilization{"": {}},
        

        Alias: (*Alias)(u),
    })
}

