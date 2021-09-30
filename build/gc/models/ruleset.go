package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RulesetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RulesetDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Ruleset
type Ruleset struct { 
    


    // Name - The name of the RuleSet.
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // ContactList - A ContactList to provide user-interface suggestions for contact columns on relevant conditions and actions.
    ContactList Domainentityref `json:"contactList"`


    // Queue - A Queue to provide user-interface suggestions for wrap-up codes on relevant conditions and actions.
    Queue Domainentityref `json:"queue"`


    // Rules - The list of rules.
    Rules []Dialerrule `json:"rules"`


    

}

// String returns a JSON representation of the model
func (o *Ruleset) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Rules = []Dialerrule{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ruleset) MarshalJSON() ([]byte, error) {
    type Alias Ruleset

    if RulesetMarshalled {
        return []byte("{}"), nil
    }
    RulesetMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        
        
        Version int `json:"version"`
        
        ContactList Domainentityref `json:"contactList"`
        
        Queue Domainentityref `json:"queue"`
        
        Rules []Dialerrule `json:"rules"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Rules: []Dialerrule{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

