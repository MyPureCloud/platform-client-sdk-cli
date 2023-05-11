package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulesdomainentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulesdomainentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Learningmodulesdomainentitylisting
type Learningmodulesdomainentitylisting struct { 
    // Entities
    Entities []Learningmodule `json:"entities"`


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


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Learningmodulesdomainentitylisting) String() string {
     o.Entities = []Learningmodule{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulesdomainentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulesdomainentitylisting

    if LearningmodulesdomainentitylistingMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulesdomainentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Learningmodule `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        PreviousUri string `json:"previousUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Learningmodule{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

