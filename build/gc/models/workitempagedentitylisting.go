package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitempagedentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitempagedentitylistingDud struct { 
    


    


    


    


    

}

// Workitempagedentitylisting
type Workitempagedentitylisting struct { 
    // Entities
    Entities []Workitem `json:"entities"`


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
func (o *Workitempagedentitylisting) String() string {
     o.Entities = []Workitem{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitempagedentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Workitempagedentitylisting

    if WorkitempagedentitylistingMarshalled {
        return []byte("{}"), nil
    }
    WorkitempagedentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workitem `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Workitem{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

