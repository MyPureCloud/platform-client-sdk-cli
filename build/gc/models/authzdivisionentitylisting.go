package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthzdivisionentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthzdivisionentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Authzdivisionentitylisting
type Authzdivisionentitylisting struct { 
    // Entities
    Entities []Authzdivision `json:"entities"`


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
func (o *Authzdivisionentitylisting) String() string {
     o.Entities = []Authzdivision{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authzdivisionentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Authzdivisionentitylisting

    if AuthzdivisionentitylistingMarshalled {
        return []byte("{}"), nil
    }
    AuthzdivisionentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Authzdivision `json:"entities"`
        
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

        
        Entities: []Authzdivision{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

