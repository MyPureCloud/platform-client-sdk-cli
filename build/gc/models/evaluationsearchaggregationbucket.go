package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationsearchaggregationbucketMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationsearchaggregationbucketDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Evaluationsearchaggregationbucket
type Evaluationsearchaggregationbucket struct { 
    // Key - The key for this bucket
    Key string `json:"key"`


    // KeyAsString - The key as a string representation
    KeyAsString string `json:"keyAsString"`


    // DocumentCount - Number of documents in this bucket
    DocumentCount int `json:"documentCount"`


    // KeyValue - Numeric key value for date histograms
    KeyValue int `json:"keyValue"`


    // From - Lower bound for range buckets
    From float64 `json:"from"`


    // To - Upper bound for range buckets
    To float64 `json:"to"`


    // Value - Simple aggregation value
    Value float64 `json:"value"`


    // Count - Count of documents
    Count int `json:"count"`


    // Minimum - Minimum value in the aggregation
    Minimum float64 `json:"minimum"`


    // Maximum - Maximum value in the aggregation
    Maximum float64 `json:"maximum"`


    // Average - Average value in the aggregation
    Average float64 `json:"average"`


    // Sum - Sum of values in the aggregation
    Sum float64 `json:"sum"`


    // SubAggregations - Nested sub-aggregations
    SubAggregations map[string]Evaluationsearchaggregationresponse `json:"subAggregations"`

}

// String returns a JSON representation of the model
func (o *Evaluationsearchaggregationbucket) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.SubAggregations = map[string]Evaluationsearchaggregationresponse{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationsearchaggregationbucket) MarshalJSON() ([]byte, error) {
    type Alias Evaluationsearchaggregationbucket

    if EvaluationsearchaggregationbucketMarshalled {
        return []byte("{}"), nil
    }
    EvaluationsearchaggregationbucketMarshalled = true

    return json.Marshal(&struct {
        
        Key string `json:"key"`
        
        KeyAsString string `json:"keyAsString"`
        
        DocumentCount int `json:"documentCount"`
        
        KeyValue int `json:"keyValue"`
        
        From float64 `json:"from"`
        
        To float64 `json:"to"`
        
        Value float64 `json:"value"`
        
        Count int `json:"count"`
        
        Minimum float64 `json:"minimum"`
        
        Maximum float64 `json:"maximum"`
        
        Average float64 `json:"average"`
        
        Sum float64 `json:"sum"`
        
        SubAggregations map[string]Evaluationsearchaggregationresponse `json:"subAggregations"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        SubAggregations: map[string]Evaluationsearchaggregationresponse{"": {}},
        

        Alias: (*Alias)(u),
    })
}

