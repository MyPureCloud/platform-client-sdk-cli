package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateunavailabletimeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateunavailabletimeDud struct { 
    


    


    


    

}

// Updateunavailabletime
type Updateunavailabletime struct { 
    // Id - The ID of the unavailable time span. Should be specified to update or delete an existing unavailable time span or set to null when creating a new one
    Id string `json:"id"`


    // TimeSpan - Exact date, time and length of the unavailability time in granularity of minutes. Must be specified when creating a new unavailable time span
    TimeSpan Unavailabletimestimespan `json:"timeSpan"`


    // Notes - Comments explaining the unavailability time span
    Notes string `json:"notes"`


    // Delete - Whether the unavailable time should be deleted
    Delete bool `json:"delete"`

}

// String returns a JSON representation of the model
func (o *Updateunavailabletime) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateunavailabletime) MarshalJSON() ([]byte, error) {
    type Alias Updateunavailabletime

    if UpdateunavailabletimeMarshalled {
        return []byte("{}"), nil
    }
    UpdateunavailabletimeMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        TimeSpan Unavailabletimestimespan `json:"timeSpan"`
        
        Notes string `json:"notes"`
        
        Delete bool `json:"delete"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

