package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeanswerconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeanswerconfigDud struct { 
    

}

// Knowledgeanswerconfig
type Knowledgeanswerconfig struct { 
    // Enabled - Knowledge answer is enabled.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Knowledgeanswerconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeanswerconfig) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeanswerconfig

    if KnowledgeanswerconfigMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeanswerconfigMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

