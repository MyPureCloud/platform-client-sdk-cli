package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymentconfigurationversionentityrefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymentconfigurationversionentityrefDud struct { 
    


    


    


    

}

// Webdeploymentconfigurationversionentityref
type Webdeploymentconfigurationversionentityref struct { 
    // Id - The configuration version ID
    Id string `json:"id"`


    // Name - The configuration version name
    Name string `json:"name"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // Version - The version of the configuration
    Version string `json:"version"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentconfigurationversionentityref) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeploymentconfigurationversionentityref) MarshalJSON() ([]byte, error) {
    type Alias Webdeploymentconfigurationversionentityref

    if WebdeploymentconfigurationversionentityrefMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymentconfigurationversionentityrefMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        
        Version string `json:"version"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

