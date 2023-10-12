package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NuancechannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NuancechannelDud struct { 
    


    


    


    

}

// Nuancechannel - Model for a Nuance channel
type Nuancechannel struct { 
    // Id - The channel ID
    Id string `json:"id"`


    // Name - The channel name
    Name string `json:"name"`


    // Modes - Supported Channel Modes
    Modes []string `json:"modes"`


    // Color - The Channel Color
    Color string `json:"color"`

}

// String returns a JSON representation of the model
func (o *Nuancechannel) String() string {
    
    
     o.Modes = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nuancechannel) MarshalJSON() ([]byte, error) {
    type Alias Nuancechannel

    if NuancechannelMarshalled {
        return []byte("{}"), nil
    }
    NuancechannelMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Modes []string `json:"modes"`
        
        Color string `json:"color"`
        *Alias
    }{

        


        


        
        Modes: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

