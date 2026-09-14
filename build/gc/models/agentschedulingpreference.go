package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentschedulingpreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentschedulingpreferenceDud struct { 
    Id string `json:"id"`


    


    

}

// Agentschedulingpreference
type Agentschedulingpreference struct { 
    


    // TimeSpan - Exact date, time and length of the scheduling preference time span
    TimeSpan Wfmtimespan `json:"timeSpan"`


    // PreferenceLevel - The preference level for this time span
    PreferenceLevel string `json:"preferenceLevel"`

}

// String returns a JSON representation of the model
func (o *Agentschedulingpreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentschedulingpreference) MarshalJSON() ([]byte, error) {
    type Alias Agentschedulingpreference

    if AgentschedulingpreferenceMarshalled {
        return []byte("{}"), nil
    }
    AgentschedulingpreferenceMarshalled = true

    return json.Marshal(&struct {
        
        TimeSpan Wfmtimespan `json:"timeSpan"`
        
        PreferenceLevel string `json:"preferenceLevel"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

