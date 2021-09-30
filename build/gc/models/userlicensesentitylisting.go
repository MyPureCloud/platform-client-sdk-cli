package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserlicensesentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserlicensesentitylistingDud struct { 
    


    


    


    


    

}

// Userlicensesentitylisting
type Userlicensesentitylisting struct { 
    // Entities
    Entities []Userlicenses `json:"entities"`


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
func (o *Userlicensesentitylisting) String() string {
    
    
     o.Entities = []Userlicenses{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userlicensesentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Userlicensesentitylisting

    if UserlicensesentitylistingMarshalled {
        return []byte("{}"), nil
    }
    UserlicensesentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Userlicenses `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        *Alias
    }{
        

        
        Entities: []Userlicenses{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

