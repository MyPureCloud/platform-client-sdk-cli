package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactaddressconditionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactaddressconditionsettingsDud struct { 
    


    

}

// Contactaddressconditionsettings
type Contactaddressconditionsettings struct { 
    // Operator - The operator to use when comparing address values.
    Operator string `json:"operator"`


    // Value - The value to compare against the contact's address.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Contactaddressconditionsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactaddressconditionsettings) MarshalJSON() ([]byte, error) {
    type Alias Contactaddressconditionsettings

    if ContactaddressconditionsettingsMarshalled {
        return []byte("{}"), nil
    }
    ContactaddressconditionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

