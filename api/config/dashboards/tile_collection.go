package dashboards

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var mapping = map[TileType]Tile{
	TileTypes.CustomCharting:          new(CustomChartingTile),
	TileTypes.Markdown:                new(MarkdownTile),
	TileTypes.DTAQL:                   new(UserSessionQueryTile),
	TileTypes.Hosts:                   new(FilterableEntityTile),
	TileTypes.Applications:            new(FilterableEntityTile),
	TileTypes.Services:                new(FilterableEntityTile),
	TileTypes.DatabasesOverview:       new(FilterableEntityTile),
	TileTypes.SyntheticTests:          new(FilterableEntityTile),
	TileTypes.ApplicationWorldMap:     new(AssignedEntitiesWithMetricTile),
	TileTypes.Resources:               new(AssignedEntitiesWithMetricTile),
	TileTypes.ThirdPartyMostActive:    new(AssignedEntitiesWithMetricTile),
	TileTypes.UEMConversionsPerGoal:   new(AssignedEntitiesWithMetricTile),
	TileTypes.Host:                    new(AssignedEntitiesWithMetricTile),
	TileTypes.ProcessGroupsOne:        new(AssignedEntitiesWithMetricTile),
	TileTypes.SyntheticSingleWebCheck: new(SyntheticSingleWebCheckTile),
	TileTypes.Application:             new(AssignedEntitiesTile),
	TileTypes.Virtualization:          new(AssignedEntitiesTile),
	TileTypes.AWS:                     new(AssignedEntitiesTile),
	TileTypes.ServiceVersatile:        new(AssignedEntitiesTile),
	TileTypes.SessionMetrics:          new(AssignedEntitiesTile),
	TileTypes.Users:                   new(AssignedEntitiesTile),
	TileTypes.UEMKeyUserActions:       new(AssignedEntitiesTile),
	TileTypes.BounceRate:              new(AssignedEntitiesTile),
	TileTypes.UEMConversionsOverall:   new(AssignedEntitiesTile),
	TileTypes.UEMJserrorsOverall:      new(AssignedEntitiesTile),
	TileTypes.MobileApplication:       new(AssignedEntitiesTile),
	TileTypes.SyntheticSingleExtTest:  new(AssignedEntitiesTile),
	TileTypes.SyntheticHTTPMonitor:    new(AssignedEntitiesTile),
	TileTypes.Database:                new(AssignedEntitiesTile),
	TileTypes.CustomApplication:       new(AssignedEntitiesTile),
	TileTypes.ApplicationMethod:       new(AssignedEntitiesTile),
	TileTypes.LogAnalytics:            new(AssignedEntitiesTile),
	TileTypes.OpenStack:               new(AssignedEntitiesTile),
	TileTypes.OpenStackProject:        new(AssignedEntitiesTile),
	TileTypes.OpenStackAVZone:         new(AssignedEntitiesTile),
	TileTypes.DeviceApplicationMethod: new(AssignedEntitiesTile),
	TileTypes.DemKeyUserAction:        new(AssignedEntitiesTile),
	TileTypes.Header:                  new(AbstractTile),
}

// TileCollection has no documentation
type TileCollection []Tile

// Append adds additional tiles to this collection of Tiles
func (collection *TileCollection) Append(tiles ...Tile) {
	for _, tile := range tiles {
		if tile.GetType() == "" {
			panic("Trying to add a Tile without its Type property set")
		}
		*collection = append(*collection, tile)
	}
}

// GetTiles has no documentation
func (collection *TileCollection) GetTiles(t TileType, a interface{}) {
	switch t {
	case TileTypes.CustomCharting:
		if b, ok := a.(*[]*CustomChartingTile); !ok {
			panic(fmt.Sprintf("You need to provide a *[]%T when asking for Tiles of type %v", mapping[TileTypes.CustomCharting], TileTypes.CustomCharting))
		} else {
			for _, tile := range *collection {
				if tile.GetType() == TileTypes.CustomCharting {
					*b = append(*b, tile.(*CustomChartingTile))
				}
			}
		}
	}
}

// UnmarshalJSON TODO: documentation
func (collection *TileCollection) UnmarshalJSON(data []byte) error {
	rawValues := make([]json.RawMessage, 0)
	if e := json.Unmarshal(data, &rawValues); e != nil {
		return e
	}
	for _, rawValue := range rawValues {
		untyped := AbstractTile{}
		if e := json.Unmarshal(rawValue, &untyped); e != nil {
			return e
		}
		typed := newTyped(untyped.TileType)
		if e := json.Unmarshal(rawValue, &typed); e != nil {
			return e
		}
		typed.SetType(untyped.TileType)
		*collection = append(*collection, typed)
	}
	return nil
}

func newTyped(t TileType) Tile {
	if p := mapping[t]; p != nil {
		return reflect.New(reflect.TypeOf(p).Elem()).Interface().(Tile)
	}
	return new(AbstractTile)
}
