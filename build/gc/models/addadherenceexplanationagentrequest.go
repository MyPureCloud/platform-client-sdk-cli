package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AddadherenceexplanationagentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AddadherenceexplanationagentrequestDud struct { 
    


    


    


    

}

// Addadherenceexplanationagentrequest
type Addadherenceexplanationagentrequest struct { 
    // VarType - The type of the adherence explanation
    VarType string `json:"type"`


    // StartDate - The start timestamp of the adherence explanation in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of the adherence explanation in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Notes - Notes about the adherence explanation
    Notes string `json:"notes"`

}

// String returns a JSON representation of the model
func (o *Addadherenceexplanationagentrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Addadherenceexplanationagentrequest) MarshalJSON() ([]byte, error) {
    type Alias Addadherenceexplanationagentrequest

    if AddadherenceexplanationagentrequestMarshalled {
        return []byte("{}"), nil
    }
    AddadherenceexplanationagentrequestMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Notes string `json:"notes"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

