package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationroleentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationroleentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Organizationroleentitylisting
type Organizationroleentitylisting struct { 
    // Entities
    Entities []Domainorganizationrole `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Organizationroleentitylisting) String() string {
     o.Entities = []Domainorganizationrole{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationroleentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Organizationroleentitylisting

    if OrganizationroleentitylistingMarshalled {
        return []byte("{}"), nil
    }
    OrganizationroleentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Domainorganizationrole `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        LastUri string `json:"lastUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Domainorganizationrole{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

