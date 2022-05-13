package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentcarouselMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentcarouselDud struct { 
    

}

// Contentcarousel - Carousel content object.
type Contentcarousel struct { 
    // Cards - An array of card objects.
    Cards []Contentcard `json:"cards"`

}

// String returns a JSON representation of the model
func (o *Contentcarousel) String() string {
     o.Cards = []Contentcard{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentcarousel) MarshalJSON() ([]byte, error) {
    type Alias Contentcarousel

    if ContentcarouselMarshalled {
        return []byte("{}"), nil
    }
    ContentcarouselMarshalled = true

    return json.Marshal(&struct {
        
        Cards []Contentcard `json:"cards"`
        *Alias
    }{

        
        Cards: []Contentcard{{}},
        

        Alias: (*Alias)(u),
    })
}

