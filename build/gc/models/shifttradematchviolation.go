package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradematchviolationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradematchviolationDud struct { 
    


    

}

// Shifttradematchviolation
type Shifttradematchviolation struct { 
    // VarType - The type of constraint violation
    VarType string `json:"type"`


    // Params - Clarifying user params for constructing helpful error messages
    Params map[string]string `json:"params"`

}

// String returns a JSON representation of the model
func (o *Shifttradematchviolation) String() string {
    
    
    
    
    
    
     o.Params = map[string]string{"": ""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradematchviolation) MarshalJSON() ([]byte, error) {
    type Alias Shifttradematchviolation

    if ShifttradematchviolationMarshalled {
        return []byte("{}"), nil
    }
    ShifttradematchviolationMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Params map[string]string `json:"params"`
        
        *Alias
    }{
        

        

        

        
        Params: map[string]string{"": ""},
        

        
        Alias: (*Alias)(u),
    })
}

