package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgeidnamepairMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgeidnamepairDud struct { 
    Id string `json:"id"`


    

}

// Edgeidnamepair
type Edgeidnamepair struct { 
    


    // Name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Edgeidnamepair) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgeidnamepair) MarshalJSON() ([]byte, error) {
    type Alias Edgeidnamepair

    if EdgeidnamepairMarshalled {
        return []byte("{}"), nil
    }
    EdgeidnamepairMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

