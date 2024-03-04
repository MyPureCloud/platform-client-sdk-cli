package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetadataitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetadataitemDud struct { 
    


    

}

// Metadataitem
type Metadataitem struct { 
    // Text - The text contents of the metadata
    Text string `json:"text"`


    // Attributes - The custom attributes for the metadata
    Attributes map[string]string `json:"attributes"`

}

// String returns a JSON representation of the model
func (o *Metadataitem) String() string {
    
     o.Attributes = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metadataitem) MarshalJSON() ([]byte, error) {
    type Alias Metadataitem

    if MetadataitemMarshalled {
        return []byte("{}"), nil
    }
    MetadataitemMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Attributes map[string]string `json:"attributes"`
        *Alias
    }{

        


        
        Attributes: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

