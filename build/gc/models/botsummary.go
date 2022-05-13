package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotsummaryDud struct { 
    


    


    


    BotCompositeTag string `json:"botCompositeTag"`

}

// Botsummary - A summary description for a botConnector bot
type Botsummary struct { 
    // Name - The name of the bot.
    Name string `json:"name"`


    // Id - The id of the bot.
    Id string `json:"id"`


    // Description - An optional description of the bot.
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Botsummary) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botsummary) MarshalJSON() ([]byte, error) {
    type Alias Botsummary

    if BotsummaryMarshalled {
        return []byte("{}"), nil
    }
    BotsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Id string `json:"id"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

