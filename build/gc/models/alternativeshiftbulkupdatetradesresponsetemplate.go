package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshiftbulkupdatetradesresponsetemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshiftbulkupdatetradesresponsetemplateDud struct { 
    

}

// Alternativeshiftbulkupdatetradesresponsetemplate
type Alternativeshiftbulkupdatetradesresponsetemplate struct { 
    // Entities
    Entities []Alternativeshifttradebulkupdatetemplateitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Alternativeshiftbulkupdatetradesresponsetemplate) String() string {
     o.Entities = []Alternativeshifttradebulkupdatetemplateitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshiftbulkupdatetradesresponsetemplate) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshiftbulkupdatetradesresponsetemplate

    if AlternativeshiftbulkupdatetradesresponsetemplateMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshiftbulkupdatetradesresponsetemplateMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Alternativeshifttradebulkupdatetemplateitem `json:"entities"`
        *Alias
    }{

        
        Entities: []Alternativeshifttradebulkupdatetemplateitem{{}},
        

        Alias: (*Alias)(u),
    })
}

