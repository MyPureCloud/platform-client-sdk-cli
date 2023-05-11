package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetadataresultsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetadataresultsDud struct { 
    Entities []Metadataresultentity `json:"entities"`

}

// Metadataresults - List of resources created or modified as a result of running an accelerator
type Metadataresults struct { 
    

}

// String returns a JSON representation of the model
func (o *Metadataresults) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metadataresults) MarshalJSON() ([]byte, error) {
    type Alias Metadataresults

    if MetadataresultsMarshalled {
        return []byte("{}"), nil
    }
    MetadataresultsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

