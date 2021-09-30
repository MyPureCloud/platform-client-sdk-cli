package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgemetricsdiskMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgemetricsdiskDud struct { 
    


    


    

}

// Edgemetricsdisk
type Edgemetricsdisk struct { 
    // AvailableBytes - Available memory in bytes.
    AvailableBytes float64 `json:"availableBytes"`


    // PartitionName - Disk partition name.
    PartitionName string `json:"partitionName"`


    // TotalBytes - Total memory in bytes.
    TotalBytes float64 `json:"totalBytes"`

}

// String returns a JSON representation of the model
func (o *Edgemetricsdisk) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgemetricsdisk) MarshalJSON() ([]byte, error) {
    type Alias Edgemetricsdisk

    if EdgemetricsdiskMarshalled {
        return []byte("{}"), nil
    }
    EdgemetricsdiskMarshalled = true

    return json.Marshal(&struct { 
        AvailableBytes float64 `json:"availableBytes"`
        
        PartitionName string `json:"partitionName"`
        
        TotalBytes float64 `json:"totalBytes"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

