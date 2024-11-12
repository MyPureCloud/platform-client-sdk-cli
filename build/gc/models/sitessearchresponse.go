package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SitessearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SitessearchresponseDud struct { 
    


    


    


    


    


    


    


    


    

}

// Sitessearchresponse
type Sitessearchresponse struct { 
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
    Results []Site `json:"results"`

}

// String returns a JSON representation of the model
func (o *Sitessearchresponse) String() string {
    
    
    
    
    
    
    
     o.Types = []string{""} 
     o.Results = []Site{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sitessearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Sitessearchresponse

    if SitessearchresponseMarshalled {
        return []byte("{}"), nil
    }
    SitessearchresponseMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        PreviousPage string `json:"previousPage"`
        
        CurrentPage string `json:"currentPage"`
        
        NextPage string `json:"nextPage"`
        
        Types []string `json:"types"`
        
        Results []Site `json:"results"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Types: []string{""},
        


        
        Results: []Site{{}},
        

        Alias: (*Alias)(u),
    })
}

