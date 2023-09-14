package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowhealthMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowhealthDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Flowhealth
type Flowhealth struct { 
    // FlowVersionInfo - Info about given flow version.
    FlowVersionInfo Flowhealthversioninfo `json:"flowVersionInfo"`


    // LanguageInfo - Each language's status about its health computation.
    LanguageInfo map[string]Localeinfo `json:"languageInfo"`


    // Intents - Health metrics information for the intents.
    Intents []Flowhealthintentinfo `json:"intents"`


    

}

// String returns a JSON representation of the model
func (o *Flowhealth) String() string {
    
     o.LanguageInfo = map[string]Localeinfo{"": {}} 
     o.Intents = []Flowhealthintentinfo{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowhealth) MarshalJSON() ([]byte, error) {
    type Alias Flowhealth

    if FlowhealthMarshalled {
        return []byte("{}"), nil
    }
    FlowhealthMarshalled = true

    return json.Marshal(&struct {
        
        FlowVersionInfo Flowhealthversioninfo `json:"flowVersionInfo"`
        
        LanguageInfo map[string]Localeinfo `json:"languageInfo"`
        
        Intents []Flowhealthintentinfo `json:"intents"`
        *Alias
    }{

        


        
        LanguageInfo: map[string]Localeinfo{"": {}},
        


        
        Intents: []Flowhealthintentinfo{{}},
        


        

        Alias: (*Alias)(u),
    })
}

