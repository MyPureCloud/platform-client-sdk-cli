package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowsquerycriteriaresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowsquerycriteriaresponseDud struct { 
    


    


    


    


    

}

// Flowsquerycriteriaresponse - The response for QueryCapabilities which contains the allowed criteria, flow types and action types for the organization.
type Flowsquerycriteriaresponse struct { 
    // Criteria - The is a list of allowed criteria to query on.
    Criteria []Querycriteria `json:"criteria"`


    // FlowTypes - The is a list of flow types the organization has access to.
    FlowTypes []string `json:"flowTypes"`


    // ActionTypes - The is a list of action types the organization has access to.
    ActionTypes []string `json:"actionTypes"`


    // ErrorCodes - The is a list of potential error codes the organization may encounter.
    ErrorCodes []string `json:"errorCodes"`


    // WarningCodes - The is a list of potential warning codes the organization may encounter.
    WarningCodes []string `json:"warningCodes"`

}

// String returns a JSON representation of the model
func (o *Flowsquerycriteriaresponse) String() string {
     o.Criteria = []Querycriteria{{}} 
     o.FlowTypes = []string{""} 
     o.ActionTypes = []string{""} 
     o.ErrorCodes = []string{""} 
     o.WarningCodes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowsquerycriteriaresponse) MarshalJSON() ([]byte, error) {
    type Alias Flowsquerycriteriaresponse

    if FlowsquerycriteriaresponseMarshalled {
        return []byte("{}"), nil
    }
    FlowsquerycriteriaresponseMarshalled = true

    return json.Marshal(&struct {
        
        Criteria []Querycriteria `json:"criteria"`
        
        FlowTypes []string `json:"flowTypes"`
        
        ActionTypes []string `json:"actionTypes"`
        
        ErrorCodes []string `json:"errorCodes"`
        
        WarningCodes []string `json:"warningCodes"`
        *Alias
    }{

        
        Criteria: []Querycriteria{{}},
        


        
        FlowTypes: []string{""},
        


        
        ActionTypes: []string{""},
        


        
        ErrorCodes: []string{""},
        


        
        WarningCodes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

