package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingofferfieldsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingofferfieldsDud struct { 
    


    

}

// Webmessagingofferfields
type Webmessagingofferfields struct { 
    // OfferText - Text value to be used when inviting a visitor to engage with a web messaging offer.
    OfferText string `json:"offerText"`


    // ArchitectFlow - Flow to be invoked, overrides default flow when specified.
    ArchitectFlow Addressableentityref `json:"architectFlow"`

}

// String returns a JSON representation of the model
func (o *Webmessagingofferfields) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingofferfields) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingofferfields

    if WebmessagingofferfieldsMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingofferfieldsMarshalled = true

    return json.Marshal(&struct { 
        OfferText string `json:"offerText"`
        
        ArchitectFlow Addressableentityref `json:"architectFlow"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

