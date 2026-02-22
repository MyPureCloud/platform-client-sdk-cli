package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FollowersfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FollowersfilterDud struct { 
    


    


    


    

}

// Followersfilter
type Followersfilter struct { 
    // Operator - The comparison operator for follower count filtering.
    Operator string `json:"operator"`


    // From - The minimum followers count
    From int `json:"from"`


    // To - The maximum followers count
    To int `json:"to"`


    // Value - The specific followers count value
    Value int `json:"value"`

}

// String returns a JSON representation of the model
func (o *Followersfilter) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Followersfilter) MarshalJSON() ([]byte, error) {
    type Alias Followersfilter

    if FollowersfilterMarshalled {
        return []byte("{}"), nil
    }
    FollowersfilterMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        From int `json:"from"`
        
        To int `json:"to"`
        
        Value int `json:"value"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

