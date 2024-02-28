package vtu_ng

// NetworkID represents the network IDs for MTN, Glo, Airtel, and Etisalat.
type NetworkID string
type DataVariantID string

// Constants representing the network IDs.
const (
	Mtn      NetworkID = "mtn"
	Glo      NetworkID = "glo"
	Airtel   NetworkID = "airtel"
	Etisalat NetworkID = "etisalat"
)

// Constants representing the data plan variant IDs for MTN.
const (
	MtnSMEData500MB   DataVariantID = "500"
	MtnSMEData1GB     DataVariantID = "M1024"
	MtnSMEData2GB     DataVariantID = "M2024"
	MtnSMEData3GB     DataVariantID = "3000"
	MtnSMEData5GB     DataVariantID = "5000"
	MtnSMEData10GB    DataVariantID = "10000"
	MtnData6GB7Days   DataVariantID = "mtn-20hrs-1500"
	MtnData30GB30Days DataVariantID = "mtn-30gb-8000"
	MtnData40GB30Days DataVariantID = "mtn-40gb-10000"
	MtnData75GB30Days DataVariantID = "mtn-75gb-15000"
)

// Constants representing the data plan variant IDs for Glo.
const (
	GloData1GB5Nights    DataVariantID = "glo100x"
	GloData1_25GB1Day    DataVariantID = "glo200x"
	GloData1_35GB14Days  DataVariantID = "G500"
	GloData5_8GB30Days   DataVariantID = "G2000"
	GloData2_9GB30Days   DataVariantID = "G1000"
	GloData7_7GB30Days   DataVariantID = "G2500"
	GloData10GB30Days    DataVariantID = "G3000"
	GloData13_25GB30Days DataVariantID = "G4000"
	GloData18_25GB30Days DataVariantID = "G5000"
	GloData29_5GB30Days  DataVariantID = "G8000"
	GloData50GB30Days    DataVariantID = "glo10000"
)

// Constants representing the data plan variant IDs for Airtel.
const (
	AirtelData500MB30Days    DataVariantID = "AIRTEL500MB"
	AirtelData1GB30GiftDays  DataVariantID = "AIRTEL1GB"
	AirtelData2GB30GiftDays  DataVariantID = "AIRTEL2GB"
	AirtelData5GB30GiftDays  DataVariantID = "AIRTEL5GB"
	AirtelData10GB30GiftDays DataVariantID = "AIRTEL10GB"
	AirtelData15GB30GiftDays DataVariantID = "AIRTEL15GB"
	AirtelData20GB30GiftDays DataVariantID = "AIRTEL20GB"
	AirtelData1_5GB30Days    DataVariantID = "airt-1100"
	AirtelData2GB30Days      DataVariantID = "airt-1300"
	AirtelData3GB30Days      DataVariantID = "airt-1650"
	AirtelData4_5GB30Days    DataVariantID = "airt-2200"
	AirtelData10GB30Days     DataVariantID = "airt-3300"
	AirtelData20GB30Days     DataVariantID = "airt-5500"
	AirtelData40GB30Days     DataVariantID = "airt-11000"
	AirtelData1GB1Day        DataVariantID = "airt-330x"
	AirtelData750MB14Days    DataVariantID = "airt-550"
	AirtelData6GB7Days       DataVariantID = "airt-1650-2"
)

type DataPlanInfoType []DataPlanInfo

var NineMobileDataDescriptions DataPlanInfoType
var AirtelDataDescriptions DataPlanInfoType
var GloDataDescriptions DataPlanInfoType
var MtnDataDescriptions DataPlanInfoType

// Constants representing the data plan variant IDs for 9mobile.
const (
	NineMobileData1GB30Days    DataVariantID = "9MOB1000"
	NineMobileData2_5GB30Days  DataVariantID = "9MOB34500"
	NineMobileData11_5GB30Days DataVariantID = "9MOB8000"
	NineMobileData15GB30Days   DataVariantID = "9MOB5000"
)

type DataPlanInfo struct {
	Description string        `json:"description"`
	Price       int64         `json:"price"`
	VariationID DataVariantID `json:"variation_id"`
}

