package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentsuggestionrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentsuggestionrequestDud struct { 
    


    


    


    


    

}

// Knowledgedocumentsuggestionrequest
type Knowledgedocumentsuggestionrequest struct { 
    // Query - Query to get autocomplete suggestions for the matching knowledge documents.
    Query string `json:"query"`


    // PageSize - Page size of the returned results.
    PageSize int `json:"pageSize"`


    // IncludeDraftDocuments - Indicates whether the suggestion results would also include draft documents.
    IncludeDraftDocuments bool `json:"includeDraftDocuments"`


    // Interval - Retrieves the documents created/modified/published in specified date and time range.
    Interval Documentqueryinterval `json:"interval"`


    // Filter - Filter for the document suggestions.
    Filter Documentquery `json:"filter"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsuggestionrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentsuggestionrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentsuggestionrequest

    if KnowledgedocumentsuggestionrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentsuggestionrequestMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        PageSize int `json:"pageSize"`
        
        IncludeDraftDocuments bool `json:"includeDraftDocuments"`
        
        Interval Documentqueryinterval `json:"interval"`
        
        Filter Documentquery `json:"filter"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

