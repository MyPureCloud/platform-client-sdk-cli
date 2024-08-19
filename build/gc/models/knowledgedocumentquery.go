package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentqueryDud struct { 
    


    


    


    


    


    


    


    

}

// Knowledgedocumentquery
type Knowledgedocumentquery struct { 
    // PageSize - Page size of the returned results.
    PageSize int `json:"pageSize"`


    // PageNumber - Page number of the returned results.
    PageNumber int `json:"pageNumber"`


    // IncludeDraftDocuments - Indicates whether the results would also include draft documents.
    IncludeDraftDocuments bool `json:"includeDraftDocuments"`


    // Interval - Retrieves the documents created/modified/published in specified date and time range.
    Interval Documentqueryinterval `json:"interval"`


    // Filter - Filter for the document query.
    Filter Documentquery `json:"filter"`


    // IncludeVariations - Indicates which document variations to include in returned documents. All: all variations regardless of the filter expression; AllMatching: all variations that match the filter expression; SingleMostRelevant: single variation that matches the filter expression and has the highest priority. The default is All.
    IncludeVariations string `json:"includeVariations"`


    // SortOrder - The sort order for results.
    SortOrder string `json:"sortOrder"`


    // SortBy - The field in the documents that you want to sort the results by.
    SortBy string `json:"sortBy"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentquery) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentquery) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentquery

    if KnowledgedocumentqueryMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentqueryMarshalled = true

    return json.Marshal(&struct {
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        IncludeDraftDocuments bool `json:"includeDraftDocuments"`
        
        Interval Documentqueryinterval `json:"interval"`
        
        Filter Documentquery `json:"filter"`
        
        IncludeVariations string `json:"includeVariations"`
        
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

