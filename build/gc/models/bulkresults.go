package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkresultsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkresultsDud struct { 
    


    

}

// Bulkresults
type Bulkresults struct { 
    // Id - Id of source intent
    Id string `json:"id"`


    // Success - Result of operation
    Success bool `json:"success"`

}

// String returns a JSON representation of the model
func (o *Bulkresults) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkresults) MarshalJSON() ([]byte, error) {
    type Alias Bulkresults

    if BulkresultsMarshalled {
        return []byte("{}"), nil
    }
    BulkresultsMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Success bool `json:"success"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

