package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AsyncconversationqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AsyncconversationqueryDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Asyncconversationquery
type Asyncconversationquery struct { 
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


    // Interval - Specifies the date and time range of data being queried. Results will include all conversations that had activity during the interval. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Limit - Specify number of results to be returned
    Limit int `json:"limit"`


    // StartOfDayIntervalMatching - Add a filter to only include conversations that started after the beginning of the interval start date (UTC)
    StartOfDayIntervalMatching bool `json:"startOfDayIntervalMatching"`

}

// String returns a JSON representation of the model
func (o *Asyncconversationquery) String() string {
     o.ConversationFilters = []Conversationdetailqueryfilter{{}} 
     o.SegmentFilters = []Segmentdetailqueryfilter{{}} 
     o.EvaluationFilters = []Evaluationdetailqueryfilter{{}} 
     o.SurveyFilters = []Surveydetailqueryfilter{{}} 
     o.ResolutionFilters = []Resolutiondetailqueryfilter{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Asyncconversationquery) MarshalJSON() ([]byte, error) {
    type Alias Asyncconversationquery

    if AsyncconversationqueryMarshalled {
        return []byte("{}"), nil
    }
    AsyncconversationqueryMarshalled = true

    return json.Marshal(&struct {
        
        ConversationFilters []Conversationdetailqueryfilter `json:"conversationFilters"`
        
        SegmentFilters []Segmentdetailqueryfilter `json:"segmentFilters"`
        
        EvaluationFilters []Evaluationdetailqueryfilter `json:"evaluationFilters"`
        
        SurveyFilters []Surveydetailqueryfilter `json:"surveyFilters"`
        
        ResolutionFilters []Resolutiondetailqueryfilter `json:"resolutionFilters"`
        
        Order string `json:"order"`
        
        OrderBy string `json:"orderBy"`
        
        Interval string `json:"interval"`
        
        Limit int `json:"limit"`
        
        StartOfDayIntervalMatching bool `json:"startOfDayIntervalMatching"`
        *Alias
    }{

        
        ConversationFilters: []Conversationdetailqueryfilter{{}},
        


        
        SegmentFilters: []Segmentdetailqueryfilter{{}},
        


        
        EvaluationFilters: []Evaluationdetailqueryfilter{{}},
        


        
        SurveyFilters: []Surveydetailqueryfilter{{}},
        


        
        ResolutionFilters: []Resolutiondetailqueryfilter{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

