package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetadataresultentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetadataresultentityDud struct { 
    VarType string `json:"type"`


    Description string `json:"description"`


    Visibility string `json:"visibility"`

}

// Metadataresultentity - A Genesys Cloud resource created or modified as a result of running an accelerator
type Metadataresultentity struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Metadataresultentity) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metadataresultentity) MarshalJSON() ([]byte, error) {
    type Alias Metadataresultentity

    if MetadataresultentityMarshalled {
        return []byte("{}"), nil
    }
    MetadataresultentityMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

