package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotlistDud struct { 
    

}

// Botlist - A list of BotConnectorBots
type Botlist struct { 
    // ChatBots - A list of botConnector Bots. Max 50
    ChatBots []Botconnectorbot `json:"chatBots"`

}

// String returns a JSON representation of the model
func (o *Botlist) String() string {
    
    
     o.ChatBots = []Botconnectorbot{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botlist) MarshalJSON() ([]byte, error) {
    type Alias Botlist

    if BotlistMarshalled {
        return []byte("{}"), nil
    }
    BotlistMarshalled = true

    return json.Marshal(&struct { 
        ChatBots []Botconnectorbot `json:"chatBots"`
        
        *Alias
    }{
        

        
        ChatBots: []Botconnectorbot{{}},
        

        
        Alias: (*Alias)(u),
    })
}

