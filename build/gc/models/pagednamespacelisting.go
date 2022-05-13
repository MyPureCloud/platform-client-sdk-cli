package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PagednamespacelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PagednamespacelistingDud struct { }

// Pagednamespacelisting
type Pagednamespacelisting struct { }

// String returns a JSON representation of the model
func (o *Pagednamespacelisting) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Pagednamespacelisting) MarshalJSON() ([]byte, error) {
    type Alias Pagednamespacelisting

    if PagednamespacelistingMarshalled {
        return []byte("{}"), nil
    }
    PagednamespacelistingMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

