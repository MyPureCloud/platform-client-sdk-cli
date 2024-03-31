package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactsexportrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactsexportrequestDud struct { 
    


    


    

}

// Contactsexportrequest
type Contactsexportrequest struct { 
    // ContactListFilterId - Contact List Filter ID.
    ContactListFilterId string `json:"contactListFilterId"`


    // Criteria - Criteria to filter the contacts by.
    Criteria Contactbulksearchcriteria `json:"criteria"`


    // ContactIds - Contact IDs to be exported.
    ContactIds []string `json:"contactIds"`

}

// String returns a JSON representation of the model
func (o *Contactsexportrequest) String() string {
    
    
     o.ContactIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactsexportrequest) MarshalJSON() ([]byte, error) {
    type Alias Contactsexportrequest

    if ContactsexportrequestMarshalled {
        return []byte("{}"), nil
    }
    ContactsexportrequestMarshalled = true

    return json.Marshal(&struct {
        
        ContactListFilterId string `json:"contactListFilterId"`
        
        Criteria Contactbulksearchcriteria `json:"criteria"`
        
        ContactIds []string `json:"contactIds"`
        *Alias
    }{

        


        


        
        ContactIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

