package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetadatadocumentationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetadatadocumentationDud struct { 
    Description string `json:"description"`


    Location string `json:"location"`

}

// Metadatadocumentation - Additional documentation about an artifact
type Metadatadocumentation struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Metadatadocumentation) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metadatadocumentation) MarshalJSON() ([]byte, error) {
    type Alias Metadatadocumentation

    if MetadatadocumentationMarshalled {
        return []byte("{}"), nil
    }
    MetadatadocumentationMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

