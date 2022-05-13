package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyeventssettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyeventssettingsDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Journeyeventssettings - Settings concerning journey events
type Journeyeventssettings struct { 
    // Enabled - Whether or not journey event collection is enabled.
    Enabled bool `json:"enabled"`


    // ExcludedQueryParameters - List of parameters to be excluded from the query string.
    ExcludedQueryParameters []string `json:"excludedQueryParameters"`


    // ShouldKeepUrlFragment - Whether or not to keep the URL fragment.
    ShouldKeepUrlFragment bool `json:"shouldKeepUrlFragment"`


    // SearchQueryParameters - List of query parameters used for search (e.g. 'q').
    SearchQueryParameters []string `json:"searchQueryParameters"`


    // PageviewConfig - Controls how the pageview events are tracked.
    PageviewConfig string `json:"pageviewConfig"`


    // ClickEvents - Tracks when and where a visitor clicks on a webpage.
    ClickEvents []Selectoreventtrigger `json:"clickEvents"`


    // FormsTrackEvents - Controls how the form submitted and form abandoned events are tracked after a visitor interacts with a form element.
    FormsTrackEvents []Formstracktrigger `json:"formsTrackEvents"`


    // IdleEvents - Tracks when and where a visitor becomes inactive on a webpage.
    IdleEvents []Idleeventtrigger `json:"idleEvents"`


    // InViewportEvents - Tracks when elements become visible or hidden on screen.
    InViewportEvents []Selectoreventtrigger `json:"inViewportEvents"`


    // ScrollDepthEvents - Tracks when a visitor scrolls to a specific percentage of a webpage.
    ScrollDepthEvents []Scrollpercentageeventtrigger `json:"scrollDepthEvents"`

}

// String returns a JSON representation of the model
func (o *Journeyeventssettings) String() string {
    
     o.ExcludedQueryParameters = []string{""} 
    
     o.SearchQueryParameters = []string{""} 
    
     o.ClickEvents = []Selectoreventtrigger{{}} 
     o.FormsTrackEvents = []Formstracktrigger{{}} 
     o.IdleEvents = []Idleeventtrigger{{}} 
     o.InViewportEvents = []Selectoreventtrigger{{}} 
     o.ScrollDepthEvents = []Scrollpercentageeventtrigger{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyeventssettings) MarshalJSON() ([]byte, error) {
    type Alias Journeyeventssettings

    if JourneyeventssettingsMarshalled {
        return []byte("{}"), nil
    }
    JourneyeventssettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        ExcludedQueryParameters []string `json:"excludedQueryParameters"`
        
        ShouldKeepUrlFragment bool `json:"shouldKeepUrlFragment"`
        
        SearchQueryParameters []string `json:"searchQueryParameters"`
        
        PageviewConfig string `json:"pageviewConfig"`
        
        ClickEvents []Selectoreventtrigger `json:"clickEvents"`
        
        FormsTrackEvents []Formstracktrigger `json:"formsTrackEvents"`
        
        IdleEvents []Idleeventtrigger `json:"idleEvents"`
        
        InViewportEvents []Selectoreventtrigger `json:"inViewportEvents"`
        
        ScrollDepthEvents []Scrollpercentageeventtrigger `json:"scrollDepthEvents"`
        *Alias
    }{

        


        
        ExcludedQueryParameters: []string{""},
        


        


        
        SearchQueryParameters: []string{""},
        


        


        
        ClickEvents: []Selectoreventtrigger{{}},
        


        
        FormsTrackEvents: []Formstracktrigger{{}},
        


        
        IdleEvents: []Idleeventtrigger{{}},
        


        
        InViewportEvents: []Selectoreventtrigger{{}},
        


        
        ScrollDepthEvents: []Scrollpercentageeventtrigger{{}},
        

        Alias: (*Alias)(u),
    })
}

