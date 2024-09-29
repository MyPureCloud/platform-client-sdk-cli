package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeintegrationoptionsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeintegrationoptionsresponseDud struct { 
    

}

// Knowledgeintegrationoptionsresponse
type Knowledgeintegrationoptionsresponse struct { 
    // Filters - Filter setting options available for a knowledge source.
    Filters []Knowledgeintegrationfilter `json:"filters"`

}

// String returns a JSON representation of the model
func (o *Knowledgeintegrationoptionsresponse) String() string {
     o.Filters = []Knowledgeintegrationfilter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeintegrationoptionsresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeintegrationoptionsresponse

    if KnowledgeintegrationoptionsresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeintegrationoptionsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Filters []Knowledgeintegrationfilter `json:"filters"`
        *Alias
    }{

        
        Filters: []Knowledgeintegrationfilter{{}},
        

        Alias: (*Alias)(u),
    })
}

