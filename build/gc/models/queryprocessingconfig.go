package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryprocessingconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryprocessingconfigDud struct { 
    

}

// Queryprocessingconfig
type Queryprocessingconfig struct { 
    // KnowledgeQueryProcessing - Knowledge query processing mode applied before retrieval.
    KnowledgeQueryProcessing string `json:"knowledgeQueryProcessing"`

}

// String returns a JSON representation of the model
func (o *Queryprocessingconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryprocessingconfig) MarshalJSON() ([]byte, error) {
    type Alias Queryprocessingconfig

    if QueryprocessingconfigMarshalled {
        return []byte("{}"), nil
    }
    QueryprocessingconfigMarshalled = true

    return json.Marshal(&struct {
        
        KnowledgeQueryProcessing string `json:"knowledgeQueryProcessing"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

