package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LibrarydivisionviewentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LibrarydivisionviewentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Librarydivisionviewentitylisting
type Librarydivisionviewentitylisting struct { 
    // Entities
    Entities []Librarydivisionview `json:"entities"`


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
func (o *Librarydivisionviewentitylisting) String() string {
     o.Entities = []Librarydivisionview{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Librarydivisionviewentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Librarydivisionviewentitylisting

    if LibrarydivisionviewentitylistingMarshalled {
        return []byte("{}"), nil
    }
    LibrarydivisionviewentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Librarydivisionview `json:"entities"`
        
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

        
        Entities: []Librarydivisionview{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

