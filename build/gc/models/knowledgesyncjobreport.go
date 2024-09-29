package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesyncjobreportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesyncjobreportDud struct { 
    


    

}

// Knowledgesyncjobreport
type Knowledgesyncjobreport struct { 
    // Errors - List of errors occurred during processing sync.
    Errors []Errorbody `json:"errors"`


    // Statistics - Statistics related to the sync job.
    Statistics Knowledgesyncjobstatistics `json:"statistics"`

}

// String returns a JSON representation of the model
func (o *Knowledgesyncjobreport) String() string {
     o.Errors = []Errorbody{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesyncjobreport) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesyncjobreport

    if KnowledgesyncjobreportMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesyncjobreportMarshalled = true

    return json.Marshal(&struct {
        
        Errors []Errorbody `json:"errors"`
        
        Statistics Knowledgesyncjobstatistics `json:"statistics"`
        *Alias
    }{

        
        Errors: []Errorbody{{}},
        


        

        Alias: (*Alias)(u),
    })
}

