package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JsonnodesearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JsonnodesearchresponseDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Jsonnodesearchresponse
type Jsonnodesearchresponse struct { 
    // Total - The total number of results found
    Total int `json:"total"`


    // PageCount - The total number of pages
    PageCount int `json:"pageCount"`


    // PageSize - The current page size
    PageSize int `json:"pageSize"`


    // PageNumber - The current page number
    PageNumber int `json:"pageNumber"`


    // PreviousPage - Q64 value for the previous page of results
    PreviousPage string `json:"previousPage"`


    // CurrentPage - Q64 value for the current page of results
    CurrentPage string `json:"currentPage"`


    // NextPage - Q64 value for the next page of results
    NextPage string `json:"nextPage"`


    // Types - Resource types the search was performed against
    Types []string `json:"types"`


    // Results - Search results
    Results interface{} `json:"results"`


    // Aggregations
    Aggregations interface{} `json:"aggregations"`

}

// String returns a JSON representation of the model
func (o *Jsonnodesearchresponse) String() string {
    
    
    
    
    
    
    
     o.Types = []string{""} 
     o.Results = Interface{} 
     o.Aggregations = Interface{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Jsonnodesearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Jsonnodesearchresponse

    if JsonnodesearchresponseMarshalled {
        return []byte("{}"), nil
    }
    JsonnodesearchresponseMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        PreviousPage string `json:"previousPage"`
        
        CurrentPage string `json:"currentPage"`
        
        NextPage string `json:"nextPage"`
        
        Types []string `json:"types"`
        
        Results interface{} `json:"results"`
        
        Aggregations interface{} `json:"aggregations"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Types: []string{""},
        


        
        Results: Interface{},
        


        
        Aggregations: Interface{},
        

        Alias: (*Alias)(u),
    })
}

