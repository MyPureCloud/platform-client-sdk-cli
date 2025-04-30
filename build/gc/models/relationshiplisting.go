package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RelationshiplistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RelationshiplistingDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Relationshiplisting
type Relationshiplisting struct { 
    // Entities
    Entities []Relationship `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // PartialResults
    PartialResults bool `json:"partialResults"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Relationshiplisting) String() string {
     o.Entities = []Relationship{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Relationshiplisting) MarshalJSON() ([]byte, error) {
    type Alias Relationshiplisting

    if RelationshiplistingMarshalled {
        return []byte("{}"), nil
    }
    RelationshiplistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Relationship `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PartialResults bool `json:"partialResults"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Relationship{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

