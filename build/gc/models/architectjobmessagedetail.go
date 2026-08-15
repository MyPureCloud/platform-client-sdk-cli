package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArchitectjobmessagedetailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArchitectjobmessagedetailDud struct { 
    VarType string `json:"type"`


    Url string `json:"url"`


    Method string `json:"method"`


    RequestBody string `json:"requestBody"`


    StatusCode int `json:"statusCode"`


    StatusMessage string `json:"statusMessage"`


    CorrelationId string `json:"correlationId"`


    ResponseBody string `json:"responseBody"`


    ErrorCode string `json:"errorCode"`


    ErrorMessage string `json:"errorMessage"`

}

// Architectjobmessagedetail
type Architectjobmessagedetail struct { 
    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Architectjobmessagedetail) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Architectjobmessagedetail) MarshalJSON() ([]byte, error) {
    type Alias Architectjobmessagedetail

    if ArchitectjobmessagedetailMarshalled {
        return []byte("{}"), nil
    }
    ArchitectjobmessagedetailMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

