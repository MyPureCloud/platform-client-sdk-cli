package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LexbotaliasentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LexbotaliasentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Lexbotaliasentitylisting
type Lexbotaliasentitylisting struct { 
    // Entities
    Entities []Lexbotalias `json:"entities"`


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
func (o *Lexbotaliasentitylisting) String() string {
     o.Entities = []Lexbotalias{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lexbotaliasentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Lexbotaliasentitylisting

    if LexbotaliasentitylistingMarshalled {
        return []byte("{}"), nil
    }
    LexbotaliasentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Lexbotalias `json:"entities"`
        
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

        
        Entities: []Lexbotalias{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

