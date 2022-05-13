package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AddworkplanrotationagentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AddworkplanrotationagentrequestDud struct { 
    


    


    

}

// Addworkplanrotationagentrequest
type Addworkplanrotationagentrequest struct { 
    // UserId - The ID of an agent in this work plan rotation
    UserId string `json:"userId"`


    // DateRange - The date range to which this agent is effective in the work plan rotation
    DateRange Daterangewithoptionalend `json:"dateRange"`


    // Position - Start position of the work plan in the pattern for this agent in the work plan rotation. Position value starts from 0
    Position int `json:"position"`

}

// String returns a JSON representation of the model
func (o *Addworkplanrotationagentrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Addworkplanrotationagentrequest) MarshalJSON() ([]byte, error) {
    type Alias Addworkplanrotationagentrequest

    if AddworkplanrotationagentrequestMarshalled {
        return []byte("{}"), nil
    }
    AddworkplanrotationagentrequestMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        DateRange Daterangewithoptionalend `json:"dateRange"`
        
        Position int `json:"position"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

