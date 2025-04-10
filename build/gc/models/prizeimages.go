package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PrizeimagesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PrizeimagesDud struct { 
    ImageUrl string `json:"imageUrl"`

}

// Prizeimages
type Prizeimages struct { 
    

}

// String returns a JSON representation of the model
func (o *Prizeimages) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Prizeimages) MarshalJSON() ([]byte, error) {
    type Alias Prizeimages

    if PrizeimagesMarshalled {
        return []byte("{}"), nil
    }
    PrizeimagesMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

