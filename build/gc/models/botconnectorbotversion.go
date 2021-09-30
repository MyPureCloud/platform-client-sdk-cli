package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotconnectorbotversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotconnectorbotversionDud struct { 
    


    


    

}

// Botconnectorbotversion - A version description for a botConnector bot.
type Botconnectorbotversion struct { 
    // Version - The name of the version. This can be up to 100 characters long and must be comprised of displayable characters without leading or trailing whitespace
    Version string `json:"version"`


    // SupportedLanguages - The supported languages for this bot. EG 'en-us' or 'es', etc; These language codes are W3C language identification tags (ISO 639-1 for the language name and ISO 3166 for the country code)
    SupportedLanguages []string `json:"supportedLanguages"`


    // Intents - A list of potential intents this bot will return, limit of 50
    Intents []Botintent `json:"intents"`

}

// String returns a JSON representation of the model
func (o *Botconnectorbotversion) String() string {
    
    
    
    
    
    
     o.SupportedLanguages = []string{""} 
    
    
    
     o.Intents = []Botintent{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botconnectorbotversion) MarshalJSON() ([]byte, error) {
    type Alias Botconnectorbotversion

    if BotconnectorbotversionMarshalled {
        return []byte("{}"), nil
    }
    BotconnectorbotversionMarshalled = true

    return json.Marshal(&struct { 
        Version string `json:"version"`
        
        SupportedLanguages []string `json:"supportedLanguages"`
        
        Intents []Botintent `json:"intents"`
        
        *Alias
    }{
        

        

        

        
        SupportedLanguages: []string{""},
        

        

        
        Intents: []Botintent{{}},
        

        
        Alias: (*Alias)(u),
    })
}

