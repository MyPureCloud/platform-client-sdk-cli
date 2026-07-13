package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationsexportlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationsexportlistingDud struct { 
    


    


    


    

}

// Organizationsexportlisting
type Organizationsexportlisting struct { 
    // Entities
    Entities []Organizationsexport `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Organizationsexportlisting) String() string {
     o.Entities = []Organizationsexport{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationsexportlisting) MarshalJSON() ([]byte, error) {
    type Alias Organizationsexportlisting

    if OrganizationsexportlistingMarshalled {
        return []byte("{}"), nil
    }
    OrganizationsexportlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Organizationsexport `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Organizationsexport{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

