package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomerintentusageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomerintentusageDud struct { 
    EntityType string `json:"entityType"`


    SourceType string `json:"sourceType"`


    Source Domainentityref `json:"source"`


    DateCreated time.Time `json:"dateCreated"`

}

// Customerintentusage - A place where a customer intent or category is used
type Customerintentusage struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Customerintentusage) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customerintentusage) MarshalJSON() ([]byte, error) {
    type Alias Customerintentusage

    if CustomerintentusageMarshalled {
        return []byte("{}"), nil
    }
    CustomerintentusageMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

