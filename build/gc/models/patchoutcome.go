package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchoutcomeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchoutcomeDud struct { 
    


    


    


    


    


    


    


    

}

// Patchoutcome
type Patchoutcome struct { 
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
    Context Patchcontext `json:"context"`


    // Journey - The pattern of rules defining the filter of the outcome.
    Journey Patchjourney `json:"journey"`


    // AssociatedValueField - The field from the event indicating the associated value.
    AssociatedValueField Patchassociatedvaluefield `json:"associatedValueField"`

}

// String returns a JSON representation of the model
func (o *Patchoutcome) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchoutcome) MarshalJSON() ([]byte, error) {
    type Alias Patchoutcome

    if PatchoutcomeMarshalled {
        return []byte("{}"), nil
    }
    PatchoutcomeMarshalled = true

    return json.Marshal(&struct {
        
        IsActive bool `json:"isActive"`
        
        DisplayName string `json:"displayName"`
        
        Version int `json:"version"`
        
        Description string `json:"description"`
        
        IsPositive bool `json:"isPositive"`
        
        Context Patchcontext `json:"context"`
        
        Journey Patchjourney `json:"journey"`
        
        AssociatedValueField Patchassociatedvaluefield `json:"associatedValueField"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

