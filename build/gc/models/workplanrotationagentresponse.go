package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanrotationagentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanrotationagentresponseDud struct { 
    


    


    

}

// Workplanrotationagentresponse
type Workplanrotationagentresponse struct { 
    // User - The user associated with this work plan rotation
    User Userreference `json:"user"`


    // DateRange - The date range to which this agent is effective in the work plan rotation
    DateRange Daterangewithoptionalend `json:"dateRange"`


    // Position - Start position of the work plan in the pattern for this agent in the work plan rotation. Position value starts from 0
    Position int `json:"position"`

}

// String returns a JSON representation of the model
func (o *Workplanrotationagentresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanrotationagentresponse) MarshalJSON() ([]byte, error) {
    type Alias Workplanrotationagentresponse

    if WorkplanrotationagentresponseMarshalled {
        return []byte("{}"), nil
    }
    WorkplanrotationagentresponseMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        DateRange Daterangewithoptionalend `json:"dateRange"`
        
        Position int `json:"position"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

