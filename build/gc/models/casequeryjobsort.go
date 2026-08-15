package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasequeryjobsortMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasequeryjobsortDud struct { 
    


    

}

// Casequeryjobsort
type Casequeryjobsort struct { 
    // Name - The attribute to sort by. The default is dateDue.
    Name string `json:"name"`


    // Ascending - Whether to sort in ascending order. The default is false.
    Ascending bool `json:"ascending"`

}

// String returns a JSON representation of the model
func (o *Casequeryjobsort) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casequeryjobsort) MarshalJSON() ([]byte, error) {
    type Alias Casequeryjobsort

    if CasequeryjobsortMarshalled {
        return []byte("{}"), nil
    }
    CasequeryjobsortMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Ascending bool `json:"ascending"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

