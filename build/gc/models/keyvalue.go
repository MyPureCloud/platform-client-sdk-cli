package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KeyvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KeyvalueDud struct { 
    


    

}

// Keyvalue
type Keyvalue struct { 
    // Key - Key for free-form data.
    Key string `json:"key"`


    // Value - Value for free-form data.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Keyvalue) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Keyvalue) MarshalJSON() ([]byte, error) {
    type Alias Keyvalue

    if KeyvalueMarshalled {
        return []byte("{}"), nil
    }
    KeyvalueMarshalled = true

    return json.Marshal(&struct { 
        Key string `json:"key"`
        
        Value string `json:"value"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

