package autotags

// DCRumDecoderComparisonValue The value to compare to.
type DCRumDecoderComparisonValue string

// DCRumDecoderComparisonValues offers the known enum values
var DCRumDecoderComparisonValues = struct {
	AllOther         DCRumDecoderComparisonValue
	CitrixAppFlow    DCRumDecoderComparisonValue
	CitrixIca        DCRumDecoderComparisonValue
	CitrixIcaOverSSL DCRumDecoderComparisonValue
	DB2Drda          DCRumDecoderComparisonValue
	HTTP             DCRumDecoderComparisonValue
	HTTPS            DCRumDecoderComparisonValue
	HTTPExpress      DCRumDecoderComparisonValue
	Informix         DCRumDecoderComparisonValue
	MySQL            DCRumDecoderComparisonValue
	Oracle           DCRumDecoderComparisonValue
	SAPGUI           DCRumDecoderComparisonValue
	SAPGUIOverHTTP   DCRumDecoderComparisonValue
	SAPGUIOverHTTPS  DCRumDecoderComparisonValue
	SAPHanaDB        DCRumDecoderComparisonValue
	SAPRfc           DCRumDecoderComparisonValue
	SSL              DCRumDecoderComparisonValue
	TDS              DCRumDecoderComparisonValue
}{
	"ALL_OTHER",
	"CITRIX_APPFLOW",
	"CITRIX_ICA",
	"CITRIX_ICA_OVER_SSL",
	"DB2_DRDA",
	"HTTP",
	"HTTPS",
	"HTTP_EXPRESS",
	"INFORMIX",
	"MYSQL",
	"ORACLE",
	"SAP_GUI",
	"SAP_GUI_OVER_HTTP",
	"SAP_GUI_OVER_HTTPS",
	"SAP_HANA_DB",
	"SAP_RFC",
	"SSL",
	"TDS",
}
