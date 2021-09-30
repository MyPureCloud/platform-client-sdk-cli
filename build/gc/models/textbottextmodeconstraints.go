package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbottextmodeconstraintsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbottextmodeconstraintsDud struct { 
    

}

// Textbottextmodeconstraints - Mode constraints to observe when operating on a bot flow.
type Textbottextmodeconstraints struct { 
    // LanguagePreferences - The list of language preferences by their ISO language code.
    LanguagePreferences []string `json:"languagePreferences"`

}

// String returns a JSON representation of the model
func (o *Textbottextmodeconstraints) String() string {
    
    
     o.LanguagePreferences = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbottextmodeconstraints) MarshalJSON() ([]byte, error) {
    type Alias Textbottextmodeconstraints

    if TextbottextmodeconstraintsMarshalled {
        return []byte("{}"), nil
    }
    TextbottextmodeconstraintsMarshalled = true

    return json.Marshal(&struct { 
        LanguagePreferences []string `json:"languagePreferences"`
        
        *Alias
    }{
        

        
        LanguagePreferences: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

