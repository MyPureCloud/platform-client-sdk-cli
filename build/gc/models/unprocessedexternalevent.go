package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnprocessedexternaleventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnprocessedexternaleventDud struct { 
    


    


    


    


    

}

// Unprocessedexternalevent
type Unprocessedexternalevent struct { 
    // Event - The event that failed processing.
    Event Externalevent `json:"event"`


    // OriginalRequestIndex - The index of the event in the original request.
    OriginalRequestIndex int `json:"originalRequestIndex"`


    // IsRetryable - Whether the error is retryable.
    IsRetryable bool `json:"isRetryable"`


    // ErrorMessage - The error message.
    ErrorMessage string `json:"errorMessage"`


    // StatusCode - The HTTP status code associated with the error.
    StatusCode int `json:"statusCode"`

}

// String returns a JSON representation of the model
func (o *Unprocessedexternalevent) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unprocessedexternalevent) MarshalJSON() ([]byte, error) {
    type Alias Unprocessedexternalevent

    if UnprocessedexternaleventMarshalled {
        return []byte("{}"), nil
    }
    UnprocessedexternaleventMarshalled = true

    return json.Marshal(&struct {
        
        Event Externalevent `json:"event"`
        
        OriginalRequestIndex int `json:"originalRequestIndex"`
        
        IsRetryable bool `json:"isRetryable"`
        
        ErrorMessage string `json:"errorMessage"`
        
        StatusCode int `json:"statusCode"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

