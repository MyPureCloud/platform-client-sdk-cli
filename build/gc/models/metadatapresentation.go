package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetadatapresentationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetadatapresentationDud struct { 
    Title string `json:"title"`


    Schema Metadataschema `json:"schema"`

}

// Metadatapresentation - A representation of data fields to be gathered for installing the accelerator
type Metadatapresentation struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Metadatapresentation) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metadatapresentation) MarshalJSON() ([]byte, error) {
    type Alias Metadatapresentation

    if MetadatapresentationMarshalled {
        return []byte("{}"), nil
    }
    MetadatapresentationMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

