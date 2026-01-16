package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CarouselMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CarouselDud struct { }

// Carousel - A WhatsApp Carousel messaging template definition
type Carousel struct { }

// String returns a JSON representation of the model
func (o *Carousel) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Carousel) MarshalJSON() ([]byte, error) {
    type Alias Carousel

    if CarouselMarshalled {
        return []byte("{}"), nil
    }
    CarouselMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

