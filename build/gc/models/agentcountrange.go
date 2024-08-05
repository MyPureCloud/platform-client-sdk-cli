package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentcountrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentcountrangeDud struct { 
    


    

}

// Agentcountrange
type Agentcountrange struct { 
    // Minimum - The minimum value of agent count per work plan
    Minimum int `json:"minimum"`


    // Maximum - The maximum value of agent count per work plan
    Maximum int `json:"maximum"`

}

// String returns a JSON representation of the model
func (o *Agentcountrange) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentcountrange) MarshalJSON() ([]byte, error) {
    type Alias Agentcountrange

    if AgentcountrangeMarshalled {
        return []byte("{}"), nil
    }
    AgentcountrangeMarshalled = true

    return json.Marshal(&struct {
        
        Minimum int `json:"minimum"`
        
        Maximum int `json:"maximum"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

