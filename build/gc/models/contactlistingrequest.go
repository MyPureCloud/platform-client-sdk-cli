package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlistingrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlistingrequestDud struct { 
    


    


    


    


    

}

// Contactlistingrequest
type Contactlistingrequest struct { 
    // ContactListFilterId - Contact List Filter ID.
    ContactListFilterId string `json:"contactListFilterId"`


    // Criteria - Criteria to filter the contacts by.
    Criteria Contactbulksearchcriteria `json:"criteria"`


    // PageNumber - Page number.
    PageNumber int `json:"pageNumber"`


    // PageSize - Page size. The max that will be returned is 100.
    PageSize int `json:"pageSize"`


    // ContactSorts - The order in which to sort contacts.
    ContactSorts []Contactsort `json:"contactSorts"`

}

// String returns a JSON representation of the model
func (o *Contactlistingrequest) String() string {
    
    
    
    
     o.ContactSorts = []Contactsort{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlistingrequest) MarshalJSON() ([]byte, error) {
    type Alias Contactlistingrequest

    if ContactlistingrequestMarshalled {
        return []byte("{}"), nil
    }
    ContactlistingrequestMarshalled = true

    return json.Marshal(&struct {
        
        ContactListFilterId string `json:"contactListFilterId"`
        
        Criteria Contactbulksearchcriteria `json:"criteria"`
        
        PageNumber int `json:"pageNumber"`
        
        PageSize int `json:"pageSize"`
        
        ContactSorts []Contactsort `json:"contactSorts"`
        *Alias
    }{

        


        


        


        


        
        ContactSorts: []Contactsort{{}},
        

        Alias: (*Alias)(u),
    })
}

