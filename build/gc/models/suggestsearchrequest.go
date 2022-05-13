package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestsearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestsearchrequestDud struct { 
    


    


    

}

// Suggestsearchrequest
type Suggestsearchrequest struct { 
    // Expand - Provides more details about a specified resource
    Expand []string `json:"expand"`


    // Types - Resource domain type to search
    Types []string `json:"types"`


    // Query - Suggest query
    Query []Suggestsearchcriteria `json:"query"`

}

// String returns a JSON representation of the model
func (o *Suggestsearchrequest) String() string {
     o.Expand = []string{""} 
     o.Types = []string{""} 
     o.Query = []Suggestsearchcriteria{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestsearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Suggestsearchrequest

    if SuggestsearchrequestMarshalled {
        return []byte("{}"), nil
    }
    SuggestsearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        Expand []string `json:"expand"`
        
        Types []string `json:"types"`
        
        Query []Suggestsearchcriteria `json:"query"`
        *Alias
    }{

        
        Expand: []string{""},
        


        
        Types: []string{""},
        


        
        Query: []Suggestsearchcriteria{{}},
        

        Alias: (*Alias)(u),
    })
}

