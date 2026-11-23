package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesflowtypelanguagesupportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesflowtypelanguagesupportDud struct { 
    


    


    

}

// Capabilitiesflowtypelanguagesupport
type Capabilitiesflowtypelanguagesupport struct { 
    // IsAvailable
    IsAvailable bool `json:"isAvailable"`


    // UnavailableReason
    UnavailableReason Capabilitiesunavailablereason `json:"unavailableReason"`


    // MoreThanOneLanguage
    MoreThanOneLanguage bool `json:"moreThanOneLanguage"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesflowtypelanguagesupport) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesflowtypelanguagesupport) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesflowtypelanguagesupport

    if CapabilitiesflowtypelanguagesupportMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesflowtypelanguagesupportMarshalled = true

    return json.Marshal(&struct {
        
        IsAvailable bool `json:"isAvailable"`
        
        UnavailableReason Capabilitiesunavailablereason `json:"unavailableReason"`
        
        MoreThanOneLanguage bool `json:"moreThanOneLanguage"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

