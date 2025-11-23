package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SecondarypresencelookupitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SecondarypresencelookupitemDud struct { 
    


    

}

// Secondarypresencelookupitem
type Secondarypresencelookupitem struct { 
    // LookupId - The lookupId of secondary presence id
    LookupId string `json:"lookupId"`


    // SecondaryPresence - The secondary presence Id
    SecondaryPresence Secondarypresence `json:"secondaryPresence"`

}

// String returns a JSON representation of the model
func (o *Secondarypresencelookupitem) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Secondarypresencelookupitem) MarshalJSON() ([]byte, error) {
    type Alias Secondarypresencelookupitem

    if SecondarypresencelookupitemMarshalled {
        return []byte("{}"), nil
    }
    SecondarypresencelookupitemMarshalled = true

    return json.Marshal(&struct {
        
        LookupId string `json:"lookupId"`
        
        SecondaryPresence Secondarypresence `json:"secondaryPresence"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

