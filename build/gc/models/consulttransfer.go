package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConsulttransferMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConsulttransferDud struct { 
    


    

}

// Consulttransfer
type Consulttransfer struct { 
    // SpeakTo - Determines to whom the initiating participant is speaking. Defaults to DESTINATION
    SpeakTo string `json:"speakTo"`


    // Destination - Destination phone number and name.
    Destination Destination `json:"destination"`

}

// String returns a JSON representation of the model
func (o *Consulttransfer) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Consulttransfer) MarshalJSON() ([]byte, error) {
    type Alias Consulttransfer

    if ConsulttransferMarshalled {
        return []byte("{}"), nil
    }
    ConsulttransferMarshalled = true

    return json.Marshal(&struct { 
        SpeakTo string `json:"speakTo"`
        
        Destination Destination `json:"destination"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

