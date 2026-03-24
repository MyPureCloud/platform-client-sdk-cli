package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseplanupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseplanupdateDud struct { 
    


    


    


    


    


    


    


    

}

// Caseplanupdate
type Caseplanupdate struct { 
    // Name - The name of the Caseplan. Valid length between 3 and 256 characters.
    Name string `json:"name"`


    // DefaultDueDurationInSeconds - The default due duration in seconds for Cases created from the Caseplan. Valid range is between 1 and 31536000 seconds.
    DefaultDueDurationInSeconds int `json:"defaultDueDurationInSeconds"`


    // DefaultTtlSeconds - The default TTL in seconds for Cases created from the Caseplan. Valid range is between 86400 and 31536000 seconds.
    DefaultTtlSeconds int `json:"defaultTtlSeconds"`


    // ReferencePrefix - The reference of the Caseplan. Valid length between 2 and 8 alphanumeric characters.
    ReferencePrefix string `json:"referencePrefix"`


    // CustomerIntentId - The ID of the customer intent associated with this Caseplan.
    CustomerIntentId string `json:"customerIntentId"`


    // Description - The description of the Caseplan. Maximum length of 512 characters.
    Description string `json:"description"`


    // DefaultCaseOwnerId - The ID of the default owner of a Case created from the Caseplan. Must be a valid UUID.
    DefaultCaseOwnerId string `json:"defaultCaseOwnerId"`


    // DivisionId - The ID of the division the Caseplan belongs to. If divisionId is null or '*', the Caseplan will be divisionless.
    DivisionId string `json:"divisionId"`

}

// String returns a JSON representation of the model
func (o *Caseplanupdate) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseplanupdate) MarshalJSON() ([]byte, error) {
    type Alias Caseplanupdate

    if CaseplanupdateMarshalled {
        return []byte("{}"), nil
    }
    CaseplanupdateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        DefaultDueDurationInSeconds int `json:"defaultDueDurationInSeconds"`
        
        DefaultTtlSeconds int `json:"defaultTtlSeconds"`
        
        ReferencePrefix string `json:"referencePrefix"`
        
        CustomerIntentId string `json:"customerIntentId"`
        
        Description string `json:"description"`
        
        DefaultCaseOwnerId string `json:"defaultCaseOwnerId"`
        
        DivisionId string `json:"divisionId"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

