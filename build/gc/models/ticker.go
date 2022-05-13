package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TickerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TickerDud struct { 
    


    

}

// Ticker
type Ticker struct { 
    // Symbol - The ticker symbol for this organization. Example: ININ, AAPL, MSFT, etc.
    Symbol string `json:"symbol"`


    // Exchange - The exchange for this ticker symbol. Examples: NYSE, FTSE, NASDAQ, etc.
    Exchange string `json:"exchange"`

}

// String returns a JSON representation of the model
func (o *Ticker) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ticker) MarshalJSON() ([]byte, error) {
    type Alias Ticker

    if TickerMarshalled {
        return []byte("{}"), nil
    }
    TickerMarshalled = true

    return json.Marshal(&struct {
        
        Symbol string `json:"symbol"`
        
        Exchange string `json:"exchange"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

