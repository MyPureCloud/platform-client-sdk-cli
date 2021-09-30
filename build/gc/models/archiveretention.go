package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArchiveretentionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArchiveretentionDud struct { 
    


    

}

// Archiveretention
type Archiveretention struct { 
    // Days
    Days int `json:"days"`


    // StorageMedium
    StorageMedium string `json:"storageMedium"`

}

// String returns a JSON representation of the model
func (o *Archiveretention) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Archiveretention) MarshalJSON() ([]byte, error) {
    type Alias Archiveretention

    if ArchiveretentionMarshalled {
        return []byte("{}"), nil
    }
    ArchiveretentionMarshalled = true

    return json.Marshal(&struct { 
        Days int `json:"days"`
        
        StorageMedium string `json:"storageMedium"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

