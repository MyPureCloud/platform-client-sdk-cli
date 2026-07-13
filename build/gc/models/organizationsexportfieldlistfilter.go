package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationsexportfieldlistfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationsexportfieldlistfilterDud struct { 
    


    

}

// Organizationsexportfieldlistfilter
type Organizationsexportfieldlistfilter struct { 
    // Field - Field name to apply the filter
    Field string `json:"field"`


    // Values - Values to check field's value against
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Organizationsexportfieldlistfilter) String() string {
    
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationsexportfieldlistfilter) MarshalJSON() ([]byte, error) {
    type Alias Organizationsexportfieldlistfilter

    if OrganizationsexportfieldlistfilterMarshalled {
        return []byte("{}"), nil
    }
    OrganizationsexportfieldlistfilterMarshalled = true

    return json.Marshal(&struct {
        
        Field string `json:"field"`
        
        Values []string `json:"values"`
        *Alias
    }{

        


        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

