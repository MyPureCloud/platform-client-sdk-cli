package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemqueryjobsortMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemqueryjobsortDud struct { 
    


    

}

// Workitemqueryjobsort
type Workitemqueryjobsort struct { 
    // Name - Specify an attribute for sorting.
    Name string `json:"name"`


    // Ascending - Sort Ascending
    Ascending bool `json:"ascending"`

}

// String returns a JSON representation of the model
func (o *Workitemqueryjobsort) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemqueryjobsort) MarshalJSON() ([]byte, error) {
    type Alias Workitemqueryjobsort

    if WorkitemqueryjobsortMarshalled {
        return []byte("{}"), nil
    }
    WorkitemqueryjobsortMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Ascending bool `json:"ascending"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

