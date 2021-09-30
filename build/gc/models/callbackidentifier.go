package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallbackidentifierMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallbackidentifierDud struct { 
    


    

}

// Callbackidentifier
type Callbackidentifier struct { 
    // VarType - The type of the associated callback participant
    VarType string `json:"type"`


    // Id - The identifier of the callback
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Callbackidentifier) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callbackidentifier) MarshalJSON() ([]byte, error) {
    type Alias Callbackidentifier

    if CallbackidentifierMarshalled {
        return []byte("{}"), nil
    }
    CallbackidentifierMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Id string `json:"id"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

