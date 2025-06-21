package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidejobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidejobDud struct { 
    Id string `json:"id"`


    Status string `json:"status"`


    Errors []Errorbody `json:"errors"`


    


    SelfUri string `json:"selfUri"`

}

// Guidejob
type Guidejob struct { 
    


    


    


    // Guide
    Guide Addressableentityref `json:"guide"`


    

}

// String returns a JSON representation of the model
func (o *Guidejob) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidejob) MarshalJSON() ([]byte, error) {
    type Alias Guidejob

    if GuidejobMarshalled {
        return []byte("{}"), nil
    }
    GuidejobMarshalled = true

    return json.Marshal(&struct {
        
        Guide Addressableentityref `json:"guide"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

