package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserlanguageentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserlanguageentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Userlanguageentitylisting
type Userlanguageentitylisting struct { 
    // Entities
    Entities []Userroutinglanguage `json:"entities"`


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
func (o *Userlanguageentitylisting) String() string {
     o.Entities = []Userroutinglanguage{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userlanguageentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Userlanguageentitylisting

    if UserlanguageentitylistingMarshalled {
        return []byte("{}"), nil
    }
    UserlanguageentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Userroutinglanguage `json:"entities"`
        
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

        
        Entities: []Userroutinglanguage{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

