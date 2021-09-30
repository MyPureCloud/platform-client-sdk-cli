package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SearchsortMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SearchsortDud struct { 
    


    

}

// Searchsort
type Searchsort struct { 
    // SortOrder - The sort order for results
    SortOrder string `json:"sortOrder"`


    // SortBy - The field in the resource that you want to sort the results by
    SortBy string `json:"sortBy"`

}

// String returns a JSON representation of the model
func (o *Searchsort) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Searchsort) MarshalJSON() ([]byte, error) {
    type Alias Searchsort

    if SearchsortMarshalled {
        return []byte("{}"), nil
    }
    SearchsortMarshalled = true

    return json.Marshal(&struct { 
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

