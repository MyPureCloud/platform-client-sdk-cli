package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialerauditrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialerauditrequestDud struct { 
    


    


    


    

}

// Dialerauditrequest
type Dialerauditrequest struct { 
    // QueryPhrase - The word or words to search for.
    QueryPhrase string `json:"queryPhrase"`


    // QueryFields - The fields in which to search for the queryPhrase.
    QueryFields []string `json:"queryFields"`


    // Facets - The fields to facet on.
    Facets []Auditfacet `json:"facets"`


    // Filters - The fields to filter on.
    Filters []Auditfilter `json:"filters"`

}

// String returns a JSON representation of the model
func (o *Dialerauditrequest) String() string {
    
    
    
    
    
    
     o.QueryFields = []string{""} 
    
    
    
     o.Facets = []Auditfacet{{}} 
    
    
    
     o.Filters = []Auditfilter{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialerauditrequest) MarshalJSON() ([]byte, error) {
    type Alias Dialerauditrequest

    if DialerauditrequestMarshalled {
        return []byte("{}"), nil
    }
    DialerauditrequestMarshalled = true

    return json.Marshal(&struct { 
        QueryPhrase string `json:"queryPhrase"`
        
        QueryFields []string `json:"queryFields"`
        
        Facets []Auditfacet `json:"facets"`
        
        Filters []Auditfilter `json:"filters"`
        
        *Alias
    }{
        

        

        

        
        QueryFields: []string{""},
        

        

        
        Facets: []Auditfacet{{}},
        

        

        
        Filters: []Auditfilter{{}},
        

        
        Alias: (*Alias)(u),
    })
}

