package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryrequestpredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryrequestpredicateDud struct { 
    


    

}

// Queryrequestpredicate
type Queryrequestpredicate struct { 
    // Dimension - The dimension to be filtered
    Dimension string `json:"dimension"`


    // Value - The value to filter by
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Queryrequestpredicate) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryrequestpredicate) MarshalJSON() ([]byte, error) {
    type Alias Queryrequestpredicate

    if QueryrequestpredicateMarshalled {
        return []byte("{}"), nil
    }
    QueryrequestpredicateMarshalled = true

    return json.Marshal(&struct { 
        Dimension string `json:"dimension"`
        
        Value string `json:"value"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

