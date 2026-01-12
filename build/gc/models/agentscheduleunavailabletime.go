package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentscheduleunavailabletimeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentscheduleunavailabletimeDud struct { 
    


    

}

// Agentscheduleunavailabletime
type Agentscheduleunavailabletime struct { 
    // TimeSpan - Exact date, time and length of the unavailability time span
    TimeSpan Unavailabletimestimespan `json:"timeSpan"`


    // Notes - Comments explaining the unavailability time span
    Notes string `json:"notes"`

}

// String returns a JSON representation of the model
func (o *Agentscheduleunavailabletime) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentscheduleunavailabletime) MarshalJSON() ([]byte, error) {
    type Alias Agentscheduleunavailabletime

    if AgentscheduleunavailabletimeMarshalled {
        return []byte("{}"), nil
    }
    AgentscheduleunavailabletimeMarshalled = true

    return json.Marshal(&struct {
        
        TimeSpan Unavailabletimestimespan `json:"timeSpan"`
        
        Notes string `json:"notes"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

