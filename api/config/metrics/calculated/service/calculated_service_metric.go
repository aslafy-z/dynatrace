package service

import api "github.com/dtcookie/dynatrace/api/config"

// CalculatedServiceMetric Descriptor of a calculated service metric.
type CalculatedServiceMetric struct {
	EntityID            *string                    `json:"entityId,omitempty"`            // Restricts the metric usage to the specified service.   This field is mutually exclusive with the **managementZones** field.
	Name                string                     `json:"name"`                          // The displayed name of the metric.
	Unit                Unit                       `json:"unit"`                          // The unit of the metric.
	DimensionDefinition *DimensionDefinition       `json:"dimensionDefinition,omitempty"` // Parameters of a definition of a calculated service metric.
	Enabled             bool                       `json:"enabled"`                       // The metric is enabled (`true`) or disabled (`false`).
	ManagementZones     []string                   `json:"managementZones,omitempty"`     // Restricts the metric usage to specified management zones.   This field is mutually exclusive with the **entityId** field.
	MetricDefinition    CalculatedMetricDefinition `json:"metricDefinition"`              // The definition of a calculated service metric.
	TsmMetricKey        string                     `json:"tsmMetricKey"`                  // The key of the calculated service metric.
	UnitDisplayName     *string                    `json:"unitDisplayName,omitempty"`     // The display name of the metric's unit.   Only applicable when the **unit** parameter is set to `UNSPECIFIED`.
	Conditions          []*Condition               `json:"conditions,omitempty"`          // The set of conditions for the metric usage.   **All** the specified conditions must be fulfilled to use the metric.
	Metadata            *api.ConfigMetadata        `json:"metadata,omitempty"`            // Metadata useful for debugging
}

// Unit The unit of the metric.
type Unit string

// Units offers the known enum values
var Units = struct {
	Bit                  Unit
	BitPerHour           Unit
	BitPerMinute         Unit
	BitPerSecond         Unit
	Byte                 Unit
	BytePerHour          Unit
	BytePerMinute        Unit
	BytePerSecond        Unit
	Cores                Unit
	Count                Unit
	Day                  Unit
	DecibelMilliWatt     Unit
	GibiByte             Unit
	Giga                 Unit
	GigaByte             Unit
	Hour                 Unit
	KibiByte             Unit
	KibiBytePerHour      Unit
	KibiBytePerMinute    Unit
	KibiBytePerSecond    Unit
	Kilo                 Unit
	KiloByte             Unit
	KiloBytePerHour      Unit
	KiloBytePerMinute    Unit
	KiloBytePerSecond    Unit
	MebiByte             Unit
	MebiBytePerHour      Unit
	MebiBytePerMinute    Unit
	MebiBytePerSecond    Unit
	Mega                 Unit
	MegaByte             Unit
	MegaBytePerHour      Unit
	MegaBytePerMinute    Unit
	MegaBytePerSecond    Unit
	MicroSecond          Unit
	MilliCores           Unit
	MilliSecond          Unit
	MilliSecondPerMinute Unit
	Minute               Unit
	Month                Unit
	Msu                  Unit
	NanoSecond           Unit
	NanoSecondPerMinute  Unit
	NotApplicable        Unit
	Percent              Unit
	PerHour              Unit
	PerMinute            Unit
	PerSecond            Unit
	Pixel                Unit
	Promille             Unit
	Ratio                Unit
	Second               Unit
	State                Unit
	Unspecified          Unit
	Week                 Unit
	Year                 Unit
}{
	"BIT",
	"BIT_PER_HOUR",
	"BIT_PER_MINUTE",
	"BIT_PER_SECOND",
	"BYTE",
	"BYTE_PER_HOUR",
	"BYTE_PER_MINUTE",
	"BYTE_PER_SECOND",
	"CORES",
	"COUNT",
	"DAY",
	"DECIBEL_MILLI_WATT",
	"GIBI_BYTE",
	"GIGA",
	"GIGA_BYTE",
	"HOUR",
	"KIBI_BYTE",
	"KIBI_BYTE_PER_HOUR",
	"KIBI_BYTE_PER_MINUTE",
	"KIBI_BYTE_PER_SECOND",
	"KILO",
	"KILO_BYTE",
	"KILO_BYTE_PER_HOUR",
	"KILO_BYTE_PER_MINUTE",
	"KILO_BYTE_PER_SECOND",
	"MEBI_BYTE",
	"MEBI_BYTE_PER_HOUR",
	"MEBI_BYTE_PER_MINUTE",
	"MEBI_BYTE_PER_SECOND",
	"MEGA",
	"MEGA_BYTE",
	"MEGA_BYTE_PER_HOUR",
	"MEGA_BYTE_PER_MINUTE",
	"MEGA_BYTE_PER_SECOND",
	"MICRO_SECOND",
	"MILLI_CORES",
	"MILLI_SECOND",
	"MILLI_SECOND_PER_MINUTE",
	"MINUTE",
	"MONTH",
	"MSU",
	"NANO_SECOND",
	"NANO_SECOND_PER_MINUTE",
	"NOT_APPLICABLE",
	"PERCENT",
	"PER_HOUR",
	"PER_MINUTE",
	"PER_SECOND",
	"PIXEL",
	"PROMILLE",
	"RATIO",
	"SECOND",
	"STATE",
	"UNSPECIFIED",
	"WEEK",
	"YEAR",
}
