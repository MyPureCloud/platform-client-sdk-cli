package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MaskingrulevalidaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MaskingrulevalidaterequestDud struct { 
    


    

}

// Maskingrulevalidaterequest
type Maskingrulevalidaterequest struct { 
    // Text - Text to mask.
    Text string `json:"text"`


    // Definition - Regex to be applied
    Definition string `json:"definition"`

}

// String returns a JSON representation of the model
func (o *Maskingrulevalidaterequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Maskingrulevalidaterequest) MarshalJSON() ([]byte, error) {
    type Alias Maskingrulevalidaterequest

    if MaskingrulevalidaterequestMarshalled {
        return []byte("{}"), nil
    }
    MaskingrulevalidaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Definition string `json:"definition"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

