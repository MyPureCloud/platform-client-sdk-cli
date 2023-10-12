package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomeachievedeventoutcomeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomeachievedeventoutcomeDud struct { 
    


    


    


    

}

// Outcomeachievedeventoutcome
type Outcomeachievedeventoutcome struct { 
    // Id
    Id string `json:"id"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // DisplayName - The display name of the outcome.
    DisplayName string `json:"displayName"`


    // Version - The version of the outcome.
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Outcomeachievedeventoutcome) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomeachievedeventoutcome) MarshalJSON() ([]byte, error) {
    type Alias Outcomeachievedeventoutcome

    if OutcomeachievedeventoutcomeMarshalled {
        return []byte("{}"), nil
    }
    OutcomeachievedeventoutcomeMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SelfUri string `json:"selfUri"`
        
        DisplayName string `json:"displayName"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

