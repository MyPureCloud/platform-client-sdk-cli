package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DigitalrulesetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DigitalrulesetDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Digitalruleset
type Digitalruleset struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // ContactList - A ContactList to provide suggestions for contact columns on relevant conditions and actions.
    ContactList Domainentityref `json:"contactList"`


    // Rules - The list of rules.
    Rules []Digitalrule `json:"rules"`


    

}

// String returns a JSON representation of the model
func (o *Digitalruleset) String() string {
    
    
    
     o.Rules = []Digitalrule{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Digitalruleset) MarshalJSON() ([]byte, error) {
    type Alias Digitalruleset

    if DigitalrulesetMarshalled {
        return []byte("{}"), nil
    }
    DigitalrulesetMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        ContactList Domainentityref `json:"contactList"`
        
        Rules []Digitalrule `json:"rules"`
        *Alias
    }{

        


        


        


        


        


        


        
        Rules: []Digitalrule{{}},
        


        

        Alias: (*Alias)(u),
    })
}

