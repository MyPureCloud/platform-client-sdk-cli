package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ThirdpartysuggestionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ThirdpartysuggestionDud struct { 
    


    


    

}

// Thirdpartysuggestion
type Thirdpartysuggestion struct { 
    // Text - The third party suggestion text.
    Text string `json:"text"`


    // Title - The title of the suggestion.
    Title string `json:"title"`


    // Sources - A list of source references attributing the suggestion to its origin sources.
    Sources []Thirdpartysuggestionsource `json:"sources"`

}

// String returns a JSON representation of the model
func (o *Thirdpartysuggestion) String() string {
    
    
     o.Sources = []Thirdpartysuggestionsource{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Thirdpartysuggestion) MarshalJSON() ([]byte, error) {
    type Alias Thirdpartysuggestion

    if ThirdpartysuggestionMarshalled {
        return []byte("{}"), nil
    }
    ThirdpartysuggestionMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Title string `json:"title"`
        
        Sources []Thirdpartysuggestionsource `json:"sources"`
        *Alias
    }{

        


        


        
        Sources: []Thirdpartysuggestionsource{{}},
        

        Alias: (*Alias)(u),
    })
}

