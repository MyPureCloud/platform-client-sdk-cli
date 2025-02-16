package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthorizationpolicyentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthorizationpolicyentitylistingDud struct { 
    


    


    


    

}

// Authorizationpolicyentitylisting
type Authorizationpolicyentitylisting struct { 
    // Entities
    Entities []Authorizationpolicy `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Authorizationpolicyentitylisting) String() string {
     o.Entities = []Authorizationpolicy{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authorizationpolicyentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Authorizationpolicyentitylisting

    if AuthorizationpolicyentitylistingMarshalled {
        return []byte("{}"), nil
    }
    AuthorizationpolicyentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Authorizationpolicy `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Authorizationpolicy{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

