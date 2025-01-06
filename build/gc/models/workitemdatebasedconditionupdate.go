package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemdatebasedconditionupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemdatebasedconditionupdateDud struct { 
    


    

}

// Workitemdatebasedconditionupdate
type Workitemdatebasedconditionupdate struct { 
    // Attribute - The name of the workitem date attribute.
    Attribute string `json:"attribute"`


    // RelativeMinutesToInvocation - The time in minutes before or after the date attribute.
    RelativeMinutesToInvocation int `json:"relativeMinutesToInvocation"`

}

// String returns a JSON representation of the model
func (o *Workitemdatebasedconditionupdate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemdatebasedconditionupdate) MarshalJSON() ([]byte, error) {
    type Alias Workitemdatebasedconditionupdate

    if WorkitemdatebasedconditionupdateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemdatebasedconditionupdateMarshalled = true

    return json.Marshal(&struct {
        
        Attribute string `json:"attribute"`
        
        RelativeMinutesToInvocation int `json:"relativeMinutesToInvocation"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

