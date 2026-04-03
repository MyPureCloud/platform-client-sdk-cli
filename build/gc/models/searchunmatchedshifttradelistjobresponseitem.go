package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SearchunmatchedshifttradelistjobresponseitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SearchunmatchedshifttradelistjobresponseitemDud struct { 
    

}

// Searchunmatchedshifttradelistjobresponseitem
type Searchunmatchedshifttradelistjobresponseitem struct { 
    // Trades - The shift trades that match the search criteria
    Trades []Searchunmatchedshifttraderesponseitem `json:"trades"`

}

// String returns a JSON representation of the model
func (o *Searchunmatchedshifttradelistjobresponseitem) String() string {
     o.Trades = []Searchunmatchedshifttraderesponseitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Searchunmatchedshifttradelistjobresponseitem) MarshalJSON() ([]byte, error) {
    type Alias Searchunmatchedshifttradelistjobresponseitem

    if SearchunmatchedshifttradelistjobresponseitemMarshalled {
        return []byte("{}"), nil
    }
    SearchunmatchedshifttradelistjobresponseitemMarshalled = true

    return json.Marshal(&struct {
        
        Trades []Searchunmatchedshifttraderesponseitem `json:"trades"`
        *Alias
    }{

        
        Trades: []Searchunmatchedshifttraderesponseitem{{}},
        

        Alias: (*Alias)(u),
    })
}

