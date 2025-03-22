package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactsexportfieldlistfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactsexportfieldlistfilterDud struct { 
    


    

}

// Contactsexportfieldlistfilter
type Contactsexportfieldlistfilter struct { 
    // Field - Field name to apply the filter
    Field string `json:"field"`


    // Values - Values to check field's value against
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Contactsexportfieldlistfilter) String() string {
    
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactsexportfieldlistfilter) MarshalJSON() ([]byte, error) {
    type Alias Contactsexportfieldlistfilter

    if ContactsexportfieldlistfilterMarshalled {
        return []byte("{}"), nil
    }
    ContactsexportfieldlistfilterMarshalled = true

    return json.Marshal(&struct {
        
        Field string `json:"field"`
        
        Values []string `json:"values"`
        *Alias
    }{

        


        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

