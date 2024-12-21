package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemdatebasedconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemdatebasedconditionDud struct { 
    


    

}

// Workitemdatebasedcondition
type Workitemdatebasedcondition struct { 
    // Attribute - The name of the workitem date attribute.
    Attribute string `json:"attribute"`


    // RelativeMinutesToInvocation - The time in minutes before or after the date attribute.
    RelativeMinutesToInvocation int `json:"relativeMinutesToInvocation"`

}

// String returns a JSON representation of the model
func (o *Workitemdatebasedcondition) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemdatebasedcondition) MarshalJSON() ([]byte, error) {
    type Alias Workitemdatebasedcondition

    if WorkitemdatebasedconditionMarshalled {
        return []byte("{}"), nil
    }
    WorkitemdatebasedconditionMarshalled = true

    return json.Marshal(&struct {
        
        Attribute string `json:"attribute"`
        
        RelativeMinutesToInvocation int `json:"relativeMinutesToInvocation"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

