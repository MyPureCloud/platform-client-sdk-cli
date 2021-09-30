package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LexslotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LexslotDud struct { 
    


    


    


    

}

// Lexslot
type Lexslot struct { 
    // Name - The slot name
    Name string `json:"name"`


    // Description - The slot description
    Description string `json:"description"`


    // VarType - The slot type
    VarType string `json:"type"`


    // Priority - The priority of the slot
    Priority int `json:"priority"`

}

// String returns a JSON representation of the model
func (o *Lexslot) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lexslot) MarshalJSON() ([]byte, error) {
    type Alias Lexslot

    if LexslotMarshalled {
        return []byte("{}"), nil
    }
    LexslotMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        VarType string `json:"type"`
        
        Priority int `json:"priority"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

