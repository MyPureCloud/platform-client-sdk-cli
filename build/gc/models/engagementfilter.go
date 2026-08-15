package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EngagementfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EngagementfilterDud struct { 
    


    


    


    

}

// Engagementfilter
type Engagementfilter struct { 
    // Operator - The comparison operator for engagement metric filtering.
    Operator string `json:"operator"`


    // From - The inclusive lower bound of the engagement metric count. Required when operator is Between, not allowed otherwise.
    From int `json:"from"`


    // To - The inclusive upper bound of the engagement metric count. Required when operator is Between, not allowed otherwise.
    To int `json:"to"`


    // Value - The engagement metric count to compare against. Required for every operator except Between, not allowed for Between.
    Value int `json:"value"`

}

// String returns a JSON representation of the model
func (o *Engagementfilter) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Engagementfilter) MarshalJSON() ([]byte, error) {
    type Alias Engagementfilter

    if EngagementfilterMarshalled {
        return []byte("{}"), nil
    }
    EngagementfilterMarshalled = true

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

