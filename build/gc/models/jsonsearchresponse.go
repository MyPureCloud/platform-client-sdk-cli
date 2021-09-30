package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JsonsearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JsonsearchresponseDud struct { 
    


    


    


    


    


    


    

}

// Jsonsearchresponse
type Jsonsearchresponse struct { 
    // Total - The total number of results found
    Total int `json:"total"`


    // PageCount - The total number of pages
    PageCount int `json:"pageCount"`


    // PageSize - The current page size
    PageSize int `json:"pageSize"`


    // PageNumber - The current page number
    PageNumber int `json:"pageNumber"`


    // Types - Resource types the search was performed against
    Types []string `json:"types"`


    // Results - Search results
    Results Arraynode `json:"results"`


    // Aggregations
    Aggregations Arraynode `json:"aggregations"`

}

// String returns a JSON representation of the model
func (o *Jsonsearchresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Types = []string{""} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Jsonsearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Jsonsearchresponse

    if JsonsearchresponseMarshalled {
        return []byte("{}"), nil
    }
    JsonsearchresponseMarshalled = true

    return json.Marshal(&struct { 
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Types []string `json:"types"`
        
        Results Arraynode `json:"results"`
        
        Aggregations Arraynode `json:"aggregations"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Types: []string{""},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

