package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulermessageargumentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulermessageargumentDud struct { 
    


    

}

// Schedulermessageargument
type Schedulermessageargument struct { 
    // VarType - The type of this message parameter
    VarType string `json:"type"`


    // Value - The value of this message parameter
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Schedulermessageargument) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulermessageargument) MarshalJSON() ([]byte, error) {
    type Alias Schedulermessageargument

    if SchedulermessageargumentMarshalled {
        return []byte("{}"), nil
    }
    SchedulermessageargumentMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Value string `json:"value"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

