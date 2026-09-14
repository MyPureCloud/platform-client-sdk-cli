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
    // Name - Specify an attribute for sorting. Use 'customFields' to sort by a custom field, in which case the customField property is required.
    Name string `json:"name"`


    // Ascending - Sort Ascending
    Ascending bool `json:"ascending"`


    // CustomField - The key of the custom field to sort by. Required when name is 'customFields' and must not be set otherwise.
    CustomField string `json:"customField"`

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
        
        CustomField string `json:"customField"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

