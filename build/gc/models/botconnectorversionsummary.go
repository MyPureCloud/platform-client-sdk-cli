package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotconnectorversionsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotconnectorversionsummaryDud struct { 
    

}

// Botconnectorversionsummary - A version summary for a botConnector bot.
type Botconnectorversionsummary struct { 
    // Version - The name of the version.
    Version string `json:"version"`

}

// String returns a JSON representation of the model
func (o *Botconnectorversionsummary) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botconnectorversionsummary) MarshalJSON() ([]byte, error) {
    type Alias Botconnectorversionsummary

    if BotconnectorversionsummaryMarshalled {
        return []byte("{}"), nil
    }
    BotconnectorversionsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Version string `json:"version"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

