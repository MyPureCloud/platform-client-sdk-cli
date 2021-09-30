package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DatatableimportentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DatatableimportentitylistingDud struct { 
    


    


    


    


    

}

// Datatableimportentitylisting
type Datatableimportentitylisting struct { 
    // Entities
    Entities []Datatableimportjob `json:"entities"`


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
func (o *Datatableimportentitylisting) String() string {
    
    
     o.Entities = []Datatableimportjob{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Datatableimportentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Datatableimportentitylisting

    if DatatableimportentitylistingMarshalled {
        return []byte("{}"), nil
    }
    DatatableimportentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Datatableimportjob `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        *Alias
    }{
        

        
        Entities: []Datatableimportjob{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

