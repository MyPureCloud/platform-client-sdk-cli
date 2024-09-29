package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PagedskillcombinationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PagedskillcombinationlistingDud struct { 
    


    


    


    


    

}

// Pagedskillcombinationlisting
type Pagedskillcombinationlisting struct { 
    // Entities
    Entities []Skillcombinationinfo `json:"entities"`


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
func (o *Pagedskillcombinationlisting) String() string {
     o.Entities = []Skillcombinationinfo{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Pagedskillcombinationlisting) MarshalJSON() ([]byte, error) {
    type Alias Pagedskillcombinationlisting

    if PagedskillcombinationlistingMarshalled {
        return []byte("{}"), nil
    }
    PagedskillcombinationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Skillcombinationinfo `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Skillcombinationinfo{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

