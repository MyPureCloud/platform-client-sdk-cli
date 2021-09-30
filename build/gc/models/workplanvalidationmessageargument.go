package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanvalidationmessageargumentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanvalidationmessageargumentDud struct { 
    


    

}

// Workplanvalidationmessageargument
type Workplanvalidationmessageargument struct { 
    // VarType - The type of the argument associated with violation messages
    VarType string `json:"type"`


    // Value - The value of the argument
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Workplanvalidationmessageargument) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanvalidationmessageargument) MarshalJSON() ([]byte, error) {
    type Alias Workplanvalidationmessageargument

    if WorkplanvalidationmessageargumentMarshalled {
        return []byte("{}"), nil
    }
    WorkplanvalidationmessageargumentMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Value string `json:"value"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

