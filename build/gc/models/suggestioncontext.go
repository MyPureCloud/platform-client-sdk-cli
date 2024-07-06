package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestioncontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestioncontextDud struct { 
    Queue Addressableentityref `json:"queue"`


    MediaType string `json:"mediaType"`


    User Userreference `json:"user"`


    ExternalContact Addressableentityref `json:"externalContact"`


    Utterance Entity `json:"utterance"`


    Message Addressableentityref `json:"message"`


    QueryStatement string `json:"queryStatement"`

}

// Suggestioncontext
type Suggestioncontext struct { 
    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Suggestioncontext) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestioncontext) MarshalJSON() ([]byte, error) {
    type Alias Suggestioncontext

    if SuggestioncontextMarshalled {
        return []byte("{}"), nil
    }
    SuggestioncontextMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

