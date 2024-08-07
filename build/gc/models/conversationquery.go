package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationqueryDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Conversationquery
type Conversationquery struct { 
    // ConversationFilters - Filters that target conversation-level data
    ConversationFilters []Conversationdetailqueryfilter `json:"conversationFilters"`


    // SegmentFilters - Filters that target individual segments within a conversation
    SegmentFilters []Segmentdetailqueryfilter `json:"segmentFilters"`


    // EvaluationFilters - Filters that target evaluations
    EvaluationFilters []Evaluationdetailqueryfilter `json:"evaluationFilters"`


    // SurveyFilters - Filters that target surveys
    SurveyFilters []Surveydetailqueryfilter `json:"surveyFilters"`


    // ResolutionFilters - Filters that target resolutions
    ResolutionFilters []Resolutiondetailqueryfilter `json:"resolutionFilters"`


    // Order - Sort the result set in ascending/descending order. Default is ascending
    Order string `json:"order"`


    // OrderBy - Specify which data element within the result set to use for sorting. The options  to use as a basis for sorting the results: conversationStart, segmentStart, and segmentEnd. If not specified, the default is conversationStart
    OrderBy string `json:"orderBy"`


    // Interval - Specifies the date and time range of data being queried. Results will only include conversations that started on a day touched by the interval. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Aggregations - Include faceted search and aggregate roll-ups describing your search results. This does not function as a filter, but rather, summary data about the data matching your filters
    Aggregations []Analyticsqueryaggregation `json:"aggregations"`


    // Paging - Page size and number to control iterating through large result sets. Default page size is 25
    Paging Pagingspec `json:"paging"`

}

// String returns a JSON representation of the model
func (o *Conversationquery) String() string {
     o.ConversationFilters = []Conversationdetailqueryfilter{{}} 
     o.SegmentFilters = []Segmentdetailqueryfilter{{}} 
     o.EvaluationFilters = []Evaluationdetailqueryfilter{{}} 
     o.SurveyFilters = []Surveydetailqueryfilter{{}} 
     o.ResolutionFilters = []Resolutiondetailqueryfilter{{}} 
    
    
    
     o.Aggregations = []Analyticsqueryaggregation{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationquery) MarshalJSON() ([]byte, error) {
    type Alias Conversationquery

    if ConversationqueryMarshalled {
        return []byte("{}"), nil
    }
    ConversationqueryMarshalled = true

    return json.Marshal(&struct {
        
        ConversationFilters []Conversationdetailqueryfilter `json:"conversationFilters"`
        
        SegmentFilters []Segmentdetailqueryfilter `json:"segmentFilters"`
        
        EvaluationFilters []Evaluationdetailqueryfilter `json:"evaluationFilters"`
        
        SurveyFilters []Surveydetailqueryfilter `json:"surveyFilters"`
        
        ResolutionFilters []Resolutiondetailqueryfilter `json:"resolutionFilters"`
        
        Order string `json:"order"`
        
        OrderBy string `json:"orderBy"`
        
        Interval string `json:"interval"`
        
        Aggregations []Analyticsqueryaggregation `json:"aggregations"`
        
        Paging Pagingspec `json:"paging"`
        *Alias
    }{

        
        ConversationFilters: []Conversationdetailqueryfilter{{}},
        


        
        SegmentFilters: []Segmentdetailqueryfilter{{}},
        


        
        EvaluationFilters: []Evaluationdetailqueryfilter{{}},
        


        
        SurveyFilters: []Surveydetailqueryfilter{{}},
        


        
        ResolutionFilters: []Resolutiondetailqueryfilter{{}},
        


        


        


        


        
        Aggregations: []Analyticsqueryaggregation{{}},
        


        

        Alias: (*Alias)(u),
    })
}

