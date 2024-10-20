package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserreferenceentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserreferenceentitylistingDud struct { 
    


    


    


    


    

}

// Userreferenceentitylisting
type Userreferenceentitylisting struct { 
    // Entities
    Entities []Userreference `json:"entities"`


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
func (o *Userreferenceentitylisting) String() string {
     o.Entities = []Userreference{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userreferenceentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Userreferenceentitylisting

    if UserreferenceentitylistingMarshalled {
        return []byte("{}"), nil
    }
    UserreferenceentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Userreference `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Userreference{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