func init() {
	// Initialize the array of maps with descriptions
	NineMobileDataDescriptions = []DataPlanInfo{
		{
			Description: "9mobile Data 1GB - 30 Days",
			Price:       989,
			VariationID: NineMobileData1GB30Days,
		},
		{
			Description: "9mobile Data 2.5GB - 30 Days",
			Price:       1989,
			VariationID: NineMobileData2_5GB30Days,
		},
		{
			Description: "9mobile Data 11.5GB - 30 Days",
			Price:       7969,
			VariationID: NineMobileData11_5GB30Days,
		},
		{
			Description: "9mobile Data 15GB - 30 Days",
			Price:       9899,
			VariationID: NineMobileData15GB30Days,
		},
	}

	AirtelDataDescriptions = []DataPlanInfo{
		{
			Description: "Airtel Data 500MB (Gift) – 30 Days",
			Price:       159,
			VariationID: AirtelData500MB30Days,
		},
		{
			Description: "Airtel Data 1GB (Gift) – 30 Days",
			Price:       299,
			VariationID: AirtelData1GB30GiftDays,
		},
		{
			Description: "Airtel Data 2GB (Gift)– 30 Days",
			Price:       599,
			VariationID: AirtelData2GB30GiftDays,
		},
		{
			Description: "Airtel Data 5GB (Gift)– 30 Days",
			Price:       1499,
			VariationID: AirtelData5GB30GiftDays,
		},
		{
			Description: "Airtel Data 10GB (Gift)– 30 Days",
			Price:       2899,
			VariationID: AirtelData10GB30GiftDays,
		}, {
			Description: "Airtel Data 15GB (Gift)– 30 Days",
			Price:       4339,
			VariationID: AirtelData15GB30GiftDays,
		}, {
			Description: "Airtel Data 20GB (Gift)– 30 Days",
			Price:       5779,
			VariationID: AirtelData20GB30GiftDays,
		}, {
			Description: "Airtel Data 1.5GB – 30 Days",
			Price:       1079,
			VariationID: AirtelData1_5GB30Days,
		}, {
			Description: "Airtel Data 2GB – 30 Days",
			Price:       1289,
			VariationID: AirtelData2GB30Days,
		}, {
			Description: "Airtel Data 3GB – 30 Days",
			Price:       1639,
			VariationID: AirtelData3GB30Days,
		}, {
			Description: "Airtel Data 4.5GB – 30 Days",
			Price:       2189,
			VariationID: AirtelData4_5GB30Days,
		},
		{
			Description: "Airtel Data 10GB – 30 Days",
			Price:       3289,
			VariationID: AirtelData10GB30Days,
		}, {
			Description: "Airtel Data 20GB – 30 Days",
			Price:       5489,
			VariationID: AirtelData20GB30Days,
		},
		{
			Description: "Airtel Data 40GB – 30 Days",
			Price:       10799,
			VariationID: AirtelData40GB30Days,
		}, {
			Description: "Airtel Data 1GB – 1 Day",
			Price:       329,
			VariationID: AirtelData1GB1Day,
		},
		{
			Description: "Airtel Data 750MB – 14 Days",
			Price:       545,
			VariationID: AirtelData750MB14Days,
		},
		{
			Description: "Airtel Data 6GB – 7 Days",
			Price:       1639,
			VariationID: AirtelData6GB7Days,
		},
	}

	GloDataDescriptions = []DataPlanInfo{
		{
			Description: "Glo Data 1GB – 5 Nights",
			Price:       99,
			VariationID: GloData1GB5Nights,
		},
		{
			Description: "Glo Data 1.25GB – 1 Day (Sunday)",
			Price:       199,
			VariationID: GloData1_25GB1Day,
		},
		{
			Description: "Glo Data 1.35GB – 14 Days",
			Price:       489,
			VariationID: GloData1_35GB14Days,
		},
		{
			Description: "Glo Data 5.8GB – 30 Days",
			Price:       1949,
			VariationID: GloData5_8GB30Days,
		},
		{
			Description: "Glo Data 2.9GB – 30 Days",
			Price:       979,
			VariationID: GloData7_7GB30Days,
		},
		{
			Description: "Glo Data 10GB – 30 Days",
			Price:       2949,
			VariationID: GloData10GB30Days,
		},
		{
			Description: "Glo Data 13.25GB – 30 Days",
			Price:       3889,
			VariationID: GloData13_25GB30Days,
		}, {
			Description: "Glo Data 18.25GB – 30 Days",
			Price:       4849,
			VariationID: GloData18_25GB30Days,
		}, {
			Description: "Glo Data 29.5GB – 30 Days",
			Price:       7799,
			VariationID: GloData29_5GB30Days,
		},
		{
			Description: "Glo Data 50GB – 30 Days",
			Price:       9899,
			VariationID: GloData50GB30Days,
		},
	}

	MtnDataDescriptions = []DataPlanInfo{
		{
			Description: "MTN SME Data 500MB – 30 Days",
			Price:       199,
			VariationID: MtnSMEData500MB,
		},
		{
			Description: "MTN SME Data 1GB – 30 Days",
			Price:       339,
			VariationID: MtnSMEData1GB,
		},
		{
			Description: "MTN SME Data 2GB – 30 Days",
			Price:       679,
			VariationID: MtnSMEData2GB,
		},
		{
			Description: "MTN SME Data 3GB – 30 Days",
			Price:       1019,
			VariationID: MtnSMEData3GB,
		},
		{
			Description: "MTN Data 5GB – 30 Days",
			Price:       1699,
			VariationID: MtnSMEData5GB,
		},
		{
			Description: "MTN Data 10GB – 30 Days",
			Price:       3399,
			VariationID: MtnSMEData10GB,
		},
		{
			Description: "MTN Data 6GB – 7 Days",
			Price:       1499,
			VariationID: MtnData6GB7Days,
		},
		{
			Description: "MTN Data 30GB – 30 Days",
			Price:       7959,
			VariationID: MtnData30GB30Days,
		},
		{
			Description: "MTN Data 40GB – 30 Days",
			Price:       9899,
			VariationID: MtnData40GB30Days,
		},
		{
			Description: "MTN Data 75GB – 30 Days",
			Price:       14979,
			VariationID: MtnData75GB30Days,
		},
	}

}
