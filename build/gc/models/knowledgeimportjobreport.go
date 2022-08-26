package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeimportjobreportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeimportjobreportDud struct { 
    


    

}

// Knowledgeimportjobreport
type Knowledgeimportjobreport struct { 
    // Errors - List of errors occurred during processing import.
    Errors []Knowledgeimportjoberror `json:"errors"`


    // Statistics - Statistics related to the import job.
    Statistics Knowledgeimportjobstatistics `json:"statistics"`

}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobreport) String() string {
     o.Errors = []Knowledgeimportjoberror{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeimportjobreport) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeimportjobreport

    if KnowledgeimportjobreportMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeimportjobreportMarshalled = true

    return json.Marshal(&struct {
        
        Errors []Knowledgeimportjoberror `json:"errors"`
        
        Statistics Knowledgeimportjobstatistics `json:"statistics"`
        *Alias
    }{

        
        Errors: []Knowledgeimportjoberror{{}},
        


        

        Alias: (*Alias)(u),
    })
}

