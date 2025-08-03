package models

type Excalidraw struct {
	InitialData   *InitialDataExcalidraw
	TextElements  map[string]string
	EmbeddedFiles map[string]string
}

// Scene represents the top-level Excalidraw object.
type InitialDataExcalidraw struct {
	Type     string                     `json:"type"`
	Version  int                        `json:"version"`
	Source   string                     `json:"source"`
	Elements []*Element                 `json:"elements"`
	AppState *AppState                  `json:"appState"`
	Files    map[string]*ExcalidrawFile `json:"files"`
}

type ElementType string

const (
	ElementTypeRectangle  ElementType = "rectangle"
	ElementTypeEllipse    ElementType = "ellipse"
	ElementTypeDiamond    ElementType = "diamond"
	ElementTypeArrow      ElementType = "arrow"
	ElementTypeLine       ElementType = "line"
	ElementTypeText       ElementType = "text"
	ElementTypeImage      ElementType = "image"
	ElementTypeFrame      ElementType = "frame"
	ElementTypeFreedraw   ElementType = "freedraw"
	ElementTypeSelection  ElementType = "selection"
	ElementTypeMagicFrame ElementType = "magic-frame"
)

// Element represents a single shape or text box.
type Element struct {
	ID              string          `json:"id"`
	Type            ElementType     `json:"type"`
	X               float64         `json:"x"`
	Y               float64         `json:"y"`
	Width           float64         `json:"width"`
	Height          float64         `json:"height"`
	Angle           float64         `json:"angle"`
	StrokeColor     string          `json:"strokeColor"`
	BackgroundColor string          `json:"backgroundColor"`
	FillStyle       string          `json:"fillStyle"`
	StrokeWidth     float64         `json:"strokeWidth"`
	StrokeStyle     string          `json:"strokeStyle"`
	Roughness       float64         `json:"roughness"`
	Opacity         float64         `json:"opacity"`
	GroupIds        []string        `json:"groupIds"`
	FrameId         *string         `json:"frameId"`
	Index           string          `json:"index"`
	Roundness       *Roundness      `json:"roundness"`
	Seed            int64           `json:"seed"`
	Version         int             `json:"version"`
	VersionNonce    int64           `json:"versionNonce"`
	IsDeleted       bool            `json:"isDeleted"`
	BoundElements   *[]BoundElement `json:"boundElements"`

	Updated       int64   `json:"updated"`
	Link          *string `json:"link"`
	Locked        bool    `json:"locked"`
	Text          string  `json:"text"`
	FontSize      float64 `json:"fontSize,omitempty"`
	FontFamily    int     `json:"fontFamily,omitempty"`
	TextAlign     string  `json:"textAlign,omitempty"`
	VerticalAlign string  `json:"verticalAlign,omitempty"`
	ContainerID   *string `json:"containerId,omitempty"`

	// Image-specific properties (only for Type = "image")
	FileID string    `json:"fileId,omitempty"` // ID of the associated image file
	Scale  []float64 `json:"scale,omitempty"`  // [xScale, yScale] for image flipping, e.g., [-1, 1]

	// Linear element-specific properties (only for Type = "arrow", "line", "freedraw")
	Points       [][]float64 `json:"points,omitempty"`       // Array of [x, y] coordinates for path-based elements
	StartBinding *Binding    `json:"startBinding,omitempty"` // Binding at the start of an arrow/line
	EndBinding   *Binding    `json:"endBinding,omitempty"`   // Binding at the end of an arrow/line
	Arrowheads   []string    `json:"arrowheads,omitempty"`   // Array of arrow head types, e.g., ["arrow", "bar"]

	// Frame-specific properties (only for Type = "frame")
	Children []string `json:"children,omitempty"` // IDs of elements contained within the frame
}

// Roundness represents the roundness settings of an element
type Roundness struct {
	Type  int     `json:"type"`
	Value float64 `json:"value,omitempty"`
}

// BoundElement represents an element bound to another element
type BoundElement struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// AppState represents the current state of the Excalidraw app
type AppState struct {
	Theme                      string          `json:"theme"`
	ViewBackgroundColor        string          `json:"viewBackgroundColor"`
	CurrentItemStrokeColor     string          `json:"currentItemStrokeColor"`
	CurrentItemBackgroundColor string          `json:"currentItemBackgroundColor"`
	CurrentItemFillStyle       string          `json:"currentItemFillStyle"`
	CurrentItemStrokeWidth     float64         `json:"currentItemStrokeWidth"`
	CurrentItemStrokeStyle     string          `json:"currentItemStrokeStyle"`
	CurrentItemRoughness       float64         `json:"currentItemRoughness"`
	CurrentItemOpacity         float64         `json:"currentItemOpacity"`
	CurrentItemFontFamily      int             `json:"currentItemFontFamily"`
	CurrentItemFontSize        float64         `json:"currentItemFontSize"`
	CurrentItemTextAlign       string          `json:"currentItemTextAlign"`
	CurrentItemStartArrowhead  *string         `json:"currentItemStartArrowhead"`
	CurrentItemEndArrowhead    string          `json:"currentItemEndArrowhead"`
	CurrentItemArrowType       string          `json:"currentItemArrowType"`
	ScrollX                    float64         `json:"scrollX"`
	ScrollY                    float64         `json:"scrollY"`
	Zoom                       *Zoom           `json:"zoom"`
	CurrentItemRoundness       string          `json:"currentItemRoundness"`
	GridSize                   float64         `json:"gridSize"`
	GridStep                   float64         `json:"gridStep"`
	GridModeEnabled            bool            `json:"gridModeEnabled"`
	GridColor                  *GridColor      `json:"gridColor"`
	CurrentStrokeOptions       *interface{}    `json:"currentStrokeOptions"`
	FrameRendering             *FrameRendering `json:"frameRendering"`
	ObjectsSnapModeEnabled     bool            `json:"objectsSnapModeEnabled"`
	ActiveTool                 *ActiveTool     `json:"activeTool"`
}

type Zoom struct {
	Value float64 `json:"value"`
}

type GridColor struct {
	Bold    string `json:"Bold"`
	Regular string `json:"Regular"`
}

type FrameRendering struct {
	Enabled bool `json:"enabled"`
	Clip    bool `json:"clip"`
	Name    bool `json:"name"`
	Outline bool `json:"outline"`
}

type ActiveTool struct {
	Type           string  `json:"type"`
	CustomType     *string `json:"customType"`
	Locked         bool    `json:"locked"`
	FromSelection  bool    `json:"fromSelection"`
	LastActiveTool *string `json:"lastActiveTool"`
}

// File represents an embedded file's metadata and data.
type ExcalidrawFile struct {
	MimeType string `json:"mimeType"`
	ID       string `json:"id"`
	DataURL  string `json:"dataURL"`
	Created  int64  `json:"created"`
}

type Binding struct {
	ElementID string  `json:"elementId"` // ID of the element being bound to
	Focus     float64 `json:"focus"`     // A value between 0 and 1 indicating position along the bound element's perimeter
	Gap       float64 `json:"gap"`       // Distance from the bound element's perimeter to the linear element's start/end point
}
