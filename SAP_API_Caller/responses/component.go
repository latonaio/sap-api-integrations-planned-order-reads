package responses

type Component struct {
	Value             []struct {
		PlannedOrder                   string    `json:"PlannedOrder"`
		Reservation                    string    `json:"Reservation"`
		ReservationItem                string    `json:"ReservationItem"`
		BOMItem                        string    `json:"BOMItem"`
		BOMItemDescription             string    `json:"BOMItemDescription"`
		BillOfMaterialCategory         string    `json:"BillOfMaterialCategory"`
		BOMItemSorter                  string    `json:"BOMItemSorter"`
		BillOfMaterialItemNumber       string    `json:"BillOfMaterialItemNumber"`
		BillOfMaterialInternalID       string    `json:"BillOfMaterialInternalID"`
		BillOfMaterialVariant          string    `json:"BillOfMaterialVariant"`
		BOMItemCategory                string    `json:"BOMItemCategory"`
		Material                       string    `json:"Material"`
		MatlCompRequirementDate        string    `json:"MatlCompRequirementDate"`
		GoodsMovementEntryQty          int       `json:"GoodsMovementEntryQty"`
		EntryUnit                      string    `json:"EntryUnit"`
		RequiredQuantity               int       `json:"RequiredQuantity"`
		BaseUnit                       string    `json:"BaseUnit"`
		WithdrawnQuantity              int       `json:"WithdrawnQuantity"`
		DebitCreditCode                string    `json:"DebitCreditCode"`
		ComponentScrapInPercent        float64   `json:"ComponentScrapInPercent"`
		QuantityIsFixed                bool      `json:"QuantityIsFixed"`
		Plant                          string    `json:"Plant"`
		StorageLocation                string    `json:"StorageLocation"`
		SupplyArea                     string    `json:"SupplyArea"`
		MRPController                  string    `json:"MRPController"`
		PlannedOrderLastChangeDateTime string    `json:"PlannedOrderLastChangeDateTime"`
	} `json:"value"`
}
