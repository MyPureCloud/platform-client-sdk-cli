package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactbulkeditrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactbulkeditrequestDud struct { 
    


    


    


    

}

// Contactbulkeditrequest
type Contactbulkeditrequest struct { 
    // ContactListFilterId - Contact List Filter ID.
    ContactListFilterId string `json:"contactListFilterId"`


    // Criteria - Criteria to filter the contacts by.
    Criteria Contactbulksearchcriteria `json:"criteria"`


    // ContactIds - Contact IDs to be bulk edited.
    ContactIds []string `json:"contactIds"`


    // Contact - Contact object with details of fields used for patching.
    Contact Dialercontact `json:"contact"`

}

// String returns a JSON representation of the model
func (o *Contactbulkeditrequest) String() string {
    
    
     o.ContactIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactbulkeditrequest) MarshalJSON() ([]byte, error) {
    type Alias Contactbulkeditrequest

    if ContactbulkeditrequestMarshalled {
        return []byte("{}"), nil
    }
    ContactbulkeditrequestMarshalled = true

    return json.Marshal(&struct {
        
        ContactListFilterId string `json:"contactListFilterId"`
        
        Criteria Contactbulksearchcriteria `json:"criteria"`
        
        ContactIds []string `json:"contactIds"`
        
        Contact Dialercontact `json:"contact"`
        *Alias
    }{

        


        


        
        ContactIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

