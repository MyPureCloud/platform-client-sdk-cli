package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthzsubjectMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthzsubjectDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Authzsubject
type Authzsubject struct { 
    


    // Name
    Name string `json:"name"`


    // Grants
    Grants []Authzgrant `json:"grants"`


    // Version
    Version int `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Authzsubject) String() string {
    
     o.Grants = []Authzgrant{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authzsubject) MarshalJSON() ([]byte, error) {
    type Alias Authzsubject

    if AuthzsubjectMarshalled {
        return []byte("{}"), nil
    }
    AuthzsubjectMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Grants []Authzgrant `json:"grants"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        
        Grants: []Authzgrant{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

