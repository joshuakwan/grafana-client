package models

import (
	"time"
)

type Meta struct {
	Type      string `json:"type"`
	CanSave   bool   `json:"canSave"`
	CanEdit   bool   `json:"canEdit"`
	CanStar   bool   `json:"canStar"`
	Slug      string `json:"slug"`
	Expires   string `json:"expires"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
	UpdatedBy string `json:"updatedBy"`
	CreatedBy string `json:"createdBy"`
	Version   int    `json:"version"`
}

type PanelGridPosition struct {
	H string `json:"h"`
	W string `json:"w"`
	X string `json:"x"`
	Y string `json:"y"`
}

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

type Target struct {
	RefID          string `json:"refId"`
	Expr           string `json:"expr"'`
	LegendFormat   string `json:"legendFormat"`
	IntervalFactor int    `json:"intervalFactor"`
}

type PanelXAxis struct {
	Show bool   `json:"show"`
	Mode string `json:"mode"`
}

type PanelYAxis struct {
	Format  string `json:"format"`
	LogBase int    `json:"logBase"`
	Show    bool   `json:"show"`
	Min     int    `json:"min"`
	Max     int    `json:"max"`
	Label   string `json:"label"`
}

type PanelTooltip struct {
	MsResolution bool   `json:"msResolution"`
	Shared       bool   `json:"shared"`
	Sort         int    `json:"sort"`
	ValueType    string `json:"valueType"`
}

type PanelAlertConditionEvaluator struct {
	Params []float64 `json:"params"`
	Type   string    `json:"type"`
}

type PanelAlertConditionOperator struct {
	Type string `json:"type"`
}

type PanelAlertConditionQuery struct {
	DatasourceID int      `json:"datasourceId"`
	Model        *Target  `json:"model"`
	Params       []string `json:"params"`
}

type PanelAlertConditionReducer struct {
	Type   string   `json:"type"`
	Params []string `json:"params"`
}

type PanelAlertCondition struct {
	Evaluator *PanelAlertConditionEvaluator `json:"evaluator"`
	Operator  *PanelAlertConditionOperator  `json:"operator"`
	Query     *PanelAlertConditionQuery     `json:"query"`
	Reducer   *PanelAlertConditionReducer   `json:"reducer"`
}

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

type PanelThreshold struct {
	ColorMode string `json:"colorMode"`
	Fill      bool   `json:"fill"`
	Line      bool   `json:"line"`
	Op        string `json:"op"`
	Value     int64  `json:"value"`
}

type Panel struct {
	Title           string             `json:"title"`
	Datasource      string             `json:"datasource"`
	Type            string             `json:"type"`
	Format          string             `json:"format"`
	Fill            int                `json:"fill"`
	Linewidth       int                `json:"linewidth"`
	Lines           bool               `json:"lines"`
	Stack           bool               `json:"stack"`
	Decimals        int                `json:"decimals"`
	ValueName       string             `json:"valueName"`
	NullPointMode   string             `json:"nullPointMode"`
	Repeat          string             `json:"repeat"`
	RepeatDirection string             `json:"repeatDirection"`
	MinSpan         int                `json:"minSpan"`
	GridPos         *PanelGridPosition `json:"gridPos"`
	Legend          *PanelLegend       `json:"legend"`
	Tooltip         *PanelTooltip      `json:"tooltip"`
	Targets         []*Target          `json:"targets"`
	XAxis           *PanelXAxis        `json:"xAxis"`
	YAxes           []*PanelYAxis      `json:"yAxes"`
	Editable        bool               `json:"editable"`
	Error           bool               `json:"error"`
	ID              int                `json:"id"`
	Alert           PanelAlert         `json:"alert"`
	Thresholds      []PanelThreshold   `json:"thresholds"`
}

type TemplateListCurrent struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type TemplateListOption struct {
	Selected bool   `json:"selected"`
	Text     string `json:"text"`
	Value    string `json:"value"`
}

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

type Template struct {
	List []*TemplateList `json:"list"`
}

type Dashboard struct {
	ID            int       `json:"id"`
	UID           string    `json:"uid"`
	Version       int       `json:"version"`
	Timezone      string    `json:"timezone"`
	Tags          []string  `json:"tags"`
	Title         string    `json:"title"`
	Datasource    string    `json:"datasource"`
	Refresh       string    `json:"refresh"`
	SchemaVersion int       `json:"schemaVersion"`
	Time          time.Time `json:"time"`
	Panels        []*Panel  `json:"panels"`
	Templating    Template  `json:"templating"`
}

type GrafanaDashboard struct {
	Meta      *Meta      `json:"meta"`
	Dashboard *Dashboard `json:"dashboard"`
}
