package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseplanqueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseplanqueryrequestDud struct { 
    


    


    


    


    


    


    

}

// Caseplanqueryrequest
type Caseplanqueryrequest struct { 
    // Name - Filter by Caseplan name (case-insensitive, partial match). Omitting name returns all Caseplans (subject to pagination).
    Name string `json:"name"`


    // NameSearchType - Type of name search to perform. Default is BEGINS_WITH.
    NameSearchType string `json:"nameSearchType"`


    // DivisionIds - Divisions to filter by. Accepts a list of UUIDs and/or '*'.
    DivisionIds []string `json:"divisionIds"`


    // Filters - List of filter objects to be used in the search. Valid filter names are: 'id', 'name', 'divisionId'. Multiple filters are combined with AND logic.
    Filters []Caseplanfilter `json:"filters"`


    // Attributes - List of entity attributes to be retrieved in the result.
    Attributes []string `json:"attributes"`


    // PageSize - Number of results per page. Maximum is 200. Default is 25.
    PageSize int `json:"pageSize"`


    // After - Cursor for pagination. Use the \"after\" value from the previous response.
    After string `json:"after"`

}

// String returns a JSON representation of the model
func (o *Caseplanqueryrequest) String() string {
    
    
     o.DivisionIds = []string{""} 
     o.Filters = []Caseplanfilter{{}} 
     o.Attributes = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseplanqueryrequest) MarshalJSON() ([]byte, error) {
    type Alias Caseplanqueryrequest

    if CaseplanqueryrequestMarshalled {
        return []byte("{}"), nil
    }
    CaseplanqueryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        NameSearchType string `json:"nameSearchType"`
        
        DivisionIds []string `json:"divisionIds"`
        
        Filters []Caseplanfilter `json:"filters"`
        
        Attributes []string `json:"attributes"`
        
        PageSize int `json:"pageSize"`
        
        After string `json:"after"`
        *Alias
    }{

        


        


        
        DivisionIds: []string{""},
        


        
        Filters: []Caseplanfilter{{}},
        


        
        Attributes: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

