package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiessttengineMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiessttengineDud struct { 
    


    


    


    ReplacedBy string `json:"replacedBy"`


    

}

// Capabilitiessttengine
type Capabilitiessttengine struct { 
    // Name
    Name string `json:"name"`


    // IsGrammarBased
    IsGrammarBased bool `json:"isGrammarBased"`


    // IsDefault
    IsDefault bool `json:"isDefault"`


    


    // SupportedLanguages
    SupportedLanguages []string `json:"supportedLanguages"`

}

// String returns a JSON representation of the model
func (o *Capabilitiessttengine) String() string {
    
    
    
     o.SupportedLanguages = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiessttengine) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiessttengine

    if CapabilitiessttengineMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiessttengineMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        IsGrammarBased bool `json:"isGrammarBased"`
        
        IsDefault bool `json:"isDefault"`
        
        SupportedLanguages []string `json:"supportedLanguages"`
        *Alias
    }{

        


        


        


        


        
        SupportedLanguages: []string{""},
        

        Alias: (*Alias)(u),
    })
}

