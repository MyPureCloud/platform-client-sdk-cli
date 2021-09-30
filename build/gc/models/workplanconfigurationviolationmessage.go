package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanconfigurationviolationmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanconfigurationviolationmessageDud struct { 
    


    


    

}

// Workplanconfigurationviolationmessage
type Workplanconfigurationviolationmessage struct { 
    // VarType - Type of configuration violation message for this work plan
    VarType string `json:"type"`


    // Arguments - Arguments of the message that provide information about the misconfigured value or the threshold that is exceeded by the misconfigured value
    Arguments []Workplanvalidationmessageargument `json:"arguments"`


    // Severity - Severity of the message. A message with Error severity indicates the scheduler won't be able to produce schedules and thus the work plan is invalid.
    Severity string `json:"severity"`

}

// String returns a JSON representation of the model
func (o *Workplanconfigurationviolationmessage) String() string {
    
    
    
    
    
    
     o.Arguments = []Workplanvalidationmessageargument{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanconfigurationviolationmessage) MarshalJSON() ([]byte, error) {
    type Alias Workplanconfigurationviolationmessage

    if WorkplanconfigurationviolationmessageMarshalled {
        return []byte("{}"), nil
    }
    WorkplanconfigurationviolationmessageMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Arguments []Workplanvalidationmessageargument `json:"arguments"`
        
        Severity string `json:"severity"`
        
        *Alias
    }{
        

        

        

        
        Arguments: []Workplanvalidationmessageargument{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

