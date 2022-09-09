package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DigitalrulesetentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DigitalrulesetentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Digitalrulesetentitylisting
type Digitalrulesetentitylisting struct { 
    // Entities
    Entities []Digitalruleset `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Digitalrulesetentitylisting) String() string {
     o.Entities = []Digitalruleset{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Digitalrulesetentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Digitalrulesetentitylisting

    if DigitalrulesetentitylistingMarshalled {
        return []byte("{}"), nil
    }
    DigitalrulesetentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Digitalruleset `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Digitalruleset{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

