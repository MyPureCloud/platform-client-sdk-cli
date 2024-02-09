package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeparseimportresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeparseimportresultDud struct { 
    


    


    

}

// Knowledgeparseimportresult
type Knowledgeparseimportresult struct { 
    // Success - Number of imported articles.
    Success int `json:"success"`


    // Failure - Number of articles failed to import.
    Failure int `json:"failure"`


    // Errors - Error information about the failures.
    Errors []Errorbody `json:"errors"`

}

// String returns a JSON representation of the model
func (o *Knowledgeparseimportresult) String() string {
    
    
     o.Errors = []Errorbody{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeparseimportresult) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeparseimportresult

    if KnowledgeparseimportresultMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeparseimportresultMarshalled = true

    return json.Marshal(&struct {
        
        Success int `json:"success"`
        
        Failure int `json:"failure"`
        
        Errors []Errorbody `json:"errors"`
        *Alias
    }{

        


        


        
        Errors: []Errorbody{{}},
        

        Alias: (*Alias)(u),
    })
}

