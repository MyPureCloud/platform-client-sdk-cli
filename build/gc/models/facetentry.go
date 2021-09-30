package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacetentryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacetentryDud struct { 
    


    


    


    


    


    


    


    

}

// Facetentry
type Facetentry struct { 
    // Attribute
    Attribute Termattribute `json:"attribute"`


    // Statistics
    Statistics Facetstatistics `json:"statistics"`


    // Other
    Other int `json:"other"`


    // Total
    Total int `json:"total"`


    // Missing
    Missing int `json:"missing"`


    // TermCount
    TermCount int `json:"termCount"`


    // TermType
    TermType string `json:"termType"`


    // Terms
    Terms []Facetterm `json:"terms"`

}

// String returns a JSON representation of the model
func (o *Facetentry) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Terms = []Facetterm{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facetentry) MarshalJSON() ([]byte, error) {
    type Alias Facetentry

    if FacetentryMarshalled {
        return []byte("{}"), nil
    }
    FacetentryMarshalled = true

    return json.Marshal(&struct { 
        Attribute Termattribute `json:"attribute"`
        
        Statistics Facetstatistics `json:"statistics"`
        
        Other int `json:"other"`
        
        Total int `json:"total"`
        
        Missing int `json:"missing"`
        
        TermCount int `json:"termCount"`
        
        TermType string `json:"termType"`
        
        Terms []Facetterm `json:"terms"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Terms: []Facetterm{{}},
        

        
        Alias: (*Alias)(u),
    })
}

