package dashboards

import (
	"encoding/json"
)

// Tile the configuration of a tile.
// The actual set of fields depends on the type of the tile. See the description of the **tileType** field
type Tile interface {
	GetName() string
	GetType() TileType
	SetType(TileType)
	IsConfigured() bool
	GetBounds() *TileBounds
	GetFilter() *TileFilter
	AsCustomChartingTile() *CustomChartingTile
	AsMarkdownTile() *MarkdownTile
	AsUserSessionQueryTile() *UserSessionQueryTile
	AsFilterableEntityTile() *FilterableEntityTile
	AsAssignedEntitiesWithMetricTile() *AssignedEntitiesWithMetricTile
	AsAssignedEntitiesTile() *AssignedEntitiesTile
	AsSyntheticSingleWebCheckTile() *SyntheticSingleWebCheckTile
	AsGenericTile() *GenericTile
	Initialize(*AbstractTile)
	Implementors() []Tile
}

// AbstractTile the configuration of a tile.
// The actual set of fields depends on the type of the tile. See the description of the **tileType** field
type AbstractTile struct {
	Name       string                 `json:"name"` // the name of the tile
	TileType   TileType               `json:"tileType" json-x:"discriminator"`
	Configured *bool                  `json:"configured,omitempty"` // The tile is configured and ready to use (`true`) or just placed on the dashboard (`false`)
	Bounds     *TileBounds            `json:"bounds"`               // Bounds the position and size of a tile
	Filter     *TileFilter            `json:"tileFilter,omitempty"` // is filter applied to a tile. It overrides dashboard's filter
	Properties map[string]interface{} `json:"-" json-x:"*"`
}

// Initialize has no documentation
func (at *AbstractTile) Initialize(other *AbstractTile) {
	at.Name = other.Name
	at.TileType = other.TileType
	at.Configured = other.Configured
	at.Bounds = other.Bounds
	at.Filter = other.Filter
}

// Implementors has no documentation
func (at *AbstractTile) Implementors() []Tile {
	return []Tile{
		new(CustomChartingTile),
		new(MarkdownTile),
		new(UserSessionQueryTile),
		new(FilterableEntityTile),
		new(AssignedEntitiesWithMetricTile),
		new(SyntheticSingleWebCheckTile),
		new(AssignedEntitiesTile),
		new(AbstractTile),
	}
}

// GetName has no documentation
func (at *AbstractTile) GetName() string {
	return at.Name
}

// GetType has no documentation
func (at *AbstractTile) GetType() TileType {
	return at.TileType
}

// SetType has no documentation
func (at *AbstractTile) SetType(t TileType) {
	at.TileType = t
}

// IsConfigured has no documentation
func (at *AbstractTile) IsConfigured() bool {
	return *at.Configured
}

// GetBounds has no documentation
func (at *AbstractTile) GetBounds() *TileBounds {
	return at.Bounds
}

// GetFilter has no documentation
func (at *AbstractTile) GetFilter() *TileFilter {
	return at.Filter
}

// AsCustomChartingTile has not documentation
func (at *AbstractTile) AsCustomChartingTile() *CustomChartingTile {
	return nil
}

// AsMarkdownTile has not documentation
func (at *AbstractTile) AsMarkdownTile() *MarkdownTile {
	return nil
}

// AsUserSessionQueryTile has not documentation
func (at *AbstractTile) AsUserSessionQueryTile() *UserSessionQueryTile {
	return nil
}

// AsFilterableEntityTile has not documentation
func (at *AbstractTile) AsFilterableEntityTile() *FilterableEntityTile {
	return nil
}

// AsAssignedEntitiesWithMetricTile has not documentation
func (at *AbstractTile) AsAssignedEntitiesWithMetricTile() *AssignedEntitiesWithMetricTile {
	return nil
}

// AsSyntheticSingleWebCheckTile has not documentation
func (at *AbstractTile) AsSyntheticSingleWebCheckTile() *SyntheticSingleWebCheckTile {
	return nil
}

// AsAssignedEntitiesTile has not documentation
func (at *AbstractTile) AsAssignedEntitiesTile() *AssignedEntitiesTile {
	return nil
}

// AsGenericTile has not documentation
func (at *AbstractTile) AsGenericTile() *GenericTile {
	return nil
}

// GenericTile has no documentation
type GenericTile struct {
	AbstractTile `json-x:"*"`
	Properties   map[string]interface{} `json:"-"`
}

// MarshalJSON TODO: documentation
func (gt *GenericTile) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, len(gt.Properties))
	m["name"] = gt.Name
	m["tileType"] = gt.TileType
	m["configured"] = gt.Configured
	m["bounds"] = gt.Bounds
	m["tileFilter"] = gt.Filter
	for k, v := range gt.Properties {
		m[k] = v
	}
	return json.Marshal(m)
}

// UnmarshalJSON TODO: documentation
func (gt *GenericTile) UnmarshalJSON(data []byte) error {
	var err error
	rawProperties := make(map[string]json.RawMessage)
	if err = json.Unmarshal(data, &rawProperties); err != nil {
		return err
	}
	if val, ok := rawProperties["name"]; ok {
		if err = json.Unmarshal(val, &gt.Name); err != nil {
			return err
		}
	}
	if val, ok := rawProperties["tileType"]; ok {
		if err = json.Unmarshal(val, &gt.TileType); err != nil {
			return err
		}
	}
	if val, ok := rawProperties["configured"]; ok {
		if err = json.Unmarshal(val, &gt.Configured); err != nil {
			return err
		}
	}
	if val, ok := rawProperties["bounds"]; ok {
		gt.Bounds = new(TileBounds)
		if err = json.Unmarshal(val, gt.Bounds); err != nil {
			return err
		}
	}
	if val, ok := rawProperties["tileFilter"]; ok {
		gt.Filter = new(TileFilter)
		if err = json.Unmarshal(val, gt.Filter); err != nil {
			return err
		}
	}
	gt.Properties = make(map[string]interface{})
	if err = json.Unmarshal(data, &gt.Properties); err != nil {
		delete(gt.Properties, "name")
		delete(gt.Properties, "tileType")
		delete(gt.Properties, "configured")
		delete(gt.Properties, "bounds")
		delete(gt.Properties, "tileFilter")
	}

	return nil
}

// AsCustomChartingTile has not documentation
func (gt *GenericTile) AsCustomChartingTile() *CustomChartingTile {
	return nil
}

// AsMarkdownTile has not documentation
func (gt *GenericTile) AsMarkdownTile() *MarkdownTile {
	return nil
}

// AsUserSessionQueryTile has not documentation
func (gt *GenericTile) AsUserSessionQueryTile() *UserSessionQueryTile {
	return nil
}

// AsFilterableEntityTile has not documentation
func (gt *GenericTile) AsFilterableEntityTile() *FilterableEntityTile {
	return nil
}

// AsAssignedEntitiesWithMetricTile has not documentation
func (gt *GenericTile) AsAssignedEntitiesWithMetricTile() *AssignedEntitiesWithMetricTile {
	return nil
}

// AsSyntheticSingleWebCheckTile has not documentation
func (gt *GenericTile) AsSyntheticSingleWebCheckTile() *SyntheticSingleWebCheckTile {
	return nil
}

// AsAssignedEntitiesTile has not documentation
func (gt *GenericTile) AsAssignedEntitiesTile() *AssignedEntitiesTile {
	return nil
}

// AsGenericTile has not documentation
func (gt *GenericTile) AsGenericTile() *GenericTile {
	return gt
}
