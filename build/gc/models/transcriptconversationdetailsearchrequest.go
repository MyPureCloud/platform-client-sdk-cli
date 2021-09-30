package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptconversationdetailsearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptconversationdetailsearchrequestDud struct { 
    


    


    


    


    


    


    

}

// Transcriptconversationdetailsearchrequest
type Transcriptconversationdetailsearchrequest struct { 
    // SortOrder - The sort order for results
    SortOrder string `json:"sortOrder"`


    // SortBy - The field in the resource that you want to sort the results by
    SortBy string `json:"sortBy"`


    // PageSize - The number of results per page
    PageSize int `json:"pageSize"`


    // PageNumber - The page of resources you want to retrieve
    PageNumber int `json:"pageNumber"`


    // Sort - Multi-value sort order, list of multiple sort values
    Sort []Searchsort `json:"sort"`


    // Types - Resource domain type to search
    Types []string `json:"types"`


    // Query - The search criteria
    Query []Transcriptconversationdetailsearchcriteria `json:"query"`

}

// String returns a JSON representation of the model
func (o *Transcriptconversationdetailsearchrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Sort = []Searchsort{{}} 
    
    
    
     o.Types = []string{""} 
    
    
    
     o.Query = []Transcriptconversationdetailsearchcriteria{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptconversationdetailsearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Transcriptconversationdetailsearchrequest

    if TranscriptconversationdetailsearchrequestMarshalled {
        return []byte("{}"), nil
    }
    TranscriptconversationdetailsearchrequestMarshalled = true

    return json.Marshal(&struct { 
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Sort []Searchsort `json:"sort"`
        
        Types []string `json:"types"`
        
        Query []Transcriptconversationdetailsearchcriteria `json:"query"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Sort: []Searchsort{{}},
        

        

        
        Types: []string{""},
        

        

        
        Query: []Transcriptconversationdetailsearchcriteria{{}},
        

        
        Alias: (*Alias)(u),
    })
}

