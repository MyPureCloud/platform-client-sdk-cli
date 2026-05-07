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
    // Name - Filter by caseplan name (case-insensitive, partial match). Omitting name returns all caseplans (subject to pagination).
    Name string `json:"name"`


    // PageSize - Number of results per page. Maximum is 200. Default is 25.
    PageSize int `json:"pageSize"`


    // After - Cursor for pagination. Use the \"after\" value from the previous response.
    After string `json:"after"`


    // DivisionIds - Divisions to filter by. Accepts a list of UUIDs and/or '*'.
    DivisionIds []string `json:"divisionIds"`

}

// String returns a JSON representation of the model
func (o *Caseplanqueryrequest) String() string {
    
    
    
     o.DivisionIds = []string{""} 

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
        
        PageSize int `json:"pageSize"`
        
        After string `json:"after"`
        
        DivisionIds []string `json:"divisionIds"`
        *Alias
    }{

        


        


        


        
        DivisionIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

