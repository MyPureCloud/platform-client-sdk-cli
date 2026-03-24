package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcesitedetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcesitedetailsDud struct { 
    


    

}

// V3sourcesitedetails
type V3sourcesitedetails struct { 
    // Id - The site's id.
    Id string `json:"id"`


    // Name - The site's display name.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *V3sourcesitedetails) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcesitedetails) MarshalJSON() ([]byte, error) {
    type Alias V3sourcesitedetails

    if V3sourcesitedetailsMarshalled {
        return []byte("{}"), nil
    }
    V3sourcesitedetailsMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

