package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptsearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptsearchrequestDud struct { 
    


    


    


    


    


    


    


    

}

// Transcriptsearchrequest
type Transcriptsearchrequest struct { 
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


    // ReturnFields
    ReturnFields []string `json:"returnFields"`


    // Types - Resource domain type to search
    Types []string `json:"types"`


    // Query - The search criteria
    Query []Transcriptsearchcriteria `json:"query"`

}

// String returns a JSON representation of the model
func (o *Transcriptsearchrequest) String() string {
    
    
    
    
     o.Sort = []Searchsort{{}} 
     o.ReturnFields = []string{""} 
     o.Types = []string{""} 
     o.Query = []Transcriptsearchcriteria{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptsearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Transcriptsearchrequest

    if TranscriptsearchrequestMarshalled {
        return []byte("{}"), nil
    }
    TranscriptsearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Sort []Searchsort `json:"sort"`
        
        ReturnFields []string `json:"returnFields"`
        
        Types []string `json:"types"`
        
        Query []Transcriptsearchcriteria `json:"query"`
        *Alias
    }{

        


        


        


        


        
        Sort: []Searchsort{{}},
        


        
        ReturnFields: []string{""},
        


        
        Types: []string{""},
        


        
        Query: []Transcriptsearchcriteria{{}},
        

        Alias: (*Alias)(u),
    })
}

