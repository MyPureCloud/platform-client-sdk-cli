package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateutilizationlabelrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateutilizationlabelrequestDud struct { 
    


    

}

// Createutilizationlabelrequest
type Createutilizationlabelrequest struct { 
    // Name - The utilization label name.
    Name string `json:"name"`


    // Utilization - Org level utilization settings for the new label. If not specified, default utilization settings will be applied.
    Utilization Labelutilizationrequest `json:"utilization"`

}

// String returns a JSON representation of the model
func (o *Createutilizationlabelrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createutilizationlabelrequest) MarshalJSON() ([]byte, error) {
    type Alias Createutilizationlabelrequest

    if CreateutilizationlabelrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateutilizationlabelrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Utilization Labelutilizationrequest `json:"utilization"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

