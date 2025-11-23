package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummarysettingentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummarysettingentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Summarysettingentitylisting
type Summarysettingentitylisting struct { 
    // Entities
    Entities []Summarysetting `json:"entities"`


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
func (o *Summarysettingentitylisting) String() string {
     o.Entities = []Summarysetting{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summarysettingentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Summarysettingentitylisting

    if SummarysettingentitylistingMarshalled {
        return []byte("{}"), nil
    }
    SummarysettingentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Summarysetting `json:"entities"`
        
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

        
        Entities: []Summarysetting{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

