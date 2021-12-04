package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SourceplanninggrouprequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SourceplanninggrouprequestDud struct { 
    


    

}

// Sourceplanninggrouprequest
type Sourceplanninggrouprequest struct { 
    // Id - The ID of the planning group
    Id string `json:"id"`


    // Metadata - Version metadata for the planning group
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Sourceplanninggrouprequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sourceplanninggrouprequest) MarshalJSON() ([]byte, error) {
    type Alias Sourceplanninggrouprequest

    if SourceplanninggrouprequestMarshalled {
        return []byte("{}"), nil
    }
    SourceplanninggrouprequestMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

