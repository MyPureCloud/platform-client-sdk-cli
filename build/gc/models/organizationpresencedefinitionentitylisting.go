package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationpresencedefinitionentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationpresencedefinitionentitylistingDud struct { 
    


    


    

}

// Organizationpresencedefinitionentitylisting
type Organizationpresencedefinitionentitylisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Organizationpresencedefinition `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Organizationpresencedefinitionentitylisting) String() string {
    
     o.Entities = []Organizationpresencedefinition{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationpresencedefinitionentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Organizationpresencedefinitionentitylisting

    if OrganizationpresencedefinitionentitylistingMarshalled {
        return []byte("{}"), nil
    }
    OrganizationpresencedefinitionentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Organizationpresencedefinition `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Organizationpresencedefinition{{}},
        


        

        Alias: (*Alias)(u),
    })
}

