package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MaskingruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MaskingruleDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`

}

// Maskingrule
type Maskingrule struct { 
    


    // Name - Masking rule name.
    Name string `json:"name"`


    // Description - Description about masking rule.
    Description string `json:"description"`


    // SubstituteCharacter - Replacement character for masked text character.
    SubstituteCharacter string `json:"substituteCharacter"`


    // Definition - Definition of masking rule (a valid regex or builtin AI based mask name).
    Definition string `json:"definition"`


    // Enabled - True/False
    Enabled bool `json:"enabled"`


    // VarType - Masking rule type
    VarType string `json:"type"`


    // Direction - inbound/outbound
    Direction string `json:"direction"`


    // Integrations - Associated integration channels
    Integrations []string `json:"integrations"`


    


    

}

// String returns a JSON representation of the model
func (o *Maskingrule) String() string {
    
    
    
    
    
    
    
     o.Integrations = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Maskingrule) MarshalJSON() ([]byte, error) {
    type Alias Maskingrule

    if MaskingruleMarshalled {
        return []byte("{}"), nil
    }
    MaskingruleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        SubstituteCharacter string `json:"substituteCharacter"`
        
        Definition string `json:"definition"`
        
        Enabled bool `json:"enabled"`
        
        VarType string `json:"type"`
        
        Direction string `json:"direction"`
        
        Integrations []string `json:"integrations"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        Integrations: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

