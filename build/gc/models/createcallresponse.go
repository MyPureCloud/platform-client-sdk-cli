package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatecallresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatecallresponseDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Createcallresponse
type Createcallresponse struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Createcallresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createcallresponse) MarshalJSON() ([]byte, error) {
    type Alias Createcallresponse

    if CreatecallresponseMarshalled {
        return []byte("{}"), nil
    }
    CreatecallresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

