package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NamedentitytypeitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NamedentitytypeitemDud struct { 
    


    


    

}

// Namedentitytypeitem
type Namedentitytypeitem struct { 
    // Value - A value for an named entity type definition.
    Value string `json:"value"`


    // Synonyms - Synonyms for the given named entity value.
    Synonyms []string `json:"synonyms"`


    // AdditionalLanguages - Additional Language Synonyms for the given named entity value.
    AdditionalLanguages map[string]Additionallanguagessynonyms `json:"additionalLanguages"`

}

// String returns a JSON representation of the model
func (o *Namedentitytypeitem) String() string {
    
     o.Synonyms = []string{""} 
     o.AdditionalLanguages = map[string]Additionallanguagessynonyms{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Namedentitytypeitem) MarshalJSON() ([]byte, error) {
    type Alias Namedentitytypeitem

    if NamedentitytypeitemMarshalled {
        return []byte("{}"), nil
    }
    NamedentitytypeitemMarshalled = true

    return json.Marshal(&struct {
        
        Value string `json:"value"`
        
        Synonyms []string `json:"synonyms"`
        
        AdditionalLanguages map[string]Additionallanguagessynonyms `json:"additionalLanguages"`
        *Alias
    }{

        


        
        Synonyms: []string{""},
        


        
        AdditionalLanguages: map[string]Additionallanguagessynonyms{"": {}},
        

        Alias: (*Alias)(u),
    })
}

