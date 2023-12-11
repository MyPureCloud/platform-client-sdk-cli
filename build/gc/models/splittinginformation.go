package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SplittinginformationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SplittinginformationDud struct { 
    


    


    


    

}

// Splittinginformation
type Splittinginformation struct { 
    // Criteria - The splitting criteria type
    Criteria string `json:"criteria"`


    // CriteriaValue - The criteria value for the specified criteria type
    CriteriaValue string `json:"criteriaValue"`


    // CreateRemainderContactList - Whether to create remainder contact list
    CreateRemainderContactList bool `json:"createRemainderContactList"`


    // UseWaterfallRule - Whether to use waterfall rule
    UseWaterfallRule bool `json:"useWaterfallRule"`

}

// String returns a JSON representation of the model
func (o *Splittinginformation) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Splittinginformation) MarshalJSON() ([]byte, error) {
    type Alias Splittinginformation

    if SplittinginformationMarshalled {
        return []byte("{}"), nil
    }
    SplittinginformationMarshalled = true

    return json.Marshal(&struct {
        
        Criteria string `json:"criteria"`
        
        CriteriaValue string `json:"criteriaValue"`
        
        CreateRemainderContactList bool `json:"createRemainderContactList"`
        
        UseWaterfallRule bool `json:"useWaterfallRule"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

