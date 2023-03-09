package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DncpatchcustomexclusioncolumnsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DncpatchcustomexclusioncolumnsrequestDud struct { 
    


    


    

}

// Dncpatchcustomexclusioncolumnsrequest
type Dncpatchcustomexclusioncolumnsrequest struct { 
    // Action - The action to perform
    Action string `json:"action"`


    // CustomExclusionColumnEntries - The list of custom exclusion column entries to Add to / Remove from the DNC list 
    CustomExclusionColumnEntries []string `json:"customExclusionColumnEntries"`


    // ExpirationDateTime - Expiration date for DNC customExclusionColumnEntries in yyyy-MM-ddTHH:mmZ format
    ExpirationDateTime string `json:"expirationDateTime"`

}

// String returns a JSON representation of the model
func (o *Dncpatchcustomexclusioncolumnsrequest) String() string {
    
     o.CustomExclusionColumnEntries = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dncpatchcustomexclusioncolumnsrequest) MarshalJSON() ([]byte, error) {
    type Alias Dncpatchcustomexclusioncolumnsrequest

    if DncpatchcustomexclusioncolumnsrequestMarshalled {
        return []byte("{}"), nil
    }
    DncpatchcustomexclusioncolumnsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Action string `json:"action"`
        
        CustomExclusionColumnEntries []string `json:"customExclusionColumnEntries"`
        
        ExpirationDateTime string `json:"expirationDateTime"`
        *Alias
    }{

        


        
        CustomExclusionColumnEntries: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

