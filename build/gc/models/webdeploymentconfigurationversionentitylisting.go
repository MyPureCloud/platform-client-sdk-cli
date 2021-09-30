package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymentconfigurationversionentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymentconfigurationversionentitylistingDud struct { 
    


    


    

}

// Webdeploymentconfigurationversionentitylisting
type Webdeploymentconfigurationversionentitylisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Webdeploymentconfigurationversion `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentconfigurationversionentitylisting) String() string {
    
    
    
    
    
    
     o.Entities = []Webdeploymentconfigurationversion{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeploymentconfigurationversionentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Webdeploymentconfigurationversionentitylisting

    if WebdeploymentconfigurationversionentitylistingMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymentconfigurationversionentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Total int `json:"total"`
        
        Entities []Webdeploymentconfigurationversion `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        
        Entities: []Webdeploymentconfigurationversion{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

