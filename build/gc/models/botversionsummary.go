package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotversionsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotversionsummaryDud struct { 
    


    


    


    BotCompositeTag string `json:"botCompositeTag"`


    

}

// Botversionsummary - A version summary for a botConnector bot.
type Botversionsummary struct { 
    // Name - The name of the bot.
    Name string `json:"name"`


    // Id - The id of the bot.
    Id string `json:"id"`


    // Description - An optional description of the bot.
    Description string `json:"description"`


    


    // Version - The name of the version.
    Version string `json:"version"`

}

// String returns a JSON representation of the model
func (o *Botversionsummary) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botversionsummary) MarshalJSON() ([]byte, error) {
    type Alias Botversionsummary

    if BotversionsummaryMarshalled {
        return []byte("{}"), nil
    }
    BotversionsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Id string `json:"id"`
        
        Description string `json:"description"`
        
        Version string `json:"version"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

