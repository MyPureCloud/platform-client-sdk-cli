package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacettermMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacettermDud struct { 
    


    


    


    


    


    

}

// Facetterm
type Facetterm struct { 
    // Term
    Term string `json:"term"`


    // Key
    Key int `json:"key"`


    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // Count
    Count int `json:"count"`


    // Time - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Time time.Time `json:"time"`

}

// String returns a JSON representation of the model
func (o *Facetterm) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facetterm) MarshalJSON() ([]byte, error) {
    type Alias Facetterm

    if FacettermMarshalled {
        return []byte("{}"), nil
    }
    FacettermMarshalled = true

    return json.Marshal(&struct {
        
        Term string `json:"term"`
        
        Key int `json:"key"`
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Count int `json:"count"`
        
        Time time.Time `json:"time"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

