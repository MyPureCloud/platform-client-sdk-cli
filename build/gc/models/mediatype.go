package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediatypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediatypeDud struct { 
    

}

// Mediatype - Media type definition
type Mediatype struct { 
    // VarType - The media type string as defined by RFC 2046. You can define specific types such as 'image/jpeg', 'video/mpeg', or specify wild cards for a range of types, 'image/_*', or all types '*_/_*'. See https://www.iana.org/assignments/media-types/media-types.xhtml for a list of registered media types.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Mediatype) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediatype) MarshalJSON() ([]byte, error) {
    type Alias Mediatype

    if MediatypeMarshalled {
        return []byte("{}"), nil
    }
    MediatypeMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

