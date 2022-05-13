package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentqueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentqueryrequestDud struct { 
    


    


    


    


    


    


    


    

}

// Contentqueryrequest
type Contentqueryrequest struct { 
    // QueryPhrase
    QueryPhrase string `json:"queryPhrase"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // PageSize
    PageSize int `json:"pageSize"`


    // FacetNameRequests
    FacetNameRequests []string `json:"facetNameRequests"`


    // Sort
    Sort []Contentsortitem `json:"sort"`


    // Filters
    Filters []Contentfacetfilteritem `json:"filters"`


    // AttributeFilters
    AttributeFilters []Contentattributefilteritem `json:"attributeFilters"`


    // IncludeShares
    IncludeShares bool `json:"includeShares"`

}

// String returns a JSON representation of the model
func (o *Contentqueryrequest) String() string {
    
    
    
     o.FacetNameRequests = []string{""} 
     o.Sort = []Contentsortitem{{}} 
     o.Filters = []Contentfacetfilteritem{{}} 
     o.AttributeFilters = []Contentattributefilteritem{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentqueryrequest) MarshalJSON() ([]byte, error) {
    type Alias Contentqueryrequest

    if ContentqueryrequestMarshalled {
        return []byte("{}"), nil
    }
    ContentqueryrequestMarshalled = true

    return json.Marshal(&struct {
        
        QueryPhrase string `json:"queryPhrase"`
        
        PageNumber int `json:"pageNumber"`
        
        PageSize int `json:"pageSize"`
        
        FacetNameRequests []string `json:"facetNameRequests"`
        
        Sort []Contentsortitem `json:"sort"`
        
        Filters []Contentfacetfilteritem `json:"filters"`
        
        AttributeFilters []Contentattributefilteritem `json:"attributeFilters"`
        
        IncludeShares bool `json:"includeShares"`
        *Alias
    }{

        


        


        


        
        FacetNameRequests: []string{""},
        


        
        Sort: []Contentsortitem{{}},
        


        
        Filters: []Contentfacetfilteritem{{}},
        


        
        AttributeFilters: []Contentattributefilteritem{{}},
        


        

        Alias: (*Alias)(u),
    })
}

