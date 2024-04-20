package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NluorganizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NluorganizationDud struct { 
    


    

}

// Nluorganization
type Nluorganization struct { 
    // Limits - The NLU limits defined for this Organization
    Limits map[string]int `json:"limits"`


    // SupportedLanguagesInfo - The list of Supported features for each languages for this Organization
    SupportedLanguagesInfo []Supportedlanguagesinfodefinition `json:"supportedLanguagesInfo"`

}

// String returns a JSON representation of the model
func (o *Nluorganization) String() string {
     o.Limits = map[string]int{"": 0} 
     o.SupportedLanguagesInfo = []Supportedlanguagesinfodefinition{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nluorganization) MarshalJSON() ([]byte, error) {
    type Alias Nluorganization

    if NluorganizationMarshalled {
        return []byte("{}"), nil
    }
    NluorganizationMarshalled = true

    return json.Marshal(&struct {
        
        Limits map[string]int `json:"limits"`
        
        SupportedLanguagesInfo []Supportedlanguagesinfodefinition `json:"supportedLanguagesInfo"`
        *Alias
    }{

        
        Limits: map[string]int{"": 0},
        


        
        SupportedLanguagesInfo: []Supportedlanguagesinfodefinition{{}},
        

        Alias: (*Alias)(u),
    })
}

