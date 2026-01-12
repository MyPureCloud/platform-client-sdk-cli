package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnavailabletimeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnavailabletimeDud struct { 
    Id string `json:"id"`


    


    

}

// Unavailabletime
type Unavailabletime struct { 
    


    // TimeSpan - Exact date, time and length of the unavailability time span
    TimeSpan Unavailabletimestimespan `json:"timeSpan"`


    // Notes - Comments explaining the unavailability time span
    Notes string `json:"notes"`

}

// String returns a JSON representation of the model
func (o *Unavailabletime) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unavailabletime) MarshalJSON() ([]byte, error) {
    type Alias Unavailabletime

    if UnavailabletimeMarshalled {
        return []byte("{}"), nil
    }
    UnavailabletimeMarshalled = true

    return json.Marshal(&struct {
        
        TimeSpan Unavailabletimestimespan `json:"timeSpan"`
        
        Notes string `json:"notes"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

