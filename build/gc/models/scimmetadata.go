package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimmetadataDud struct { 
    ResourceType string `json:"resourceType"`


    LastModified time.Time `json:"lastModified"`


    Location string `json:"location"`


    Version string `json:"version"`

}

// Scimmetadata - Defines the SCIM metadata.
type Scimmetadata struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Scimmetadata) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimmetadata) MarshalJSON() ([]byte, error) {
    type Alias Scimmetadata

    if ScimmetadataMarshalled {
        return []byte("{}"), nil
    }
    ScimmetadataMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

