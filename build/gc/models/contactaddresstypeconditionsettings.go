package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactaddresstypeconditionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactaddresstypeconditionsettingsDud struct { 
    


    

}

// Contactaddresstypeconditionsettings
type Contactaddresstypeconditionsettings struct { 
    // Operator - The operator to use when comparing the address types.
    Operator string `json:"operator"`


    // Value - The type value to compare against the contact column type.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Contactaddresstypeconditionsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactaddresstypeconditionsettings) MarshalJSON() ([]byte, error) {
    type Alias Contactaddresstypeconditionsettings

    if ContactaddresstypeconditionsettingsMarshalled {
        return []byte("{}"), nil
    }
    ContactaddresstypeconditionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

