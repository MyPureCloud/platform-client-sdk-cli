package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotconnectorsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotconnectorsummaryDud struct { 
    


    


    


    


    

}

// Botconnectorsummary - A summary description for a bot
type Botconnectorsummary struct { 
    // Id - The id of the bot.
    Id string `json:"id"`


    // Name - The name of the bot.
    Name string `json:"name"`


    // Description - An optional description of the bot.
    Description string `json:"description"`


    // BotCompositeTag - A system-generated string that contains metadata about this bot.
    BotCompositeTag string `json:"botCompositeTag"`


    // Versions - This bots version.
    Versions []Botconnectorversionsummary `json:"versions"`

}

// String returns a JSON representation of the model
func (o *Botconnectorsummary) String() string {
    
    
    
    
     o.Versions = []Botconnectorversionsummary{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botconnectorsummary) MarshalJSON() ([]byte, error) {
    type Alias Botconnectorsummary

    if BotconnectorsummaryMarshalled {
        return []byte("{}"), nil
    }
    BotconnectorsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        BotCompositeTag string `json:"botCompositeTag"`
        
        Versions []Botconnectorversionsummary `json:"versions"`
        *Alias
    }{

        


        


        


        


        
        Versions: []Botconnectorversionsummary{{}},
        

        Alias: (*Alias)(u),
    })
}

