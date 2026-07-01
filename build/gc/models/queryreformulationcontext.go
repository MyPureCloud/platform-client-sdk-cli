package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryreformulationcontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryreformulationcontextDud struct { 
    KnowledgeClassification string `json:"knowledgeClassification"`


    ReformulatedQuery string `json:"reformulatedQuery"`

}

// Queryreformulationcontext
type Queryreformulationcontext struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Queryreformulationcontext) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryreformulationcontext) MarshalJSON() ([]byte, error) {
    type Alias Queryreformulationcontext

    if QueryreformulationcontextMarshalled {
        return []byte("{}"), nil
    }
    QueryreformulationcontextMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

