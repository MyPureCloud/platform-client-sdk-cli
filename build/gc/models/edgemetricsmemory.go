package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgemetricsmemoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgemetricsmemoryDud struct { 
    


    


    

}

// Edgemetricsmemory
type Edgemetricsmemory struct { 
    // AvailableBytes - Available memory in bytes.
    AvailableBytes float64 `json:"availableBytes"`


    // VarType - Type of memory. Virtual or physical.
    VarType string `json:"type"`


    // TotalBytes - Total memory in bytes.
    TotalBytes float64 `json:"totalBytes"`

}

// String returns a JSON representation of the model
func (o *Edgemetricsmemory) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgemetricsmemory) MarshalJSON() ([]byte, error) {
    type Alias Edgemetricsmemory

    if EdgemetricsmemoryMarshalled {
        return []byte("{}"), nil
    }
    EdgemetricsmemoryMarshalled = true

    return json.Marshal(&struct { 
        AvailableBytes float64 `json:"availableBytes"`
        
        VarType string `json:"type"`
        
        TotalBytes float64 `json:"totalBytes"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

