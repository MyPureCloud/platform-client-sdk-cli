package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LanguageoverrideMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LanguageoverrideDud struct { 
    


    


    

}

// Languageoverride
type Languageoverride struct { 
    // Language - The language code of the language being overridden
    Language string `json:"language"`


    // Engine - The ID of the TTS engine to use for this language override
    Engine string `json:"engine"`


    // Voice - The ID of the voice to use for this language override. The voice must be supported by the chosen engine.
    Voice string `json:"voice"`

}

// String returns a JSON representation of the model
func (o *Languageoverride) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Languageoverride) MarshalJSON() ([]byte, error) {
    type Alias Languageoverride

    if LanguageoverrideMarshalled {
        return []byte("{}"), nil
    }
    LanguageoverrideMarshalled = true

    return json.Marshal(&struct {
        
        Language string `json:"language"`
        
        Engine string `json:"engine"`
        
        Voice string `json:"voice"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

