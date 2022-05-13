package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchoutcomeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchoutcomeDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`


    


    

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
    Context Context `json:"context"`


    // Journey - The pattern of rules defining the filter of the outcome.
    Journey Journey `json:"journey"`


    // AssociatedValueField - The field from the event indicating the associated value.
    AssociatedValueField Associatedvaluefield `json:"associatedValueField"`


    


    // CreatedDate - Timestamp indicating when the outcome was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // ModifiedDate - Timestamp indicating when the outcome was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`

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
        
        Context Context `json:"context"`
        
        Journey Journey `json:"journey"`
        
        AssociatedValueField Associatedvaluefield `json:"associatedValueField"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

