package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotintentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotintentDud struct { 
    


    

}

// Botintent - A botConnector's bot intention
type Botintent struct { 
    // Name - The name of this intent.  This can be up to 100 characters long and must be comprised of displayable characters without leading or trailing whitespace
    Name string `json:"name"`


    // Slots - Optional returned data values associated with this intent, limit of 50.
    Slots map[string]Botslot `json:"slots"`

}

// String returns a JSON representation of the model
func (o *Botintent) String() string {
    
    
    
    
    
    
     o.Slots = map[string]Botslot{"": {}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botintent) MarshalJSON() ([]byte, error) {
    type Alias Botintent

    if BotintentMarshalled {
        return []byte("{}"), nil
    }
    BotintentMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Slots map[string]Botslot `json:"slots"`
        
        *Alias
    }{
        

        

        

        
        Slots: map[string]Botslot{"": {}},
        

        
        Alias: (*Alias)(u),
    })
}

