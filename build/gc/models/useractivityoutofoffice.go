package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivityoutofofficeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivityoutofofficeDud struct { 
    


    

}

// Useractivityoutofoffice
type Useractivityoutofoffice struct { 
    // Active - Whether the user is currently out of office
    Active bool `json:"active"`


    // ModifiedDate - The date the out of office state was last modified. Date time is represented as an ISO-8601 string
    ModifiedDate time.Time `json:"modifiedDate"`

}

// String returns a JSON representation of the model
func (o *Useractivityoutofoffice) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivityoutofoffice) MarshalJSON() ([]byte, error) {
    type Alias Useractivityoutofoffice

    if UseractivityoutofofficeMarshalled {
        return []byte("{}"), nil
    }
    UseractivityoutofofficeMarshalled = true

    return json.Marshal(&struct {
        
        Active bool `json:"active"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

