package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsphonenumberpatchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsphonenumberpatchrequestDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Smsphonenumberpatchrequest
type Smsphonenumberpatchrequest struct { 
    


    // SupportedContent - Defines the media SupportedContent profile configured for an MMS capable phone number.
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    

}

// String returns a JSON representation of the model
func (o *Smsphonenumberpatchrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsphonenumberpatchrequest) MarshalJSON() ([]byte, error) {
    type Alias Smsphonenumberpatchrequest

    if SmsphonenumberpatchrequestMarshalled {
        return []byte("{}"), nil
    }
    SmsphonenumberpatchrequestMarshalled = true

    return json.Marshal(&struct {
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

