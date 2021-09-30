package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExpansioncriteriumMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExpansioncriteriumDud struct { 
    


    

}

// Expansioncriterium
type Expansioncriterium struct { 
    // VarType
    VarType string `json:"type"`


    // Threshold
    Threshold float64 `json:"threshold"`

}

// String returns a JSON representation of the model
func (o *Expansioncriterium) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Expansioncriterium) MarshalJSON() ([]byte, error) {
    type Alias Expansioncriterium

    if ExpansioncriteriumMarshalled {
        return []byte("{}"), nil
    }
    ExpansioncriteriumMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Threshold float64 `json:"threshold"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

