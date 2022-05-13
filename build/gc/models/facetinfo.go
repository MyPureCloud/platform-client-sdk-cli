package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacetinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacetinfoDud struct { 
    


    

}

// Facetinfo
type Facetinfo struct { 
    // Name - The name of the field that was faceted on.
    Name string `json:"name"`


    // Entries - The entries resulting from this facet.
    Entries []Entry `json:"entries"`

}

// String returns a JSON representation of the model
func (o *Facetinfo) String() string {
    
     o.Entries = []Entry{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facetinfo) MarshalJSON() ([]byte, error) {
    type Alias Facetinfo

    if FacetinfoMarshalled {
        return []byte("{}"), nil
    }
    FacetinfoMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Entries []Entry `json:"entries"`
        *Alias
    }{

        


        
        Entries: []Entry{{}},
        

        Alias: (*Alias)(u),
    })
}

