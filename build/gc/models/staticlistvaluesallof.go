package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StaticlistvaluesallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StaticlistvaluesallofDud struct { 
    


    

}

// Staticlistvaluesallof
type Staticlistvaluesallof struct { 
    // Items - Array of list items. Each item contains a value and optional synonyms.
    Items []Listitem `json:"items"`


    // MatchType - Defines how matching should work.
    MatchType string `json:"matchType"`

}

// String returns a JSON representation of the model
func (o *Staticlistvaluesallof) String() string {
     o.Items = []Listitem{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Staticlistvaluesallof) MarshalJSON() ([]byte, error) {
    type Alias Staticlistvaluesallof

    if StaticlistvaluesallofMarshalled {
        return []byte("{}"), nil
    }
    StaticlistvaluesallofMarshalled = true

    return json.Marshal(&struct {
        
        Items []Listitem `json:"items"`
        
        MatchType string `json:"matchType"`
        *Alias
    }{

        
        Items: []Listitem{{}},
        


        

        Alias: (*Alias)(u),
    })
}

