package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationsearchsubaggregationdtoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationsearchsubaggregationdtoDud struct { 
    


    


    


    


    


    


    


    

}

// Evaluationsearchsubaggregationdto
type Evaluationsearchsubaggregationdto struct { 
    // Name - The name of the aggregation
    Name string `json:"name"`


    // Field - The field to aggregate on.ALLOWED FIELDS BY AGGREGATION TYPE:TERM/COUNT: evaluationStatus, aiScoringFailureType, questionAiAnswerFailureType,TERM: formId, formIdReleased, contextId, questionGroupId, questionId, questionAnswerId, released, questionGroupMarkedNA, questionMarkedNA, questionAiScored, questionEaScored,SUM/AVERAGE/STATS/RANGE: totalScore, totalCriticalScore, questionGroupScorePercentage,criticalQuestionGroupScorePercentage, questionScore,SUM: disputeCount, rescoreCount, eaSuggestionCount, eaAcceptedSuggestionCount,aiSuggestionCount, aiAcceptedSuggestionCount,DATE_HISTOGRAM: conversationDate, createdDate, submittedDate, releaseDate
    Field string `json:"field"`


    // VarType - The Type of Aggregation to Perform
    VarType string `json:"type"`


    // Size - The size limit for term aggregations, 100 size limit for all fields
    Size int `json:"size"`


    // CalendarInterval - The calendar interval for date histogram aggregations
    CalendarInterval string `json:"calendarInterval"`


    // Format - The format for date histogram aggregations
    Format string `json:"format"`


    // Ranges - The ranges for range aggregations
    Ranges []Queryapisearchaggregationrange `json:"ranges"`


    // SubAggregations - Sub-aggregations to be computed within this aggregation
    SubAggregations []Evaluationsearchsubaggregationdto `json:"subAggregations"`

}

// String returns a JSON representation of the model
func (o *Evaluationsearchsubaggregationdto) String() string {
    
    
    
    
    
    
     o.Ranges = []Queryapisearchaggregationrange{{}} 
     o.SubAggregations = []Evaluationsearchsubaggregationdto{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationsearchsubaggregationdto) MarshalJSON() ([]byte, error) {
    type Alias Evaluationsearchsubaggregationdto

    if EvaluationsearchsubaggregationdtoMarshalled {
        return []byte("{}"), nil
    }
    EvaluationsearchsubaggregationdtoMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Field string `json:"field"`
        
        VarType string `json:"type"`
        
        Size int `json:"size"`
        
        CalendarInterval string `json:"calendarInterval"`
        
        Format string `json:"format"`
        
        Ranges []Queryapisearchaggregationrange `json:"ranges"`
        
        SubAggregations []Evaluationsearchsubaggregationdto `json:"subAggregations"`
        *Alias
    }{

        


        


        


        


        


        


        
        Ranges: []Queryapisearchaggregationrange{{}},
        


        
        SubAggregations: []Evaluationsearchsubaggregationdto{{}},
        

        Alias: (*Alias)(u),
    })
}

