package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestioncannedresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestioncannedresponseDud struct { 
    Response Addressableentityref `json:"response"`


    Library Addressableentityref `json:"library"`

}

// Suggestioncannedresponse
type Suggestioncannedresponse struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Suggestioncannedresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestioncannedresponse) MarshalJSON() ([]byte, error) {
    type Alias Suggestioncannedresponse

    if SuggestioncannedresponseMarshalled {
        return []byte("{}"), nil
    }
    SuggestioncannedresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

