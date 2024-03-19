package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryuserstaffinggrouplistrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryuserstaffinggrouplistrequestDud struct { 
    

}

// Queryuserstaffinggrouplistrequest
type Queryuserstaffinggrouplistrequest struct { 
    // UserIds - The set of user IDs to fetch associated staffing groups
    UserIds []string `json:"userIds"`

}

// String returns a JSON representation of the model
func (o *Queryuserstaffinggrouplistrequest) String() string {
     o.UserIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryuserstaffinggrouplistrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryuserstaffinggrouplistrequest

    if QueryuserstaffinggrouplistrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryuserstaffinggrouplistrequestMarshalled = true

    return json.Marshal(&struct {
        
        UserIds []string `json:"userIds"`
        *Alias
    }{

        
        UserIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

