package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateworkplanrotationagentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateworkplanrotationagentrequestDud struct { 
    


    


    


    

}

// Updateworkplanrotationagentrequest
type Updateworkplanrotationagentrequest struct { 
    // UserId - The ID of an agent in this work plan rotation
    UserId string `json:"userId"`


    // DateRange - The date range to which this agent is effective in the work plan rotation
    DateRange Daterangewithoptionalend `json:"dateRange"`


    // Position - Start position of the work plan in the pattern for this agent in the work plan rotation. Position value starts from 0
    Position int `json:"position"`


    // Delete - If marked true for this agent when updating, then this agent will be removed from this work plan rotation
    Delete bool `json:"delete"`

}

// String returns a JSON representation of the model
func (o *Updateworkplanrotationagentrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateworkplanrotationagentrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateworkplanrotationagentrequest

    if UpdateworkplanrotationagentrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateworkplanrotationagentrequestMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        DateRange Daterangewithoptionalend `json:"dateRange"`
        
        Position int `json:"position"`
        
        Delete bool `json:"delete"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

