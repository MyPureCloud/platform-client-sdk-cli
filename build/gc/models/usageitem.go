package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsageitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsageitemDud struct { 
    


    


    

}

// Usageitem
type Usageitem struct { 
    // VarType
    VarType string `json:"type"`


    // TotalDocumentByteCount
    TotalDocumentByteCount int `json:"totalDocumentByteCount"`


    // TotalDocumentCount
    TotalDocumentCount int `json:"totalDocumentCount"`

}

// String returns a JSON representation of the model
func (o *Usageitem) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usageitem) MarshalJSON() ([]byte, error) {
    type Alias Usageitem

    if UsageitemMarshalled {
        return []byte("{}"), nil
    }
    UsageitemMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        TotalDocumentByteCount int `json:"totalDocumentByteCount"`
        
        TotalDocumentCount int `json:"totalDocumentCount"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

