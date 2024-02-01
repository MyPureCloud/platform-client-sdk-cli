package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlisttemplatebulkretrievebodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlisttemplatebulkretrievebodyDud struct { 
    

}

// Contactlisttemplatebulkretrievebody
type Contactlisttemplatebulkretrievebody struct { 
    // Ids - The IDs of the Contact List Templates to retrieve.
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Contactlisttemplatebulkretrievebody) String() string {
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlisttemplatebulkretrievebody) MarshalJSON() ([]byte, error) {
    type Alias Contactlisttemplatebulkretrievebody

    if ContactlisttemplatebulkretrievebodyMarshalled {
        return []byte("{}"), nil
    }
    ContactlisttemplatebulkretrievebodyMarshalled = true

    return json.Marshal(&struct {
        
        Ids []string `json:"ids"`
        *Alias
    }{

        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}

