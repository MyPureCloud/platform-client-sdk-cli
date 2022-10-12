package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestdocumentsuggestionrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestdocumentsuggestionrequestDud struct { 
    


    


    

}

// Knowledgeguestdocumentsuggestionrequest
type Knowledgeguestdocumentsuggestionrequest struct { 
    // Query - Query to get autocomplete suggestions for the matching knowledge documents.
    Query string `json:"query"`


    // PageSize - Page size of the returned results.
    PageSize int `json:"pageSize"`


    // IncludeDraftDocuments - Indicates whether the suggestion results would also include draft documents.
    IncludeDraftDocuments bool `json:"includeDraftDocuments"`

}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentsuggestionrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestdocumentsuggestionrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestdocumentsuggestionrequest

    if KnowledgeguestdocumentsuggestionrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestdocumentsuggestionrequestMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        PageSize int `json:"pageSize"`
        
        IncludeDraftDocuments bool `json:"includeDraftDocuments"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

