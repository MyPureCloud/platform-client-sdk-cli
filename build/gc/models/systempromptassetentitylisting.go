package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SystempromptassetentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SystempromptassetentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Systempromptassetentitylisting
type Systempromptassetentitylisting struct { 
    // Entities
    Entities []Systempromptasset `json:"entities"`


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
func (o *Systempromptassetentitylisting) String() string {
     o.Entities = []Systempromptasset{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Systempromptassetentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Systempromptassetentitylisting

    if SystempromptassetentitylistingMarshalled {
        return []byte("{}"), nil
    }
    SystempromptassetentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Systempromptasset `json:"entities"`
        
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

        
        Entities: []Systempromptasset{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

