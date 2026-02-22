package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Lexv2botaliasentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Lexv2botaliasentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Lexv2botaliasentitylisting
type Lexv2botaliasentitylisting struct { 
    // Entities
    Entities []Lexv2botalias `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Lexv2botaliasentitylisting) String() string {
     o.Entities = []Lexv2botalias{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lexv2botaliasentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Lexv2botaliasentitylisting

    if Lexv2botaliasentitylistingMarshalled {
        return []byte("{}"), nil
    }
    Lexv2botaliasentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Lexv2botalias `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Lexv2botalias{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

