package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffrequestbulkevaluatetemplateitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffrequestbulkevaluatetemplateitemDud struct { 
    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Timeoffrequestbulkevaluatetemplateitem
type Timeoffrequestbulkevaluatetemplateitem struct { 
    // Id - The ID of the time-off request
    Id string `json:"id"`


    // User - The user associated with this time-off request
    User Userreference `json:"user"`


    // Status - The status of the time-off request
    Status string `json:"status"`


    // Substatus - The substatus of the time-off request
    Substatus string `json:"substatus"`


    // ProcessingStatus - The processing status of the time-off request
    ProcessingStatus string `json:"processingStatus"`


    // Metadata - The version metadata of the time-off request
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Timeoffrequestbulkevaluatetemplateitem) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffrequestbulkevaluatetemplateitem) MarshalJSON() ([]byte, error) {
    type Alias Timeoffrequestbulkevaluatetemplateitem

    if TimeoffrequestbulkevaluatetemplateitemMarshalled {
        return []byte("{}"), nil
    }
    TimeoffrequestbulkevaluatetemplateitemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        User Userreference `json:"user"`
        
        Status string `json:"status"`
        
        Substatus string `json:"substatus"`
        
        ProcessingStatus string `json:"processingStatus"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

