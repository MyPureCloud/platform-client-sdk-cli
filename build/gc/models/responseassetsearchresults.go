package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponseassetsearchresultsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponseassetsearchresultsDud struct { 
    


    


    


    


    

}

// Responseassetsearchresults
type Responseassetsearchresults struct { 
    // Total - The total number of results found
    Total int `json:"total"`


    // PageCount - The total number of pages
    PageCount int `json:"pageCount"`


    // PageSize - The current page size
    PageSize int `json:"pageSize"`


    // PageNumber - The current page number
    PageNumber int `json:"pageNumber"`


    // Results - Search results
    Results []Responseasset `json:"results"`

}

// String returns a JSON representation of the model
func (o *Responseassetsearchresults) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Results = []Responseasset{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responseassetsearchresults) MarshalJSON() ([]byte, error) {
    type Alias Responseassetsearchresults

    if ResponseassetsearchresultsMarshalled {
        return []byte("{}"), nil
    }
    ResponseassetsearchresultsMarshalled = true

    return json.Marshal(&struct { 
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Results []Responseasset `json:"results"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Results: []Responseasset{{}},
        

        
        Alias: (*Alias)(u),
    })
}

