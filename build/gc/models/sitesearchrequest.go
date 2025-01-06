package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SitesearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SitesearchrequestDud struct { 
    


    


    


    


    


    


    

}

// Sitesearchrequest
type Sitesearchrequest struct { 
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


    // Expand - Provides more details about a specified resource
    Expand []string `json:"expand"`


    // Query
    Query []Sitesearchcriteria `json:"query"`

}

// String returns a JSON representation of the model
func (o *Sitesearchrequest) String() string {
    
    
    
    
     o.Sort = []Searchsort{{}} 
     o.Expand = []string{""} 
     o.Query = []Sitesearchcriteria{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sitesearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Sitesearchrequest

    if SitesearchrequestMarshalled {
        return []byte("{}"), nil
    }
    SitesearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Sort []Searchsort `json:"sort"`
        
        Expand []string `json:"expand"`
        
        Query []Sitesearchcriteria `json:"query"`
        *Alias
    }{

        


        


        


        


        
        Sort: []Searchsort{{}},
        


        
        Expand: []string{""},
        


        
        Query: []Sitesearchcriteria{{}},
        

        Alias: (*Alias)(u),
    })
}

