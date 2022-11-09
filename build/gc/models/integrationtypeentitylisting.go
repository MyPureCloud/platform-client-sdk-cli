package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntegrationtypeentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntegrationtypeentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Integrationtypeentitylisting
type Integrationtypeentitylisting struct { 
    // Entities
    Entities []Integrationtype `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Integrationtypeentitylisting) String() string {
     o.Entities = []Integrationtype{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Integrationtypeentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Integrationtypeentitylisting

    if IntegrationtypeentitylistingMarshalled {
        return []byte("{}"), nil
    }
    IntegrationtypeentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Integrationtype `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Integrationtype{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

