package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallabletimesetentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallabletimesetentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Callabletimesetentitylisting
type Callabletimesetentitylisting struct { 
    // Entities
    Entities []Callabletimeset `json:"entities"`


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


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Callabletimesetentitylisting) String() string {
     o.Entities = []Callabletimeset{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callabletimesetentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Callabletimesetentitylisting

    if CallabletimesetentitylistingMarshalled {
        return []byte("{}"), nil
    }
    CallabletimesetentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Callabletimeset `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Callabletimeset{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

