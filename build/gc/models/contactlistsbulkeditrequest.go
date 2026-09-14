package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlistsbulkeditrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlistsbulkeditrequestDud struct { 
    


    

}

// Contactlistsbulkeditrequest
type Contactlistsbulkeditrequest struct { 
    // ContactListIds - Contact List IDs to be bulk edited.
    ContactListIds []string `json:"contactListIds"`


    // ContactList - Contact list object with details of fields used for patching. Accepted fields: retentionType, retentionDays, timeZone
    ContactList Contactlist `json:"contactList"`

}

// String returns a JSON representation of the model
func (o *Contactlistsbulkeditrequest) String() string {
     o.ContactListIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlistsbulkeditrequest) MarshalJSON() ([]byte, error) {
    type Alias Contactlistsbulkeditrequest

    if ContactlistsbulkeditrequestMarshalled {
        return []byte("{}"), nil
    }
    ContactlistsbulkeditrequestMarshalled = true

    return json.Marshal(&struct {
        
        ContactListIds []string `json:"contactListIds"`
        
        ContactList Contactlist `json:"contactList"`
        *Alias
    }{

        
        ContactListIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

