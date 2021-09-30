package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DisconnectreasonMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DisconnectreasonDud struct { 
    


    


    

}

// Disconnectreason
type Disconnectreason struct { 
    // VarType - Disconnect reason protocol type.
    VarType string `json:"type"`


    // Code - Protocol specific reason code. See the Q.850 and SIP specs.
    Code int `json:"code"`


    // Phrase - Human readable English description of the disconnect reason.
    Phrase string `json:"phrase"`

}

// String returns a JSON representation of the model
func (o *Disconnectreason) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Disconnectreason) MarshalJSON() ([]byte, error) {
    type Alias Disconnectreason

    if DisconnectreasonMarshalled {
        return []byte("{}"), nil
    }
    DisconnectreasonMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Code int `json:"code"`
        
        Phrase string `json:"phrase"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

