package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomattributesbulkupdateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomattributesbulkupdateresponseDud struct { 
    


    


    

}

// Customattributesbulkupdateresponse
type Customattributesbulkupdateresponse struct { 
    // Attributes - The Custom Attributes record.
    Attributes Customattributes `json:"attributes"`


    // StatusCode - The status code result of updating the Custom Attributes record.
    StatusCode int `json:"statusCode"`


    // ErrorMessage - The error message if an error occurs while updating the record, otherwise null.
    ErrorMessage string `json:"errorMessage"`

}

// String returns a JSON representation of the model
func (o *Customattributesbulkupdateresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customattributesbulkupdateresponse) MarshalJSON() ([]byte, error) {
    type Alias Customattributesbulkupdateresponse

    if CustomattributesbulkupdateresponseMarshalled {
        return []byte("{}"), nil
    }
    CustomattributesbulkupdateresponseMarshalled = true

    return json.Marshal(&struct {
        
        Attributes Customattributes `json:"attributes"`
        
        StatusCode int `json:"statusCode"`
        
        ErrorMessage string `json:"errorMessage"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

