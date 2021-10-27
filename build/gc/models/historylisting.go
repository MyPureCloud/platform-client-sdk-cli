package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistorylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistorylistingDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Historylisting
type Historylisting struct { 
    // Id
    Id string `json:"id"`


    // Complete
    Complete bool `json:"complete"`


    // User
    User User `json:"user"`


    // Client
    Client Domainentityref `json:"client"`


    // ErrorMessage
    ErrorMessage string `json:"errorMessage"`


    // ErrorCode
    ErrorCode string `json:"errorCode"`


    // ErrorDetails
    ErrorDetails []Detail `json:"errorDetails"`


    // ErrorMessageParams
    ErrorMessageParams map[string]string `json:"errorMessageParams"`


    // ActionName - Action name
    ActionName string `json:"actionName"`


    // ActionStatus - Action status
    ActionStatus string `json:"actionStatus"`


    // Name
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // System
    System bool `json:"system"`


    // Started - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Started time.Time `json:"started"`


    // Completed - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Completed time.Time `json:"completed"`


    // Entities
    Entities []Historyentry `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // Total
    Total int `json:"total"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Historylisting) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.ErrorDetails = []Detail{{}} 
    
    
    
     o.ErrorMessageParams = map[string]string{"": ""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Entities = []Historyentry{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historylisting) MarshalJSON() ([]byte, error) {
    type Alias Historylisting

    if HistorylistingMarshalled {
        return []byte("{}"), nil
    }
    HistorylistingMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Complete bool `json:"complete"`
        
        User User `json:"user"`
        
        Client Domainentityref `json:"client"`
        
        ErrorMessage string `json:"errorMessage"`
        
        ErrorCode string `json:"errorCode"`
        
        ErrorDetails []Detail `json:"errorDetails"`
        
        ErrorMessageParams map[string]string `json:"errorMessageParams"`
        
        ActionName string `json:"actionName"`
        
        ActionStatus string `json:"actionStatus"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        System bool `json:"system"`
        
        Started time.Time `json:"started"`
        
        Completed time.Time `json:"completed"`
        
        Entities []Historyentry `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        Total int `json:"total"`
        
        PageNumber int `json:"pageNumber"`
        
        PageCount int `json:"pageCount"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        ErrorDetails: []Detail{{}},
        

        

        
        ErrorMessageParams: map[string]string{"": ""},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Entities: []Historyentry{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

