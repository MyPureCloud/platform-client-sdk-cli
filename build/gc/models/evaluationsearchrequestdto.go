package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationsearchrequestdtoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationsearchrequestdtoDud struct { 
    


    


    


    


    


    


    

}

// Evaluationsearchrequestdto
type Evaluationsearchrequestdto struct { 
    // Query
    Query []Evaluationsearchcriteriadto `json:"query"`


    // Aggregations - Aggregations to compute on the search results
    Aggregations []Evaluationsearchaggregationdto `json:"aggregations"`


    // PageSize - The number of results per page. For aggregation requests, must be omitted or 0.
    PageSize int `json:"pageSize"`


    // PageNumber - The page of resources you want to retrieve
    PageNumber int `json:"pageNumber"`


    // SortOrder - The sort order for results. Include with sortBy.
    SortOrder string `json:"sortOrder"`


    // SortBy - The field in the resource that you want to sort the results by. Include with sortOrder.
    SortBy string `json:"sortBy"`


    // SystemSubmitted - Filter for automated evaluations submitted by virtual supervisor. Defaults to false.
    SystemSubmitted bool `json:"systemSubmitted"`

}

// String returns a JSON representation of the model
func (o *Evaluationsearchrequestdto) String() string {
     o.Query = []Evaluationsearchcriteriadto{{}} 
     o.Aggregations = []Evaluationsearchaggregationdto{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationsearchrequestdto) MarshalJSON() ([]byte, error) {
    type Alias Evaluationsearchrequestdto

    if EvaluationsearchrequestdtoMarshalled {
        return []byte("{}"), nil
    }
    EvaluationsearchrequestdtoMarshalled = true

    return json.Marshal(&struct {
        
        Query []Evaluationsearchcriteriadto `json:"query"`
        
        Aggregations []Evaluationsearchaggregationdto `json:"aggregations"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        SystemSubmitted bool `json:"systemSubmitted"`
        *Alias
    }{

        
        Query: []Evaluationsearchcriteriadto{{}},
        


        
        Aggregations: []Evaluationsearchaggregationdto{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

