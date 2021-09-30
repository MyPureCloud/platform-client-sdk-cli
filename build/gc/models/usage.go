package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsageDud struct { 
    

}

// Usage
type Usage struct { 
    // Types
    Types []Usageitem `json:"types"`

}

// String returns a JSON representation of the model
func (o *Usage) String() string {
    
    
     o.Types = []Usageitem{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usage) MarshalJSON() ([]byte, error) {
    type Alias Usage

    if UsageMarshalled {
        return []byte("{}"), nil
    }
    UsageMarshalled = true

    return json.Marshal(&struct { 
        Types []Usageitem `json:"types"`
        
        *Alias
    }{
        

        
        Types: []Usageitem{{}},
        

        
        Alias: (*Alias)(u),
    })
}

