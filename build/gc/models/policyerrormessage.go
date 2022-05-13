package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PolicyerrormessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PolicyerrormessageDud struct { 
    


    


    


    


    


    


    

}

// Policyerrormessage
type Policyerrormessage struct { 
    // StatusCode
    StatusCode int `json:"statusCode"`


    // UserMessage
    UserMessage interface{} `json:"userMessage"`


    // UserParamsMessage
    UserParamsMessage string `json:"userParamsMessage"`


    // ErrorCode
    ErrorCode string `json:"errorCode"`


    // CorrelationId
    CorrelationId string `json:"correlationId"`


    // UserParams
    UserParams []Userparam `json:"userParams"`


    // InsertDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    InsertDate time.Time `json:"insertDate"`

}

// String returns a JSON representation of the model
func (o *Policyerrormessage) String() string {
    
     o.UserMessage = Interface{} 
    
    
    
     o.UserParams = []Userparam{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Policyerrormessage) MarshalJSON() ([]byte, error) {
    type Alias Policyerrormessage

    if PolicyerrormessageMarshalled {
        return []byte("{}"), nil
    }
    PolicyerrormessageMarshalled = true

    return json.Marshal(&struct {
        
        StatusCode int `json:"statusCode"`
        
        UserMessage interface{} `json:"userMessage"`
        
        UserParamsMessage string `json:"userParamsMessage"`
        
        ErrorCode string `json:"errorCode"`
        
        CorrelationId string `json:"correlationId"`
        
        UserParams []Userparam `json:"userParams"`
        
        InsertDate time.Time `json:"insertDate"`
        *Alias
    }{

        


        
        UserMessage: Interface{},
        


        


        


        


        
        UserParams: []Userparam{{}},
        


        

        Alias: (*Alias)(u),
    })
}

