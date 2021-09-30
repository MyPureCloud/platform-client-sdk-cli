package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TtssettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TtssettingsDud struct { 
    


    

}

// Ttssettings
type Ttssettings struct { 
    // DefaultEngine - ID of the global default TTS engine
    DefaultEngine string `json:"defaultEngine"`


    // LanguageOverrides - The list of default overrides for specific languages
    LanguageOverrides []Languageoverride `json:"languageOverrides"`

}

// String returns a JSON representation of the model
func (o *Ttssettings) String() string {
    
    
    
    
    
    
     o.LanguageOverrides = []Languageoverride{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ttssettings) MarshalJSON() ([]byte, error) {
    type Alias Ttssettings

    if TtssettingsMarshalled {
        return []byte("{}"), nil
    }
    TtssettingsMarshalled = true

    return json.Marshal(&struct { 
        DefaultEngine string `json:"defaultEngine"`
        
        LanguageOverrides []Languageoverride `json:"languageOverrides"`
        
        *Alias
    }{
        

        

        

        
        LanguageOverrides: []Languageoverride{{}},
        

        
        Alias: (*Alias)(u),
    })
}

