package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GenesysbotconnectorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GenesysbotconnectorDud struct { 
    

}

// Genesysbotconnector
type Genesysbotconnector struct { 
    // QueryParameters - User defined name/value parameters passed to the BotConnector bot.
    QueryParameters map[string]string `json:"queryParameters"`

}

// String returns a JSON representation of the model
func (o *Genesysbotconnector) String() string {
     o.QueryParameters = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Genesysbotconnector) MarshalJSON() ([]byte, error) {
    type Alias Genesysbotconnector

    if GenesysbotconnectorMarshalled {
        return []byte("{}"), nil
    }
    GenesysbotconnectorMarshalled = true

    return json.Marshal(&struct {
        
        QueryParameters map[string]string `json:"queryParameters"`
        *Alias
    }{

        
        QueryParameters: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

