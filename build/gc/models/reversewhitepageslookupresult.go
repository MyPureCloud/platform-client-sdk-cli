package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReversewhitepageslookupresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReversewhitepageslookupresultDud struct { 
    

}

// Reversewhitepageslookupresult
type Reversewhitepageslookupresult struct { 
    // Contacts
    Contacts []Externalcontact `json:"contacts"`

}

// String returns a JSON representation of the model
func (o *Reversewhitepageslookupresult) String() string {
     o.Contacts = []Externalcontact{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reversewhitepageslookupresult) MarshalJSON() ([]byte, error) {
    type Alias Reversewhitepageslookupresult

    if ReversewhitepageslookupresultMarshalled {
        return []byte("{}"), nil
    }
    ReversewhitepageslookupresultMarshalled = true

    return json.Marshal(&struct {
        
        Contacts []Externalcontact `json:"contacts"`
        *Alias
    }{

        
        Contacts: []Externalcontact{{}},
        

        Alias: (*Alias)(u),
    })
}

