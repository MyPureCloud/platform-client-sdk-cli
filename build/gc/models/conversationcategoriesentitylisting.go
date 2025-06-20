package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcategoriesentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcategoriesentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Conversationcategoriesentitylisting
type Conversationcategoriesentitylisting struct { 
    // Entities
    Entities []Conversationcategory `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // LastUri
    LastUri string `json:"lastUri"`


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
func (o *Conversationcategoriesentitylisting) String() string {
     o.Entities = []Conversationcategory{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcategoriesentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Conversationcategoriesentitylisting

    if ConversationcategoriesentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ConversationcategoriesentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Conversationcategory `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Conversationcategory{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

