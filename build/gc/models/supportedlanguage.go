package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportedlanguageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportedlanguageDud struct { 
    Language string `json:"language"`


    IsDefault bool `json:"isDefault"`

}

// Supportedlanguage
type Supportedlanguage struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Supportedlanguage) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportedlanguage) MarshalJSON() ([]byte, error) {
    type Alias Supportedlanguage

    if SupportedlanguageMarshalled {
        return []byte("{}"), nil
    }
    SupportedlanguageMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

