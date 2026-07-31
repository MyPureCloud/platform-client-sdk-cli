package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestionpatchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestionpatchrequestDud struct { 
    

}

// Suggestionpatchrequest
type Suggestionpatchrequest struct { 
    // ThirdPartySuggestion - The third-party suggestion to associate with the suggestion.
    ThirdPartySuggestion Thirdpartysuggestion `json:"thirdPartySuggestion"`

}

// String returns a JSON representation of the model
func (o *Suggestionpatchrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestionpatchrequest) MarshalJSON() ([]byte, error) {
    type Alias Suggestionpatchrequest

    if SuggestionpatchrequestMarshalled {
        return []byte("{}"), nil
    }
    SuggestionpatchrequestMarshalled = true

    return json.Marshal(&struct {
        
        ThirdPartySuggestion Thirdpartysuggestion `json:"thirdPartySuggestion"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

