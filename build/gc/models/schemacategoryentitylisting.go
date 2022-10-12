package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchemacategoryentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchemacategoryentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Schemacategoryentitylisting
type Schemacategoryentitylisting struct { 
    // Entities
    Entities []Schemacategory `json:"entities"`


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
func (o *Schemacategoryentitylisting) String() string {
     o.Entities = []Schemacategory{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schemacategoryentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Schemacategoryentitylisting

    if SchemacategoryentitylistingMarshalled {
        return []byte("{}"), nil
    }
    SchemacategoryentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Schemacategory `json:"entities"`
        
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

        
        Entities: []Schemacategory{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

