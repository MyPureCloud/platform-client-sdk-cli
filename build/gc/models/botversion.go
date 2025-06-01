package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotversionDud struct { 
    


    


    

}

// Botversion - A version description for a Bot.
type Botversion struct { 
    // Version - The name of the version.
    Version string `json:"version"`


    // SupportedLanguages - The supported languages for this bot. EG 'en-us' or 'es', etc; These language codes are W3C language identification tags (ISO 639-1 for the language name and ISO 3166 for the country code)
    SupportedLanguages []string `json:"supportedLanguages"`


    // Intents - A list of potential intents this bot will return, limit of 50
    Intents []Botconnectorintent `json:"intents"`

}

// String returns a JSON representation of the model
func (o *Botversion) String() string {
    
     o.SupportedLanguages = []string{""} 
     o.Intents = []Botconnectorintent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botversion) MarshalJSON() ([]byte, error) {
    type Alias Botversion

    if BotversionMarshalled {
        return []byte("{}"), nil
    }
    BotversionMarshalled = true

    return json.Marshal(&struct {
        
        Version string `json:"version"`
        
        SupportedLanguages []string `json:"supportedLanguages"`
        
        Intents []Botconnectorintent `json:"intents"`
        *Alias
    }{

        


        
        SupportedLanguages: []string{""},
        


        
        Intents: []Botconnectorintent{{}},
        

        Alias: (*Alias)(u),
    })
}

