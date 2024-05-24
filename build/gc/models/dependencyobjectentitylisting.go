package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DependencyobjectentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DependencyobjectentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Dependencyobjectentitylisting
type Dependencyobjectentitylisting struct { 
    // Entities
    Entities []Dependencyobject `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Dependencyobjectentitylisting) String() string {
     o.Entities = []Dependencyobject{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dependencyobjectentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Dependencyobjectentitylisting

    if DependencyobjectentitylistingMarshalled {
        return []byte("{}"), nil
    }
    DependencyobjectentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Dependencyobject `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Dependencyobject{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

