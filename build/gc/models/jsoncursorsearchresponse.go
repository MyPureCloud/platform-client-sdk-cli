package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JsoncursorsearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JsoncursorsearchresponseDud struct { 
    


    


    


    

}

// Jsoncursorsearchresponse
type Jsoncursorsearchresponse struct { 
    // Types - Resource types the search was performed against
    Types []string `json:"types"`


    // Results - Search results
    Results interface{} `json:"results"`


    // Aggregations
    Aggregations interface{} `json:"aggregations"`


    // Cursor - The page cursor
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Jsoncursorsearchresponse) String() string {
     o.Types = []string{""} 
     o.Results = Interface{} 
     o.Aggregations = Interface{} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Jsoncursorsearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Jsoncursorsearchresponse

    if JsoncursorsearchresponseMarshalled {
        return []byte("{}"), nil
    }
    JsoncursorsearchresponseMarshalled = true

    return json.Marshal(&struct {
        
        Types []string `json:"types"`
        
        Results interface{} `json:"results"`
        
        Aggregations interface{} `json:"aggregations"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Types: []string{""},
        


        
        Results: Interface{},
        


        
        Aggregations: Interface{},
        


        

        Alias: (*Alias)(u),
    })
}

