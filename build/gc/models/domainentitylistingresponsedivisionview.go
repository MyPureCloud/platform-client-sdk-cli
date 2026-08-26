package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainentitylistingresponsedivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainentitylistingresponsedivisionviewDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Domainentitylistingresponsedivisionview
type Domainentitylistingresponsedivisionview struct { 
    // Entities
    Entities []Responsedivisionview `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Domainentitylistingresponsedivisionview) String() string {
     o.Entities = []Responsedivisionview{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainentitylistingresponsedivisionview) MarshalJSON() ([]byte, error) {
    type Alias Domainentitylistingresponsedivisionview

    if DomainentitylistingresponsedivisionviewMarshalled {
        return []byte("{}"), nil
    }
    DomainentitylistingresponsedivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Responsedivisionview `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Responsedivisionview{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

