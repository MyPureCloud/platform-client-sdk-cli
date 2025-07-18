package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CredentialinfocursorlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CredentialinfocursorlistingDud struct { 
    


    


    


    

}

// Credentialinfocursorlisting
type Credentialinfocursorlisting struct { 
    // Entities
    Entities []Credentialinfo `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Credentialinfocursorlisting) String() string {
     o.Entities = []Credentialinfo{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Credentialinfocursorlisting) MarshalJSON() ([]byte, error) {
    type Alias Credentialinfocursorlisting

    if CredentialinfocursorlistingMarshalled {
        return []byte("{}"), nil
    }
    CredentialinfocursorlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Credentialinfo `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Credentialinfo{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

