package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotslotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotslotDud struct { 
    


    

}

// Botslot - Description of a data value returned from an intent
type Botslot struct { 
    // Name - The name of the slot. This can be up to 100 characters long and must be comprised of displayable characters without leading or trailing whitespace
    Name string `json:"name"`


    // VarType - The data type of the slot string, integer, decimal, duration, boolean, currency, datetime or the xxxCollection versions of those types
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Botslot) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botslot) MarshalJSON() ([]byte, error) {
    type Alias Botslot

    if BotslotMarshalled {
        return []byte("{}"), nil
    }
    BotslotMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

