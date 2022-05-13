package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchemareferenceentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchemareferenceentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Schemareferenceentitylisting
type Schemareferenceentitylisting struct { 
    // Entities
    Entities []Domainschemareference `json:"entities"`


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


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Schemareferenceentitylisting) String() string {
     o.Entities = []Domainschemareference{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schemareferenceentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Schemareferenceentitylisting

    if SchemareferenceentitylistingMarshalled {
        return []byte("{}"), nil
    }
    SchemareferenceentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Domainschemareference `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Domainschemareference{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

