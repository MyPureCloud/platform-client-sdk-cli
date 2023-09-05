package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmintegrationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmintegrationresponseDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Wfmintegrationresponse
type Wfmintegrationresponse struct { 
    


    // Name
    Name string `json:"name"`


    // Active - Whether integration state is active
    Active bool `json:"active"`


    

}

// String returns a JSON representation of the model
func (o *Wfmintegrationresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmintegrationresponse) MarshalJSON() ([]byte, error) {
    type Alias Wfmintegrationresponse

    if WfmintegrationresponseMarshalled {
        return []byte("{}"), nil
    }
    WfmintegrationresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Active bool `json:"active"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

