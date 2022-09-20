package sap_api_output_formatter

type PlannedOrder struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	PlannedOrder  string `json:"planned_order"`
	Deleted       bool   `json:"deleted"`
}

type Header struct {
	PlannedOrder                   string `json:"PlannedOrder"`
	PlannedOrderType               string `json:"PlannedOrderType"`
	PlannedOrderProfile            string `json:"PlannedOrderProfile"`
	Material                       string `json:"Material"`
	MaterialName                   string `json:"MaterialName"`
	ProductionPlant                string `json:"ProductionPlant"`
	MRPPlant                       string `json:"MRPPlant"`
	MRPArea                        string `json:"MRPArea"`
	ProductionVersion              string `json:"ProductionVersion"`
	MaterialProcurementCategory    string `json:"MaterialProcurementCategory"`
	MaterialProcurementType        string `json:"MaterialProcurementType"`
	StorageLocation                string `json:"StorageLocation"`
	BaseUnit                       string `json:"BaseUnit"`
	TotalQuantity                  int    `json:"TotalQuantity"`
	PlndOrderPlannedScrapQty       int    `json:"PlndOrderPlannedScrapQty"`
	GoodsReceiptQty                int    `json:"GoodsReceiptQty"`
	IssuedQuantity                 int    `json:"IssuedQuantity"`
	PlndOrderPlannedStartDate      string `json:"PlndOrderPlannedStartDate"`
	PlndOrderPlannedStartTime      string `json:"PlndOrderPlannedStartTime"`
	PlndOrderPlannedEndDate        string `json:"PlndOrderPlannedEndDate"`
	PlndOrderPlannedEndTime        string `json:"PlndOrderPlannedEndTime"`
	PlannedOrderOpeningDate        string `json:"PlannedOrderOpeningDate"`
	PlannedOrderLastChangeDateTime string `json:"PlannedOrderLastChangeDateTime"`
	ProductionStartDate            string `json:"ProductionStartDate"`
	ProductionEndDate              string `json:"ProductionEndDate"`
	SalesOrder                     string `json:"SalesOrder"`
	SalesOrderItem                 string `json:"SalesOrderItem"`
	Customer                       string `json:"Customer"`
	WBSElementInternalID           string `json:"WBSElementInternalID"`
	WBSElementExternalID           string `json:"WBSElementExternalID"`
	WBSDescription                 string `json:"WBSDescription"`
	AccountAssignmentCategory      string `json:"AccountAssignmentCategory"`
	Reservation                    string `json:"Reservation"`
	PlannedOrderLongText           string `json:"PlannedOrderLongText"`
	MRPController                  string `json:"MRPController"`
	ProductionSupervisor           string `json:"ProductionSupervisor"`
	PurchasingGroup                string `json:"PurchasingGroup"`
	PurchasingOrganization         string `json:"PurchasingOrganization"`
	FixedSupplier                  string `json:"FixedSupplier"`
	PurchasingDocument             string `json:"PurchasingDocument"`
	PurchasingDocumentItem         string `json:"PurchasingDocumentItem"`
	QuotaArrangement               string `json:"QuotaArrangement"`
	QuotaArrangementItem           string `json:"QuotaArrangementItem"`
	SupplierName                   string `json:"SupplierName"`
	PlannedOrderIsFirm             bool   `json:"PlannedOrderIsFirm"`
	PlannedOrderIsConvertible      bool   `json:"PlannedOrderIsConvertible"`
	PlannedOrderBOMIsFixed         bool   `json:"PlannedOrderBOMIsFixed"`
	PlannedOrderCapacityIsDsptchd  bool   `json:"PlannedOrderCapacityIsDsptchd"`
	CapacityRequirement            string `json:"CapacityRequirement"`
	BillOfOperationsVariant        string `json:"BillOfOperationsVariant"`
	CapacityRequirementOrigin      string `json:"CapacityRequirementOrigin"`
	BillOfOperationsType           string `json:"BillOfOperationsType"`
	BillOfOperationsGroup          string `json:"BillOfOperationsGroup"`
	LastScheduledDate              string `json:"LastScheduledDate"`
	ScheduledBasicEndDate          string `json:"ScheduledBasicEndDate"`
	ScheduledBasicEndTime          string `json:"ScheduledBasicEndTime"`
	ScheduledBasicStartDate        string `json:"ScheduledBasicStartDate"`
	ScheduledBasicStartTime        string `json:"ScheduledBasicStartTime"`
	SchedulingType                 string `json:"SchedulingType"`
}

type Component struct {
	PlannedOrder                   string  `json:"PlannedOrder"`
	Reservation                    string  `json:"Reservation"`
	ReservationItem                string  `json:"ReservationItem"`
	BOMItem                        string  `json:"BOMItem"`
	BOMItemDescription             string  `json:"BOMItemDescription"`
	BillOfMaterialCategory         string  `json:"BillOfMaterialCategory"`
	BOMItemSorter                  string  `json:"BOMItemSorter"`
	BillOfMaterialItemNumber       string  `json:"BillOfMaterialItemNumber"`
	BillOfMaterialInternalID       string  `json:"BillOfMaterialInternalID"`
	BillOfMaterialVariant          string  `json:"BillOfMaterialVariant"`
	BOMItemCategory                string  `json:"BOMItemCategory"`
	Material                       string  `json:"Material"`
	MatlCompRequirementDate        string  `json:"MatlCompRequirementDate"`
	GoodsMovementEntryQty          int     `json:"GoodsMovementEntryQty"`
	EntryUnit                      string  `json:"EntryUnit"`
	RequiredQuantity               int     `json:"RequiredQuantity"`
	BaseUnit                       string  `json:"BaseUnit"`
	WithdrawnQuantity              int     `json:"WithdrawnQuantity"`
	DebitCreditCode                string  `json:"DebitCreditCode"`
	ComponentScrapInPercent        float64 `json:"ComponentScrapInPercent"`
	QuantityIsFixed                bool    `json:"QuantityIsFixed"`
	Plant                          string  `json:"Plant"`
	StorageLocation                string  `json:"StorageLocation"`
	SupplyArea                     string  `json:"SupplyArea"`
	MRPController                  string  `json:"MRPController"`
	PlannedOrderLastChangeDateTime string  `json:"PlannedOrderLastChangeDateTime"`
}
