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
type ContactcolumntodataactionfieldmappingDud struct { }

// Contactcolumntodataactionfieldmapping
type Contactcolumntodataactionfieldmapping struct { }

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
        *Alias
    }{
        
        Alias: (*Alias)(u),
    })
}

