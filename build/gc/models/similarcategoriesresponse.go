package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SimilarcategoriesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SimilarcategoriesresponseDud struct { 
    

}

// Similarcategoriesresponse
type Similarcategoriesresponse struct { 
    // Entities
    Entities []Similarcategoryresult `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Similarcategoriesresponse) String() string {
     o.Entities = []Similarcategoryresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Similarcategoriesresponse) MarshalJSON() ([]byte, error) {
    type Alias Similarcategoriesresponse

    if SimilarcategoriesresponseMarshalled {
        return []byte("{}"), nil
    }
    SimilarcategoriesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Similarcategoryresult `json:"entities"`
        *Alias
    }{

        
        Entities: []Similarcategoryresult{{}},
        

        Alias: (*Alias)(u),
    })
}

