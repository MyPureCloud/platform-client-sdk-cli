package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalmetricdatawriterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalmetricdatawriterequestDud struct { 
    

}

// Externalmetricdatawriterequest
type Externalmetricdatawriterequest struct { 
    // Items - A list of external metric data items. A maximum of 100 items are allowed.
    Items []Externalmetricdataitem `json:"items"`

}

// String returns a JSON representation of the model
func (o *Externalmetricdatawriterequest) String() string {
     o.Items = []Externalmetricdataitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalmetricdatawriterequest) MarshalJSON() ([]byte, error) {
    type Alias Externalmetricdatawriterequest

    if ExternalmetricdatawriterequestMarshalled {
        return []byte("{}"), nil
    }
    ExternalmetricdatawriterequestMarshalled = true

    return json.Marshal(&struct {
        
        Items []Externalmetricdataitem `json:"items"`
        *Alias
    }{

        
        Items: []Externalmetricdataitem{{}},
        

        Alias: (*Alias)(u),
    })
}

