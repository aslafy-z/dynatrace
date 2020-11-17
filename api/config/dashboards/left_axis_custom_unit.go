package dashboards

import "encoding/json"

// LeftAxisCustomUnit has no documentation
type LeftAxisCustomUnit string

// LeftAxisCustomUnits has no documentation
var LeftAxisCustomUnits = struct {
	Bit                  LeftAxisCustomUnit
	Bitperhour           LeftAxisCustomUnit
	Bitperminute         LeftAxisCustomUnit
	Bitpersecond         LeftAxisCustomUnit
	Byte                 LeftAxisCustomUnit
	Byteperhour          LeftAxisCustomUnit
	Byteperminute        LeftAxisCustomUnit
	Bytepersecond        LeftAxisCustomUnit
	Cores                LeftAxisCustomUnit
	Count                LeftAxisCustomUnit
	Day                  LeftAxisCustomUnit
	Decibelmilliwatt     LeftAxisCustomUnit
	Gibibyte             LeftAxisCustomUnit
	Giga                 LeftAxisCustomUnit
	Gigabyte             LeftAxisCustomUnit
	Hour                 LeftAxisCustomUnit
	Kibibyte             LeftAxisCustomUnit
	Kibibyteperhour      LeftAxisCustomUnit
	Kibibyteperminute    LeftAxisCustomUnit
	Kibibytepersecond    LeftAxisCustomUnit
	Kilo                 LeftAxisCustomUnit
	Kilobyte             LeftAxisCustomUnit
	Kilobyteperhour      LeftAxisCustomUnit
	Kilobyteperminute    LeftAxisCustomUnit
	Kilobytepersecond    LeftAxisCustomUnit
	Mebibyte             LeftAxisCustomUnit
	Mebibyteperhour      LeftAxisCustomUnit
	Mebibyteperminute    LeftAxisCustomUnit
	Mebibytepersecond    LeftAxisCustomUnit
	Mega                 LeftAxisCustomUnit
	Megabyte             LeftAxisCustomUnit
	Megabyteperhour      LeftAxisCustomUnit
	Megabyteperminute    LeftAxisCustomUnit
	Megabytepersecond    LeftAxisCustomUnit
	Microsecond          LeftAxisCustomUnit
	Millicores           LeftAxisCustomUnit
	Millisecond          LeftAxisCustomUnit
	Millisecondperminute LeftAxisCustomUnit
	Minute               LeftAxisCustomUnit
	Month                LeftAxisCustomUnit
	Nanosecond           LeftAxisCustomUnit
	Nanosecondperminute  LeftAxisCustomUnit
	Notapplicable        LeftAxisCustomUnit
	Perhour              LeftAxisCustomUnit
	Perminute            LeftAxisCustomUnit
	Persecond            LeftAxisCustomUnit
	Percent              LeftAxisCustomUnit
	Pixel                LeftAxisCustomUnit
	Promille             LeftAxisCustomUnit
	Ratio                LeftAxisCustomUnit
	Second               LeftAxisCustomUnit
	State                LeftAxisCustomUnit
	Unspecified          LeftAxisCustomUnit
	Week                 LeftAxisCustomUnit
	Year                 LeftAxisCustomUnit
}{
	"Bit",
	"BitPerHour",
	"BitPerMinute",
	"BitPerSecond",
	"Byte",
	"BytePerHour",
	"BytePerMinute",
	"BytePerSecond",
	"Cores",
	"Count",
	"Day",
	"DecibelMilliWatt",
	"GibiByte",
	"Giga",
	"GigaByte",
	"Hour",
	"KibiByte",
	"KibiBytePerHour",
	"KibiBytePerMinute",
	"KibiBytePerSecond",
	"Kilo",
	"KiloByte",
	"KiloBytePerHour",
	"KiloBytePerMinute",
	"KiloBytePerSecond",
	"MebiByte",
	"MebiBytePerHour",
	"MebiBytePerMinute",
	"MebiBytePerSecond",
	"Mega",
	"MegaByte",
	"MegaBytePerHour",
	"MegaBytePerMinute",
	"MegaBytePerSecond",
	"MicroSecond",
	"MilliCores",
	"MilliSecond",
	"MilliSecondPerMinute",
	"Minute",
	"Month",
	"NanoSecond",
	"NanoSecondPerMinute",
	"NotApplicable",
	"PerHour",
	"PerMinute",
	"PerSecond",
	"Percent",
	"Pixel",
	"Promille",
	"Ratio",
	"Second",
	"State",
	"Unspecified",
	"Week",
	"Year",
}

// UnmarshalJSON performs custom unmarshalling of this enum type
func (t *LeftAxisCustomUnit) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = LeftAxisCustomUnit(name)
	return nil
}
