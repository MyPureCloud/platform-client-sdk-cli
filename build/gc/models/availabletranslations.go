package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AvailabletranslationsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AvailabletranslationsDud struct { 
    


    

}

// Availabletranslations
type Availabletranslations struct { 
    // OrgSpecific
    OrgSpecific []string `json:"orgSpecific"`


    // Builtin
    Builtin []string `json:"builtin"`

}

// String returns a JSON representation of the model
func (o *Availabletranslations) String() string {
     o.OrgSpecific = []string{""} 
     o.Builtin = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Availabletranslations) MarshalJSON() ([]byte, error) {
    type Alias Availabletranslations

    if AvailabletranslationsMarshalled {
        return []byte("{}"), nil
    }
    AvailabletranslationsMarshalled = true

    return json.Marshal(&struct {
        
        OrgSpecific []string `json:"orgSpecific"`
        
        Builtin []string `json:"builtin"`
        *Alias
    }{

        
        OrgSpecific: []string{""},
        


        
        Builtin: []string{""},
        

        Alias: (*Alias)(u),
    })
}

