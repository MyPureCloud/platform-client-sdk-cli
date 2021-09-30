package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallcommandMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallcommandDud struct { 
    


    

}

// Callcommand
type Callcommand struct { 
    // CallNumber - The phone number to dial for this call.
    CallNumber string `json:"callNumber"`


    // PhoneColumn - For a dialer preview or scheduled callback, the phone column associated with the phone number
    PhoneColumn string `json:"phoneColumn"`

}

// String returns a JSON representation of the model
func (o *Callcommand) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callcommand) MarshalJSON() ([]byte, error) {
    type Alias Callcommand

    if CallcommandMarshalled {
        return []byte("{}"), nil
    }
    CallcommandMarshalled = true

    return json.Marshal(&struct { 
        CallNumber string `json:"callNumber"`
        
        PhoneColumn string `json:"phoneColumn"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

