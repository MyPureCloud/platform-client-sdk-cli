package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntentclassifierinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntentclassifierinfoDud struct { 
    


    


    

}

// Intentclassifierinfo
type Intentclassifierinfo struct { 
    // ClassifierId - The intent classifier ID
    ClassifierId string `json:"classifierId"`


    // Version - The intent classifier version
    Version string `json:"version"`


    // State - The intent classifier state
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Intentclassifierinfo) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intentclassifierinfo) MarshalJSON() ([]byte, error) {
    type Alias Intentclassifierinfo

    if IntentclassifierinfoMarshalled {
        return []byte("{}"), nil
    }
    IntentclassifierinfoMarshalled = true

    return json.Marshal(&struct {
        
        ClassifierId string `json:"classifierId"`
        
        Version string `json:"version"`
        
        State string `json:"state"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

