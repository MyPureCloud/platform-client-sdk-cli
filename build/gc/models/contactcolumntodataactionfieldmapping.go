package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactcolumntodataactionfieldmappingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactcolumntodataactionfieldmappingDud struct { 
    


    

}

// Contactcolumntodataactionfieldmapping
type Contactcolumntodataactionfieldmapping struct { 
    // ContactColumnName - The name of a contact column whose data will be passed to the data action
    ContactColumnName string `json:"contactColumnName"`


    // DataActionField - The name of an input field from the data action that the contact column data will be passed to
    DataActionField string `json:"dataActionField"`

}

// String returns a JSON representation of the model
func (o *Contactcolumntodataactionfieldmapping) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactcolumntodataactionfieldmapping) MarshalJSON() ([]byte, error) {
    type Alias Contactcolumntodataactionfieldmapping

    if ContactcolumntodataactionfieldmappingMarshalled {
        return []byte("{}"), nil
    }
    ContactcolumntodataactionfieldmappingMarshalled = true

    return json.Marshal(&struct {
        
        ContactColumnName string `json:"contactColumnName"`
        
        DataActionField string `json:"dataActionField"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

