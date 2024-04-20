package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ApiusagequeryresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ApiusagequeryresultDud struct { 
    


    


    

}

// Apiusagequeryresult
type Apiusagequeryresult struct { 
    // Results - Query results
    Results []Apiusagerow `json:"results"`


    // QueryStatus - Query status
    QueryStatus string `json:"queryStatus"`


    // Cursors - Cursor tokens to be used for navigating paginated results
    Cursors Cursors `json:"cursors"`

}

// String returns a JSON representation of the model
func (o *Apiusagequeryresult) String() string {
     o.Results = []Apiusagerow{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Apiusagequeryresult) MarshalJSON() ([]byte, error) {
    type Alias Apiusagequeryresult

    if ApiusagequeryresultMarshalled {
        return []byte("{}"), nil
    }
    ApiusagequeryresultMarshalled = true

    return json.Marshal(&struct {
        
        Results []Apiusagerow `json:"results"`
        
        QueryStatus string `json:"queryStatus"`
        
        Cursors Cursors `json:"cursors"`
        *Alias
    }{

        
        Results: []Apiusagerow{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

