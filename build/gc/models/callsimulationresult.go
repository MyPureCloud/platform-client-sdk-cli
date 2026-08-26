package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallsimulationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallsimulationresultDud struct { 
    


    


    

}

// Callsimulationresult
type Callsimulationresult struct { 
    // Allowed - Whether the call is allowed
    Allowed bool `json:"allowed"`


    // Level - The simulation level
    Level string `json:"level"`


    // MatchedPrefix - The matched prefix
    MatchedPrefix string `json:"matchedPrefix"`

}

// String returns a JSON representation of the model
func (o *Callsimulationresult) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callsimulationresult) MarshalJSON() ([]byte, error) {
    type Alias Callsimulationresult

    if CallsimulationresultMarshalled {
        return []byte("{}"), nil
    }
    CallsimulationresultMarshalled = true

    return json.Marshal(&struct {
        
        Allowed bool `json:"allowed"`
        
        Level string `json:"level"`
        
        MatchedPrefix string `json:"matchedPrefix"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

