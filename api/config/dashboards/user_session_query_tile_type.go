package dashboards

import "encoding/json"

// UserSessionQueryTileType has no documentation
type UserSessionQueryTileType string

// UserSessionQueryTileTypes offers the known enum values
var UserSessionQueryTileTypes = struct {
	ColumnChart UserSessionQueryTileType
	Funnel      UserSessionQueryTileType
	LineChart   UserSessionQueryTileType
	PieChart    UserSessionQueryTileType
	SingleValue UserSessionQueryTileType
	Table       UserSessionQueryTileType
}{
	"COLUMN_CHART",
	"FUNNEL",
	"LINE_CHART",
	"PIE_CHART",
	"SINGLE_VALUE",
	"TABLE",
}

// UnmarshalJSON performs custom unmarshalling of this enum type
func (t *UserSessionQueryTileType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = UserSessionQueryTileType(name)
	return nil
}
