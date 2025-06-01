package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotentityDud struct { 
    


    

}

// Botentity - Description of a data value returned from an intent
type Botentity struct { 
    // Name - The name of the entity.
    Name string `json:"name"`


    // VarType - The data type of the entity.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Botentity) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botentity) MarshalJSON() ([]byte, error) {
    type Alias Botentity

    if BotentityMarshalled {
        return []byte("{}"), nil
    }
    BotentityMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

