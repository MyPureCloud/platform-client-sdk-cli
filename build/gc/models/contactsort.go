package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactsortMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactsortDud struct { 
    


    


    

}

// Contactsort
type Contactsort struct { 
    // FieldName
    FieldName string `json:"fieldName"`


    // Direction - The direction in which to sort contacts.
    Direction string `json:"direction"`


    // Numeric - Whether or not the column contains numeric data.
    Numeric bool `json:"numeric"`

}

// String returns a JSON representation of the model
func (o *Contactsort) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactsort) MarshalJSON() ([]byte, error) {
    type Alias Contactsort

    if ContactsortMarshalled {
        return []byte("{}"), nil
    }
    ContactsortMarshalled = true

    return json.Marshal(&struct {
        
        FieldName string `json:"fieldName"`
        
        Direction string `json:"direction"`
        
        Numeric bool `json:"numeric"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

