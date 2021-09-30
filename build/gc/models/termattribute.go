package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TermattributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TermattributeDud struct { 
    


    


    

}

// Termattribute
type Termattribute struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // VarType
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Termattribute) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Termattribute) MarshalJSON() ([]byte, error) {
    type Alias Termattribute

    if TermattributeMarshalled {
        return []byte("{}"), nil
    }
    TermattributeMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

