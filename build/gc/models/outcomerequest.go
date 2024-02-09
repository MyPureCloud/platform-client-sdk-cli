package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomerequestDud struct { 
    


    


    


    


    


    


    


    

}

// Outcomerequest
type Outcomerequest struct { 
    // IsActive - Whether or not the outcome is active.
    IsActive bool `json:"isActive"`


    // DisplayName - The display name of the outcome.
    DisplayName string `json:"displayName"`


    // Version - The version of the outcome.
    Version int `json:"version"`


    // Description - A description of the outcome.
    Description string `json:"description"`


    // IsPositive - Whether or not the outcome is positive.
    IsPositive bool `json:"isPositive"`


    // Context - The context of the outcome.
    Context Requestcontext `json:"context"`


    // Journey - The pattern of rules defining the filter of the outcome.
    Journey Requestjourney `json:"journey"`


    // AssociatedValueField - The field from the event indicating the associated value.
    AssociatedValueField Associatedvaluefield `json:"associatedValueField"`

}

// String returns a JSON representation of the model
func (o *Outcomerequest) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomerequest) MarshalJSON() ([]byte, error) {
    type Alias Outcomerequest

    if OutcomerequestMarshalled {
        return []byte("{}"), nil
    }
    OutcomerequestMarshalled = true

    return json.Marshal(&struct {
        
        IsActive bool `json:"isActive"`
        
        DisplayName string `json:"displayName"`
        
        Version int `json:"version"`
        
        Description string `json:"description"`
        
        IsPositive bool `json:"isPositive"`
        
        Context Requestcontext `json:"context"`
        
        Journey Requestjourney `json:"journey"`
        
        AssociatedValueField Associatedvaluefield `json:"associatedValueField"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

