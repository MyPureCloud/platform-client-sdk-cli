package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UservideosettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UservideosettingsDud struct { 
    


    

}

// Uservideosettings - The settings for User Video
type Uservideosettings struct { 
    // AllowCamera - whether or not user camera is allowed
    AllowCamera bool `json:"allowCamera"`


    // AllowScreenShare - whether or not user screen share is allowed
    AllowScreenShare bool `json:"allowScreenShare"`

}

// String returns a JSON representation of the model
func (o *Uservideosettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Uservideosettings) MarshalJSON() ([]byte, error) {
    type Alias Uservideosettings

    if UservideosettingsMarshalled {
        return []byte("{}"), nil
    }
    UservideosettingsMarshalled = true

    return json.Marshal(&struct {
        
        AllowCamera bool `json:"allowCamera"`
        
        AllowScreenShare bool `json:"allowScreenShare"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

