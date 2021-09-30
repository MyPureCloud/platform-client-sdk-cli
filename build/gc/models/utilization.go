package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UtilizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UtilizationDud struct { 
    

}

// Utilization
type Utilization struct { 
    // Utilization - Map of media type to utilization settings.  Valid media types include call, callback, chat, email, and message.
    Utilization map[string]Mediautilization `json:"utilization"`

}

// String returns a JSON representation of the model
func (o *Utilization) String() string {
    
    
     o.Utilization = map[string]Mediautilization{"": {}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Utilization) MarshalJSON() ([]byte, error) {
    type Alias Utilization

    if UtilizationMarshalled {
        return []byte("{}"), nil
    }
    UtilizationMarshalled = true

    return json.Marshal(&struct { 
        Utilization map[string]Mediautilization `json:"utilization"`
        
        *Alias
    }{
        

        
        Utilization: map[string]Mediautilization{"": {}},
        

        
        Alias: (*Alias)(u),
    })
}

