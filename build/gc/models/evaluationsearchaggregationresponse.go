package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationsearchaggregationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationsearchaggregationresponseDud struct { 
    


    


    


    


    


    


    


    


    

}

// Evaluationsearchaggregationresponse
type Evaluationsearchaggregationresponse struct { 
    // Value - Simple aggregation value (for SUM, COUNT, AVERAGE, etc.)
    Value float64 `json:"value"`


    // Count - Count of documents
    Count int `json:"count"`


    // Minimum - Minimum value
    Minimum float64 `json:"minimum"`


    // Maximum - Maximum value
    Maximum float64 `json:"maximum"`


    // Average - Average value
    Average float64 `json:"average"`


    // Sum - Total Sum
    Sum float64 `json:"sum"`


    // DocumentCountErrorUpperBound - Upper bound estimate of the error in document count for this aggregation
    DocumentCountErrorUpperBound int `json:"documentCountErrorUpperBound"`


    // SumOtherDocumentCount - Total document count for buckets not included in the response due to size limits
    SumOtherDocumentCount int `json:"sumOtherDocumentCount"`


    // Buckets - List of aggregation buckets
    Buckets []Evaluationsearchaggregationbucket `json:"buckets"`

}

// String returns a JSON representation of the model
func (o *Evaluationsearchaggregationresponse) String() string {
    
    
    
    
    
    
    
    
     o.Buckets = []Evaluationsearchaggregationbucket{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationsearchaggregationresponse) MarshalJSON() ([]byte, error) {
    type Alias Evaluationsearchaggregationresponse

    if EvaluationsearchaggregationresponseMarshalled {
        return []byte("{}"), nil
    }
    EvaluationsearchaggregationresponseMarshalled = true

    return json.Marshal(&struct {
        
        Value float64 `json:"value"`
        
        Count int `json:"count"`
        
        Minimum float64 `json:"minimum"`
        
        Maximum float64 `json:"maximum"`
        
        Average float64 `json:"average"`
        
        Sum float64 `json:"sum"`
        
        DocumentCountErrorUpperBound int `json:"documentCountErrorUpperBound"`
        
        SumOtherDocumentCount int `json:"sumOtherDocumentCount"`
        
        Buckets []Evaluationsearchaggregationbucket `json:"buckets"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        Buckets: []Evaluationsearchaggregationbucket{{}},
        

        Alias: (*Alias)(u),
    })
}

