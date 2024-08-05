package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LimitcountlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LimitcountlistingDud struct { 
    


    


    

}

// Limitcountlisting
type Limitcountlisting struct { 
    // Entities
    Entities []Limitcount `json:"entities"`


    // NextUri - A URI to the next page in the listing.
    NextUri string `json:"nextUri"`


    // SelfUri - A URI to the current page in the listing.
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Limitcountlisting) String() string {
     o.Entities = []Limitcount{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Limitcountlisting) MarshalJSON() ([]byte, error) {
    type Alias Limitcountlisting

    if LimitcountlistingMarshalled {
        return []byte("{}"), nil
    }
    LimitcountlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Limitcount `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        
        Entities: []Limitcount{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

