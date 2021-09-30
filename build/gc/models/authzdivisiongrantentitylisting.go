package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthzdivisiongrantentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthzdivisiongrantentitylistingDud struct { 
    


    


    


    


    

}

// Authzdivisiongrantentitylisting
type Authzdivisiongrantentitylisting struct { 
    // Entities
    Entities []Authzgrant `json:"entities"`


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
func (o *Authzdivisiongrantentitylisting) String() string {
    
    
     o.Entities = []Authzgrant{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authzdivisiongrantentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Authzdivisiongrantentitylisting

    if AuthzdivisiongrantentitylistingMarshalled {
        return []byte("{}"), nil
    }
    AuthzdivisiongrantentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Authzgrant `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        *Alias
    }{
        

        
        Entities: []Authzgrant{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

