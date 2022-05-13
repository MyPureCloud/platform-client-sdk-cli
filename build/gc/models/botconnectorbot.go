package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotconnectorbotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotconnectorbotDud struct { 
    


    


    


    


    BotCompositeTag string `json:"botCompositeTag"`


    SelfUri string `json:"selfUri"`

}

// Botconnectorbot - A botConnector Bot Instance
type Botconnectorbot struct { 
    // Id - The Botconnector Bot Id - this is configurable by the user when put
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // Description - An optional description of the bot.  This can be up to 256 characters long and must be comprised of displayable characters without leading or trailing whitespace
    Description string `json:"description"`


    // Versions - This bots versions, limit of 50 per bot
    Versions []Botconnectorbotversion `json:"versions"`


    


    

}

// String returns a JSON representation of the model
func (o *Botconnectorbot) String() string {
    
    
    
     o.Versions = []Botconnectorbotversion{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botconnectorbot) MarshalJSON() ([]byte, error) {
    type Alias Botconnectorbot

    if BotconnectorbotMarshalled {
        return []byte("{}"), nil
    }
    BotconnectorbotMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Versions []Botconnectorbotversion `json:"versions"`
        *Alias
    }{

        


        


        


        
        Versions: []Botconnectorbotversion{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

