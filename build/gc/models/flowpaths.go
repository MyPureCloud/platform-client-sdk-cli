package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowpathsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowpathsDud struct { 
    


    


    


    

}

// Flowpaths
type Flowpaths struct { 
    // Category - Category (use case) of the paths within a given domain.
    Category string `json:"category"`


    // DateStart - Start date of the date range included in the flow paths data. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // DateEnd - End date of the date range included in the flow paths data. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnd time.Time `json:"dateEnd"`


    // Elements - Unique element identifiers and their corresponding elements in the trie data structure representing the paths.
    Elements map[string]Flowpathselement `json:"elements"`

}

// String returns a JSON representation of the model
func (o *Flowpaths) String() string {
    
    
    
     o.Elements = map[string]Flowpathselement{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowpaths) MarshalJSON() ([]byte, error) {
    type Alias Flowpaths

    if FlowpathsMarshalled {
        return []byte("{}"), nil
    }
    FlowpathsMarshalled = true

    return json.Marshal(&struct {
        
        Category string `json:"category"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        Elements map[string]Flowpathselement `json:"elements"`
        *Alias
    }{

        


        


        


        
        Elements: map[string]Flowpathselement{"": {}},
        

        Alias: (*Alias)(u),
    })
}

