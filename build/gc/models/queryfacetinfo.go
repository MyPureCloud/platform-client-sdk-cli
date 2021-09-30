package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryfacetinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryfacetinfoDud struct { 
    


    

}

// Queryfacetinfo
type Queryfacetinfo struct { 
    // Attributes
    Attributes []Facetkeyattribute `json:"attributes"`


    // Facets
    Facets []Facetentry `json:"facets"`

}

// String returns a JSON representation of the model
func (o *Queryfacetinfo) String() string {
    
    
     o.Attributes = []Facetkeyattribute{{}} 
    
    
    
     o.Facets = []Facetentry{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryfacetinfo) MarshalJSON() ([]byte, error) {
    type Alias Queryfacetinfo

    if QueryfacetinfoMarshalled {
        return []byte("{}"), nil
    }
    QueryfacetinfoMarshalled = true

    return json.Marshal(&struct { 
        Attributes []Facetkeyattribute `json:"attributes"`
        
        Facets []Facetentry `json:"facets"`
        
        *Alias
    }{
        

        
        Attributes: []Facetkeyattribute{{}},
        

        

        
        Facets: []Facetentry{{}},
        

        
        Alias: (*Alias)(u),
    })
}

