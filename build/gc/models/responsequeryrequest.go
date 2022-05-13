package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponsequeryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponsequeryrequestDud struct { 
    


    


    

}

// Responsequeryrequest - Used to query for responses
type Responsequeryrequest struct { 
    // QueryPhrase - Query phrase to search response text and name. If not set will match all.
    QueryPhrase string `json:"queryPhrase"`


    // PageSize - The maximum number of hits to return. Default: 25, Maximum: 500.
    PageSize int `json:"pageSize"`


    // Filters - Filter the query results.
    Filters []Responsefilter `json:"filters"`

}

// String returns a JSON representation of the model
func (o *Responsequeryrequest) String() string {
    
    
     o.Filters = []Responsefilter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responsequeryrequest) MarshalJSON() ([]byte, error) {
    type Alias Responsequeryrequest

    if ResponsequeryrequestMarshalled {
        return []byte("{}"), nil
    }
    ResponsequeryrequestMarshalled = true

    return json.Marshal(&struct {
        
        QueryPhrase string `json:"queryPhrase"`
        
        PageSize int `json:"pageSize"`
        
        Filters []Responsefilter `json:"filters"`
        *Alias
    }{

        


        


        
        Filters: []Responsefilter{{}},
        

        Alias: (*Alias)(u),
    })
}

