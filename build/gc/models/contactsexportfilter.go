package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactsexportfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactsexportfilterDud struct { 
    


    


    


    


    

}

// Contactsexportfilter
type Contactsexportfilter struct { 
    // Eq - Filtered field should have the same value
    Eq Contactsexportfieldfilter `json:"eq"`


    // In - Filtered field should match one of the listed values
    In Contactsexportfieldlistfilter `json:"in"`


    // And - Boolean AND combination of filters
    And []Contactsexportfilter `json:"and"`


    // Or - Boolean OR combination of filters
    Or []Contactsexportfilter `json:"or"`


    // Not - Boolean negation of filters
    Not *Contactsexportfilter `json:"not"`

}

// String returns a JSON representation of the model
func (o *Contactsexportfilter) String() string {
    
    
     o.And = []Contactsexportfilter{{}} 
     o.Or = []Contactsexportfilter{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactsexportfilter) MarshalJSON() ([]byte, error) {
    type Alias Contactsexportfilter

    if ContactsexportfilterMarshalled {
        return []byte("{}"), nil
    }
    ContactsexportfilterMarshalled = true

    return json.Marshal(&struct {
        
        Eq Contactsexportfieldfilter `json:"eq"`
        
        In Contactsexportfieldlistfilter `json:"in"`
        
        And []Contactsexportfilter `json:"and"`
        
        Or []Contactsexportfilter `json:"or"`
        
        Not *Contactsexportfilter `json:"not"`
        *Alias
    }{

        


        


        
        And: []Contactsexportfilter{{}},
        


        
        Or: []Contactsexportfilter{{}},
        


        

        Alias: (*Alias)(u),
    })
}

