package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CalltargetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CalltargetDud struct { 
    


    

}

// Calltarget
type Calltarget struct { 
    // VarType - The type of call
    VarType string `json:"type"`


    // Value - The id of the station or an E.164 formatted phone number
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Calltarget) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Calltarget) MarshalJSON() ([]byte, error) {
    type Alias Calltarget

    if CalltargetMarshalled {
        return []byte("{}"), nil
    }
    CalltargetMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Value string `json:"value"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

