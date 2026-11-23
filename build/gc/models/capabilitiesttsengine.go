package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesttsengineMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesttsengineDud struct { 
    


    


    


    

}

// Capabilitiesttsengine
type Capabilitiesttsengine struct { 
    // IsSecure
    IsSecure bool `json:"isSecure"`


    // Name
    Name string `json:"name"`


    // MaxRequestLength
    MaxRequestLength int `json:"maxRequestLength"`


    // Voices
    Voices map[string]Capabilitiesttsvoice `json:"voices"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesttsengine) String() string {
    
    
    
     o.Voices = map[string]Capabilitiesttsvoice{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesttsengine) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesttsengine

    if CapabilitiesttsengineMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesttsengineMarshalled = true

    return json.Marshal(&struct {
        
        IsSecure bool `json:"isSecure"`
        
        Name string `json:"name"`
        
        MaxRequestLength int `json:"maxRequestLength"`
        
        Voices map[string]Capabilitiesttsvoice `json:"voices"`
        *Alias
    }{

        


        


        


        
        Voices: map[string]Capabilitiesttsvoice{"": {}},
        

        Alias: (*Alias)(u),
    })
}

