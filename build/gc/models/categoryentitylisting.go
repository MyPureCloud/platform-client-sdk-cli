package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CategoryentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CategoryentitylistingDud struct { 
    


    


    


    


    

}

// Categoryentitylisting
type Categoryentitylisting struct { 
    // Entities
    Entities []Category `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Categoryentitylisting) String() string {
    
    
     o.Entities = []Category{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Categoryentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Categoryentitylisting

    if CategoryentitylistingMarshalled {
        return []byte("{}"), nil
    }
    CategoryentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Category `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        *Alias
    }{
        

        
        Entities: []Category{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

