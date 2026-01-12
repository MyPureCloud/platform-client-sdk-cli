package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomerintentassignmentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomerintentassignmentresponseDud struct { 
    CustomerIntent Domainentityref `json:"customerIntent"`


    Category Addressableentityref `json:"category"`


    DateMostRecentlyAssigned time.Time `json:"dateMostRecentlyAssigned"`

}

// Customerintentassignmentresponse
type Customerintentassignmentresponse struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Customerintentassignmentresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customerintentassignmentresponse) MarshalJSON() ([]byte, error) {
    type Alias Customerintentassignmentresponse

    if CustomerintentassignmentresponseMarshalled {
        return []byte("{}"), nil
    }
    CustomerintentassignmentresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

