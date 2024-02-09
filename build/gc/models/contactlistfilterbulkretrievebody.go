package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlistfilterbulkretrievebodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlistfilterbulkretrievebodyDud struct { 
    

}

// Contactlistfilterbulkretrievebody
type Contactlistfilterbulkretrievebody struct { 
    // Ids - The IDs of the Contact List Filters to retrieve.
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Contactlistfilterbulkretrievebody) String() string {
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlistfilterbulkretrievebody) MarshalJSON() ([]byte, error) {
    type Alias Contactlistfilterbulkretrievebody

    if ContactlistfilterbulkretrievebodyMarshalled {
        return []byte("{}"), nil
    }
    ContactlistfilterbulkretrievebodyMarshalled = true

    return json.Marshal(&struct {
        
        Ids []string `json:"ids"`
        *Alias
    }{

        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}

