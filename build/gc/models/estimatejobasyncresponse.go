package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EstimatejobasyncresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EstimatejobasyncresponseDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Estimatejobasyncresponse
type Estimatejobasyncresponse struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Estimatejobasyncresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Estimatejobasyncresponse) MarshalJSON() ([]byte, error) {
    type Alias Estimatejobasyncresponse

    if EstimatejobasyncresponseMarshalled {
        return []byte("{}"), nil
    }
    EstimatejobasyncresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

