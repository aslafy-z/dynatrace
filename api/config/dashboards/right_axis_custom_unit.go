package dashboards

import "encoding/json"

// RightAxisCustomUnit has no documentation
type RightAxisCustomUnit string

// RightAxisCustomUnits has no documentation
var RightAxisCustomUnits = struct {
	Bit                  RightAxisCustomUnit
	Bitperhour           RightAxisCustomUnit
	Bitperminute         RightAxisCustomUnit
	Bitpersecond         RightAxisCustomUnit
	Byte                 RightAxisCustomUnit
	Byteperhour          RightAxisCustomUnit
	Byteperminute        RightAxisCustomUnit
	Bytepersecond        RightAxisCustomUnit
	Cores                RightAxisCustomUnit
	Count                RightAxisCustomUnit
	Day                  RightAxisCustomUnit
	Decibelmilliwatt     RightAxisCustomUnit
	Gibibyte             RightAxisCustomUnit
	Giga                 RightAxisCustomUnit
	Gigabyte             RightAxisCustomUnit
	Hour                 RightAxisCustomUnit
	Kibibyte             RightAxisCustomUnit
	Kibibyteperhour      RightAxisCustomUnit
	Kibibyteperminute    RightAxisCustomUnit
	Kibibytepersecond    RightAxisCustomUnit
	Kilo                 RightAxisCustomUnit
	Kilobyte             RightAxisCustomUnit
	Kilobyteperhour      RightAxisCustomUnit
	Kilobyteperminute    RightAxisCustomUnit
	Kilobytepersecond    RightAxisCustomUnit
	Mebibyte             RightAxisCustomUnit
	Mebibyteperhour      RightAxisCustomUnit
	Mebibyteperminute    RightAxisCustomUnit
	Mebibytepersecond    RightAxisCustomUnit
	Mega                 RightAxisCustomUnit
	Megabyte             RightAxisCustomUnit
	Megabyteperhour      RightAxisCustomUnit
	Megabyteperminute    RightAxisCustomUnit
	Megabytepersecond    RightAxisCustomUnit
	Microsecond          RightAxisCustomUnit
	Millicores           RightAxisCustomUnit
	Millisecond          RightAxisCustomUnit
	Millisecondperminute RightAxisCustomUnit
	Minute               RightAxisCustomUnit
	Month                RightAxisCustomUnit
	Nanosecond           RightAxisCustomUnit
	Nanosecondperminute  RightAxisCustomUnit
	Notapplicable        RightAxisCustomUnit
	Perhour              RightAxisCustomUnit
	Perminute            RightAxisCustomUnit
	Persecond            RightAxisCustomUnit
	Percent              RightAxisCustomUnit
	Pixel                RightAxisCustomUnit
	Promille             RightAxisCustomUnit
	Ratio                RightAxisCustomUnit
	Second               RightAxisCustomUnit
	State                RightAxisCustomUnit
	Unspecified          RightAxisCustomUnit
	Week                 RightAxisCustomUnit
	Year                 RightAxisCustomUnit
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
func (t *RightAxisCustomUnit) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = RightAxisCustomUnit(name)
	return nil
}
