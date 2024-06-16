package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningscormresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningscormresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Learningscormresponse - Learning SCORM package
type Learningscormresponse struct { 
    


    // Status - The status of the SCORM package
    Status string `json:"status"`


    // ErrorCode - The error code of the SCORM package (on failure)
    ErrorCode string `json:"errorCode"`


    // ErrorMessage - The error message associated with the error code
    ErrorMessage string `json:"errorMessage"`


    // PercentageUnpacked - The percentage of the SCORM package that has been unpacked
    PercentageUnpacked float32 `json:"percentageUnpacked"`


    

}

// String returns a JSON representation of the model
func (o *Learningscormresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningscormresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningscormresponse

    if LearningscormresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningscormresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        ErrorCode string `json:"errorCode"`
        
        ErrorMessage string `json:"errorMessage"`
        
        PercentageUnpacked float32 `json:"percentageUnpacked"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

