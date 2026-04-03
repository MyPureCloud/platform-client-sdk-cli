package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateexternaleventsconfigurationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateexternaleventsconfigurationrequestDud struct { 
    


    

}

// Updateexternaleventsconfigurationrequest
type Updateexternaleventsconfigurationrequest struct { 
    // Name - The name of the external event configuration.
    Name string `json:"name"`


    // Description - A description of the external event configuration.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Updateexternaleventsconfigurationrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateexternaleventsconfigurationrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateexternaleventsconfigurationrequest

    if UpdateexternaleventsconfigurationrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateexternaleventsconfigurationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

