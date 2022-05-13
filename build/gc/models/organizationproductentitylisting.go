package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationproductentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationproductentitylistingDud struct { 
    


    


    


    


    

}

// Organizationproductentitylisting
type Organizationproductentitylisting struct { 
    // Entities
    Entities []Domainorganizationproduct `json:"entities"`


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
func (o *Organizationproductentitylisting) String() string {
     o.Entities = []Domainorganizationproduct{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationproductentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Organizationproductentitylisting

    if OrganizationproductentitylistingMarshalled {
        return []byte("{}"), nil
    }
    OrganizationproductentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Domainorganizationproduct `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Domainorganizationproduct{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

