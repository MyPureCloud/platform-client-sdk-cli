package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotDud struct { 
    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Bot - A bot instance
type Bot struct { 
    // Id - This is a string type that should denote a unique ID of the bot for calling purposes by the Genesys service (EG a UUID).
    Id string `json:"id"`


    // Name - This is the name that will be displayed to the user in Architect.
    Name string `json:"name"`


    // Description - An optional description of the bot.
    Description string `json:"description"`


    // Provider - The provider of the bot.
    Provider string `json:"provider"`


    // Versions - This bots version.
    Versions []Botversion `json:"versions"`


    // BotCompositeTag - A system-generated string that contains metadata about this bot.
    BotCompositeTag string `json:"botCompositeTag"`


    

}

// String returns a JSON representation of the model
func (o *Bot) String() string {
    
    
    
    
     o.Versions = []Botversion{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bot) MarshalJSON() ([]byte, error) {
    type Alias Bot

    if BotMarshalled {
        return []byte("{}"), nil
    }
    BotMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Provider string `json:"provider"`
        
        Versions []Botversion `json:"versions"`
        
        BotCompositeTag string `json:"botCompositeTag"`
        *Alias
    }{

        


        


        


        


        
        Versions: []Botversion{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

