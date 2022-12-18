package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DivspermittedentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DivspermittedentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Divspermittedentitylisting
type Divspermittedentitylisting struct { 
    // Entities
    Entities []Authzdivision `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // AllDivsPermitted
    AllDivsPermitted bool `json:"allDivsPermitted"`


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
func (o *Divspermittedentitylisting) String() string {
     o.Entities = []Authzdivision{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Divspermittedentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Divspermittedentitylisting

    if DivspermittedentitylistingMarshalled {
        return []byte("{}"), nil
    }
    DivspermittedentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Authzdivision `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        AllDivsPermitted bool `json:"allDivsPermitted"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        PreviousUri string `json:"previousUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Authzdivision{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

