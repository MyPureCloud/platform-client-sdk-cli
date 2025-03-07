package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WidgetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WidgetDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Widget
type Widget struct { 
    // Row - The row number for the specific dashboard widget configuration.
    Row int `json:"row"`


    // Column - The column number for the specific dashboard widget configuration.
    Column int `json:"column"`


    // Title - The title for the dashboard widget configuration.
    Title string `json:"title"`


    // VarType - The type of dashboard widget configuration.
    VarType string `json:"type"`


    // Metrics - The list of metrics for the dashboard widget configuration.
    Metrics []string `json:"metrics"`


    // DisplayText - The display text for the dashboard widget configuration.
    DisplayText string `json:"displayText"`


    // DisplayTextColor - The color of the display text for the dashboard widget configuration in RGB hexadecimal format (for example \"#FF0000\" represents red).
    DisplayTextColor string `json:"displayTextColor"`


    // WebContentUrl - The external web URL for the dashboard widget configuration.
    WebContentUrl string `json:"webContentUrl"`


    // SplitFilters - Indicates each filter to be displayed individually.
    SplitFilters bool `json:"splitFilters"`


    // SplitByMediaType - Indicates that data for each media type should be shown individually.
    SplitByMediaType bool `json:"splitByMediaType"`


    // ShowLongest - Indicates the display be the longest time.
    ShowLongest bool `json:"showLongest"`


    // DisplayAsTable - Indicates the widget to be displayed as table.
    DisplayAsTable bool `json:"displayAsTable"`


    // ShowDuration - Indicates the display to include duration.
    ShowDuration bool `json:"showDuration"`


    // SortOrder - The sort order of the table.
    SortOrder string `json:"sortOrder"`


    // SortKey - The sort key of the table.
    SortKey string `json:"sortKey"`


    // EntityLimit - Indicates the limit of displayed entities.
    EntityLimit int `json:"entityLimit"`


    // DisplayAggregates - Indicates whether to display aggregate across all entity and media type combination.
    DisplayAggregates bool `json:"displayAggregates"`


    // IsFullWidth - Indicates whether a widget should take the full width of a dashboard or be shown only in a single slot.
    IsFullWidth bool `json:"isFullWidth"`


    // ShowPercentageChange - Indicates whether a widget should show the percentage diff between two values.
    ShowPercentageChange bool `json:"showPercentageChange"`


    // ShowProfilePicture - Indicates whether a widget should show the profile picture of an agent.
    ShowProfilePicture bool `json:"showProfilePicture"`


    // Filter - The filters to be applied for dashboard widget configuration
    Filter Viewfilter `json:"filter"`


    // Periods - The list of periods for the dashboard widget configuration
    Periods []string `json:"periods"`


    // MediaTypes - The list of media types for the dashboard widget configuration
    MediaTypes []string `json:"mediaTypes"`


    // Warnings - List of warnings for dashboard widget configuration
    Warnings []Warning `json:"warnings"`


    // ShowTimeInStatus - Indicates the show time in status of a widget configuration.
    ShowTimeInStatus bool `json:"showTimeInStatus"`


    // ShowOfflineAgents - Indicates to show offline agent widget.
    ShowOfflineAgents bool `json:"showOfflineAgents"`


    // SelectedStatuses - Indicates the selected statuses used to filter the agent widget in the dashboard.
    SelectedStatuses []string `json:"selectedStatuses"`


    // SelectedSegmentTypes - Indicates the selected segment types used to filter the agent activity in the dashboard.
    SelectedSegmentTypes []string `json:"selectedSegmentTypes"`


    // AgentInteractionSortOrder - The sort order of the interactions in the agent status widget.
    AgentInteractionSortOrder string `json:"agentInteractionSortOrder"`

}

// String returns a JSON representation of the model
func (o *Widget) String() string {
    
    
    
    
     o.Metrics = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Periods = []string{""} 
     o.MediaTypes = []string{""} 
     o.Warnings = []Warning{{}} 
    
    
     o.SelectedStatuses = []string{""} 
     o.SelectedSegmentTypes = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Widget) MarshalJSON() ([]byte, error) {
    type Alias Widget

    if WidgetMarshalled {
        return []byte("{}"), nil
    }
    WidgetMarshalled = true

    return json.Marshal(&struct {
        
        Row int `json:"row"`
        
        Column int `json:"column"`
        
        Title string `json:"title"`
        
        VarType string `json:"type"`
        
        Metrics []string `json:"metrics"`
        
        DisplayText string `json:"displayText"`
        
        DisplayTextColor string `json:"displayTextColor"`
        
        WebContentUrl string `json:"webContentUrl"`
        
        SplitFilters bool `json:"splitFilters"`
        
        SplitByMediaType bool `json:"splitByMediaType"`
        
        ShowLongest bool `json:"showLongest"`
        
        DisplayAsTable bool `json:"displayAsTable"`
        
        ShowDuration bool `json:"showDuration"`
        
        SortOrder string `json:"sortOrder"`
        
        SortKey string `json:"sortKey"`
        
        EntityLimit int `json:"entityLimit"`
        
        DisplayAggregates bool `json:"displayAggregates"`
        
        IsFullWidth bool `json:"isFullWidth"`
        
        ShowPercentageChange bool `json:"showPercentageChange"`
        
        ShowProfilePicture bool `json:"showProfilePicture"`
        
        Filter Viewfilter `json:"filter"`
        
        Periods []string `json:"periods"`
        
        MediaTypes []string `json:"mediaTypes"`
        
        Warnings []Warning `json:"warnings"`
        
        ShowTimeInStatus bool `json:"showTimeInStatus"`
        
        ShowOfflineAgents bool `json:"showOfflineAgents"`
        
        SelectedStatuses []string `json:"selectedStatuses"`
        
        SelectedSegmentTypes []string `json:"selectedSegmentTypes"`
        
        AgentInteractionSortOrder string `json:"agentInteractionSortOrder"`
        *Alias
    }{

        


        


        


        


        
        Metrics: []string{""},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Periods: []string{""},
        


        
        MediaTypes: []string{""},
        


        
        Warnings: []Warning{{}},
        


        


        


        
        SelectedStatuses: []string{""},
        


        
        SelectedSegmentTypes: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

