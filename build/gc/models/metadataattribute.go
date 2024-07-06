package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetadataattributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetadataattributeDud struct { 
    Value string `json:"value"`

}

// Metadataattribute
type Metadataattribute struct { 
    

}

// String returns a JSON representation of the model
func (o *Metadataattribute) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metadataattribute) MarshalJSON() ([]byte, error) {
    type Alias Metadataattribute

    if MetadataattributeMarshalled {
        return []byte("{}"), nil
    }
    MetadataattributeMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

