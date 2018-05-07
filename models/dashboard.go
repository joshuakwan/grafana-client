package models

import "github.com/bitly/go-simplejson"

// Meta stands for the meta data of a grafana dashboard
type Meta struct {
	Type        string `json:"type,omitempty"`
	CanSave     bool   `json:"canSave,omitempty"`
	CanEdit     bool   `json:"canEdit,omitempty"`
	CanStar     bool   `json:"canStar,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Expires     string `json:"expires,omitempty"`
	Created     string `json:"created,omitempty"`
	Updated     string `json:"updated,omitempty"`
	UpdatedBy   string `json:"updatedBy,omitempty"`
	CreatedBy   string `json:"createdBy,omitempty"`
	Version     int    `json:"version,omitempty"`
	HasACL      bool   `json:"hasAcl,omitempty"`
	IsFolder    bool   `json:"isFolder,omitempty"`
	FolderID    int    `json:"folderId,omitempty"`
	FolderTitle string `json:"folderTitle,omitempty"`
	FolderURL   string `json:"folderUrl,omitempty"`
}

// PanelGridPosition defines the panel grid position of a dashboard panel
type PanelGridPosition struct {
	H int `json:"h,omitempty"`
	W int `json:"w,omitempty"`
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

// PanelLegend is the legend settings of a dashboard panel
type PanelLegend struct {
	Avg          bool `json:"avg"`
	Current      bool `json:"current"`
	Max          bool `json:"max"`
	Min          bool `json:"min"`
	Show         bool `json:"show"`
	Totals       bool `json:"totals"`
	Values       bool `json:"values"`
	HideEmpty    bool `json:"hideEmpty"`
	HideZero     bool `json:"hideZero"`
	AlignAsTable bool `json:"alignAsTable"`
	RightSide    bool `json:"rightSide"`
}

// Target encapsulates the target setting of a dashboard panel
type Target struct {
	RefID          string `json:"refId"`
	Expr           string `json:"expr"`
	LegendFormat   string `json:"legendFormat"`
	IntervalFactor int    `json:"intervalFactor"`
}

// PanelXAxis defines the X Axis position of a dashboard panel
type PanelXAxis struct {
	Show bool   `json:"show"`
	Mode string `json:"mode"`
}

// PanelYAxis defines the Y Axis position of a dashboard panel
type PanelYAxis struct {
	Format  string `json:"format"`
	LogBase int    `json:"logBase"`
	Show    bool   `json:"show"`
	Min     int    `json:"min"`
	Max     int    `json:"max"`
	Label   string `json:"label"`
}

// PanelTooltip stands for the settings of the tooltip of a dashboard panel
type PanelTooltip struct {
	MsResolution bool   `json:"msResolution"`
	Shared       bool   `json:"shared"`
	Sort         int    `json:"sort"`
	ValueType    string `json:"valueType"`
}

// PanelAlertConditionEvaluator holds the evaluator of a panel alert
type PanelAlertConditionEvaluator struct {
	Params []float64 `json:"params"`
	Type   string    `json:"type"`
}

// PanelAlertConditionOperator holds the operator of a panel alert
type PanelAlertConditionOperator struct {
	Type string `json:"type"`
}

// PanelAlertConditionQuery holds the query of a panel alert
type PanelAlertConditionQuery struct {
	DatasourceID int      `json:"datasourceId"`
	Model        *Target  `json:"model"`
	Params       []string `json:"params"`
}

// PanelAlertConditionReducer holds the reduce condition of a panel alert
type PanelAlertConditionReducer struct {
	Type   string   `json:"type"`
	Params []string `json:"params"`
}

// PanelAlertCondition defines the condition of triggering the alert
type PanelAlertCondition struct {
	Evaluator *PanelAlertConditionEvaluator `json:"evaluator"`
	Operator  *PanelAlertConditionOperator  `json:"operator"`
	Query     *PanelAlertConditionQuery     `json:"query"`
	Reducer   *PanelAlertConditionReducer   `json:"reducer"`
}

// PanelAlert is an alert object defined in the dashboard
type PanelAlert struct {
	Conditions          []*PanelAlertCondition `json:"conditions"`
	ExecutionErrorState string                 `json:"executionErrorState"`
	Frequency           string                 `json:"frequency"`
	Handler             int                    `json:"handler"`
	Message             string                 `json:"message"`
	Name                string                 `json:"name"`
	NoDataState         string                 `json:"noDataState"`
	Notifications       []*AlertNotification   `json:"notifications"`
}

// PanelThreshold is a threshold object defined in the dashboard
type PanelThreshold struct {
	ColorMode string `json:"colorMode,omitempty"`
	Fill      bool   `json:"fill,omitempty"`
	Line      bool   `json:"line,omitempty"`
	Op        string `json:"op,omitempty"`
	Value     int64  `json:"value,omitempty"`
}

// Panel stands for a dashboard panel
type Panel struct {
	Title           string             `json:"title,omitempty"`
	Datasource      string             `json:"datasource,omitempty"`
	Type            string             `json:"type,omitempty"`
	Format          string             `json:"format,omitempty"`
	Fill            int                `json:"fill,omitempty"`
	Linewidth       int                `json:"linewidth,omitempty"`
	Lines           bool               `json:"lines,omitempty"`
	Stack           bool               `json:"stack,omitempty"`
	Decimals        int                `json:"decimals,omitempty"`
	ValueName       string             `json:"valueName,omitempty"`
	NullPointMode   string             `json:"nullPointMode,omitempty"`
	Repeat          string             `json:"repeat,omitempty"`
	RepeatDirection string             `json:"repeatDirection,omitempty"`
	MinSpan         int                `json:"minSpan,omitempty"`
	GridPos         *PanelGridPosition `json:"gridPos,omitempty"`
	Legend          *PanelLegend       `json:"legend,omitempty"`
	Tooltip         *PanelTooltip      `json:"tooltip,omitempty"`
	Targets         []*Target          `json:"targets,omitempty"`
	XAxis           *PanelXAxis        `json:"xAxis,omitempty"`
	YAxes           []*PanelYAxis      `json:"yAxes,omitempty"`
	Editable        bool               `json:"editable,omitempty"`
	Error           bool               `json:"error,omitempty"`
	ID              int                `json:"id"`
	Alert           *PanelAlert        `json:"alert,omitempty"`
	Thresholds      []*PanelThreshold  `json:"thresholds,omitempty"`
}

// TemplateListCurrent stands for the current value of a variable
type TemplateListCurrent struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

// TemplateListOption stands for the option of a TemplateList
type TemplateListOption struct {
	Selected bool   `json:"selected"`
	Text     string `json:"text"`
	Value    string `json:"value"`
}

// TemplateList defines a dashboard variable
type TemplateList struct {
	Name       string                `json:"name"`
	Label      string                `json:"label"`
	Type       string                `json:"type"`
	Datasource string                `json:"datasource"`
	Query      string                `json:"query"`
	Refresh    int                   `json:"refresh"`
	Hide       int                   `json:"hide"`
	Multi      bool                  `json:"multi"`
	IncludeAll bool                  `json:"includeAll"`
	Current    *TemplateListCurrent  `json:"current"`
	Options    []*TemplateListOption `json:"options"`
}

// Template is the variable settings of a dashboard
type Template struct {
	List []*TemplateList `json:"list"`
}

// Time defines the "from-to" property of a dashboard
type Time struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// Dashboard contains the configuration body of a grafana dashboard
type Dashboard struct {
	ID            int      `json:"id"`
	Timezone      string   `json:"timezone"`
	Tags          []string `json:"tags"`
	Datasource    string   `json:"datasource"`
	Refresh       string   `json:"refresh"`
	SchemaVersion int      `json:"schemaVersion"`
	Time          Time     `json:"time"`
	//Panels        []*Panel `json:"panels"`
	Panels     *simplejson.Json `json:"panels"`
	Templating Template         `json:"templating"`
	Title      string           `json:"title"`
	UID        string           `json:"uid"`
	Version    int              `json:"version"`
}

// GrafanaDashboard consists of a meta object and a dashboard body object
type GrafanaDashboard struct {
	Meta      *Meta      `json:"meta"`
	Dashboard *Dashboard `json:"dashboard"`
}

// SearchResult holds the dashboard search result
type SearchResult struct {
	ID          int      `json:"id"`
	UID         string   `json:"uid"`
	Title       string   `json:"title"`
	URL         string   `json:"url"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	IsStarred   bool     `json:"isStarred"`
	FolderID    int      `json:"folderId"`
	FolderUID   string   `json:"folderUid"`
	FolderTitle string   `json:"folderTitle"`
	FolderURL   string   `json:"folderUrl"`
	URI         string   `json:"uri"` // deprecated in Grafana v5.0
}

// TagResult holds the tag search result
type TagResult struct {
	Term  string `json:"term"`
	Count int    `json:"count"`
}
