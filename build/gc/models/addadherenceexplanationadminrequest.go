package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AddadherenceexplanationadminrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AddadherenceexplanationadminrequestDud struct { 
    


    


    


    


    

}

// Addadherenceexplanationadminrequest
type Addadherenceexplanationadminrequest struct { 
    // VarType - The type of the adherence explanation
    VarType string `json:"type"`


    // StartDate - The start timestamp of the adherence explanation in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of the adherence explanation in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Notes - Notes about the adherence explanation
    Notes string `json:"notes"`


    // Status - The status of the adherence explanation
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Addadherenceexplanationadminrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Addadherenceexplanationadminrequest) MarshalJSON() ([]byte, error) {
    type Alias Addadherenceexplanationadminrequest

    if AddadherenceexplanationadminrequestMarshalled {
        return []byte("{}"), nil
    }
    AddadherenceexplanationadminrequestMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Notes string `json:"notes"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

