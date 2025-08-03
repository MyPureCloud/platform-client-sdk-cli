package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsagequerysortbyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsagequerysortbyDud struct { 
    


    

}

// Usagequerysortby
type Usagequerysortby struct { 
    // FieldName - The name of the field you want to specify to sort by.
    FieldName string `json:"fieldName"`


    // Order - The order to sort by.
    Order string `json:"order"`

}

// String returns a JSON representation of the model
func (o *Usagequerysortby) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usagequerysortby) MarshalJSON() ([]byte, error) {
    type Alias Usagequerysortby

    if UsagequerysortbyMarshalled {
        return []byte("{}"), nil
    }
    UsagequerysortbyMarshalled = true

    return json.Marshal(&struct {
        
        FieldName string `json:"fieldName"`
        
        Order string `json:"order"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

