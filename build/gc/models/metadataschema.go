package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetadataschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetadataschemaDud struct { 
    Title string `json:"title"`


    Description string `json:"description"`


    VarType string `json:"type"`


    Properties []map[string]Metadataproperty `json:"properties"`


    Required []string `json:"required"`

}

// Metadataschema - A description of the contents of a data gathering interface for an accelerator
type Metadataschema struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Metadataschema) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metadataschema) MarshalJSON() ([]byte, error) {
    type Alias Metadataschema

    if MetadataschemaMarshalled {
        return []byte("{}"), nil
    }
    MetadataschemaMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

