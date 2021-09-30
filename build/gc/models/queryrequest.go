package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryrequestDud struct { 
    


    


    


    


    


    


    


    

}

// Queryrequest
type Queryrequest struct { 
    // QueryPhrase
    QueryPhrase string `json:"queryPhrase"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // PageSize
    PageSize int `json:"pageSize"`


    // FacetNameRequests
    FacetNameRequests []string `json:"facetNameRequests"`


    // Sort
    Sort []Sortitem `json:"sort"`


    // Filters
    Filters []Contentfilteritem `json:"filters"`


    // AttributeFilters
    AttributeFilters []Attributefilteritem `json:"attributeFilters"`


    // IncludeShares
    IncludeShares bool `json:"includeShares"`

}

// String returns a JSON representation of the model
func (o *Queryrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.FacetNameRequests = []string{""} 
    
    
    
     o.Sort = []Sortitem{{}} 
    
    
    
     o.Filters = []Contentfilteritem{{}} 
    
    
    
     o.AttributeFilters = []Attributefilteritem{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryrequest

    if QueryrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryrequestMarshalled = true

    return json.Marshal(&struct { 
        QueryPhrase string `json:"queryPhrase"`
        
        PageNumber int `json:"pageNumber"`
        
        PageSize int `json:"pageSize"`
        
        FacetNameRequests []string `json:"facetNameRequests"`
        
        Sort []Sortitem `json:"sort"`
        
        Filters []Contentfilteritem `json:"filters"`
        
        AttributeFilters []Attributefilteritem `json:"attributeFilters"`
        
        IncludeShares bool `json:"includeShares"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        FacetNameRequests: []string{""},
        

        

        
        Sort: []Sortitem{{}},
        

        

        
        Filters: []Contentfilteritem{{}},
        

        

        
        AttributeFilters: []Attributefilteritem{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

