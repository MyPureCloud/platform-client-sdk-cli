package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MaskingrulevalidateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MaskingrulevalidateresponseDud struct { 
    Valid bool `json:"valid"`


    ValidationMessage string `json:"validationMessage"`


    MaskedText string `json:"maskedText"`

}

// Maskingrulevalidateresponse
type Maskingrulevalidateresponse struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Maskingrulevalidateresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Maskingrulevalidateresponse) MarshalJSON() ([]byte, error) {
    type Alias Maskingrulevalidateresponse

    if MaskingrulevalidateresponseMarshalled {
        return []byte("{}"), nil
    }
    MaskingrulevalidateresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

