package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstateactivitycategorycountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstateactivitycategorycountDud struct { 
    


    

}

// Agentstateactivitycategorycount
type Agentstateactivitycategorycount struct { 
    // ActivityCategory - Activity category
    ActivityCategory string `json:"activityCategory"`


    // Count - Count of users with this activity category
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Agentstateactivitycategorycount) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstateactivitycategorycount) MarshalJSON() ([]byte, error) {
    type Alias Agentstateactivitycategorycount

    if AgentstateactivitycategorycountMarshalled {
        return []byte("{}"), nil
    }
    AgentstateactivitycategorycountMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCategory string `json:"activityCategory"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

