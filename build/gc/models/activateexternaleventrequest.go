package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivateexternaleventrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivateexternaleventrequestDud struct { 
    


    

}

// Activateexternaleventrequest - Request to activate an external event
type Activateexternaleventrequest struct { 
    // DisplayName - The display name of the external event
    DisplayName string `json:"displayName"`


    // Rank - The rank of the external event
    Rank int `json:"rank"`

}

// String returns a JSON representation of the model
func (o *Activateexternaleventrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activateexternaleventrequest) MarshalJSON() ([]byte, error) {
    type Alias Activateexternaleventrequest

    if ActivateexternaleventrequestMarshalled {
        return []byte("{}"), nil
    }
    ActivateexternaleventrequestMarshalled = true

    return json.Marshal(&struct {
        
        DisplayName string `json:"displayName"`
        
        Rank int `json:"rank"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

