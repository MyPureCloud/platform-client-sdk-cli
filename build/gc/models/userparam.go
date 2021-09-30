package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserparamMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserparamDud struct { 
    


    

}

// Userparam
type Userparam struct { 
    // Key
    Key string `json:"key"`


    // Value
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Userparam) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userparam) MarshalJSON() ([]byte, error) {
    type Alias Userparam

    if UserparamMarshalled {
        return []byte("{}"), nil
    }
    UserparamMarshalled = true

    return json.Marshal(&struct { 
        Key string `json:"key"`
        
        Value string `json:"value"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

