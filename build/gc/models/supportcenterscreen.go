package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcenterscreenMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcenterscreenDud struct { 
    


    

}

// Supportcenterscreen
type Supportcenterscreen struct { 
    // VarType - The type of the screen
    VarType string `json:"type"`


    // ModuleSettings - Module settings for the screen
    ModuleSettings []Supportcentermodulesetting `json:"moduleSettings"`

}

// String returns a JSON representation of the model
func (o *Supportcenterscreen) String() string {
    
    
    
    
    
    
     o.ModuleSettings = []Supportcentermodulesetting{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcenterscreen) MarshalJSON() ([]byte, error) {
    type Alias Supportcenterscreen

    if SupportcenterscreenMarshalled {
        return []byte("{}"), nil
    }
    SupportcenterscreenMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        ModuleSettings []Supportcentermodulesetting `json:"moduleSettings"`
        
        *Alias
    }{
        

        

        

        
        ModuleSettings: []Supportcentermodulesetting{{}},
        

        
        Alias: (*Alias)(u),
    })
}

