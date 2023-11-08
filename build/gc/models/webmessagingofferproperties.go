package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingofferpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingofferpropertiesDud struct { 
    

}

// Webmessagingofferproperties
type Webmessagingofferproperties struct { 
    // OfferText - Text value to be used when inviting a visitor to engage with a web messaging offer.
    OfferText string `json:"offerText"`

}

// String returns a JSON representation of the model
func (o *Webmessagingofferproperties) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingofferproperties) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingofferproperties

    if WebmessagingofferpropertiesMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingofferpropertiesMarshalled = true

    return json.Marshal(&struct {
        
        OfferText string `json:"offerText"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

