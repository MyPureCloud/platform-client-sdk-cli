package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactsexportqueryconditionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactsexportqueryconditionsDud struct { 
    


    

}

// Contactsexportqueryconditions
type Contactsexportqueryconditions struct { 
    // Filters - Filters to apply on export
    Filters *Contactsexportfilter `json:"filters"`


    // Limit - Maximum result count in export, default is 180 000 000
    Limit int `json:"limit"`

}

// String returns a JSON representation of the model
func (o *Contactsexportqueryconditions) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactsexportqueryconditions) MarshalJSON() ([]byte, error) {
    type Alias Contactsexportqueryconditions

    if ContactsexportqueryconditionsMarshalled {
        return []byte("{}"), nil
    }
    ContactsexportqueryconditionsMarshalled = true

    return json.Marshal(&struct {
        
        Filters *Contactsexportfilter `json:"filters"`
        
        Limit int `json:"limit"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

