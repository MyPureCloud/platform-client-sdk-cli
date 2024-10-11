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
    // Entities
    Entities []Webdeploymentconfigurationversion `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // Total
    Total int `json:"total"`

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
        
        Entities []Webdeploymentconfigurationversion `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        Total int `json:"total"`
        *Alias
    }{

        
        Entities: []Webdeploymentconfigurationversion{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

