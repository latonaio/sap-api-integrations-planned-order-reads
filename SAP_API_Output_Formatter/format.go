package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-planned-order-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		header = append(header, Header{
			PlannedOrder:                   data.PlannedOrder,
			PlannedOrderType:               data.PlannedOrderType,
			PlannedOrderProfile:            data.PlannedOrderProfile,
			Material:                       data.Material,
			MaterialName:                   data.MaterialName,
			ProductionPlant:                data.ProductionPlant,
			MRPPlant:                       data.MRPPlant,
			MRPArea:                        data.MRPArea,
			ProductionVersion:              data.ProductionVersion,
			MaterialProcurementCategory:    data.MaterialProcurementCategory,
			MaterialProcurementType:        data.MaterialProcurementType,
			StorageLocation:                data.StorageLocation,
			BaseUnit:                       data.BaseUnit,
			TotalQuantity:                  data.TotalQuantity,
			PlndOrderPlannedScrapQty:       data.PlndOrderPlannedScrapQty,
			GoodsReceiptQty:                data.GoodsReceiptQty,
			IssuedQuantity:                 data.IssuedQuantity,
			PlndOrderPlannedStartDate:      data.PlndOrderPlannedStartDate,
			PlndOrderPlannedStartTime:      data.PlndOrderPlannedStartTime,
			PlndOrderPlannedEndDate:        data.PlndOrderPlannedEndDate,
			PlndOrderPlannedEndTime:        data.PlndOrderPlannedEndTime,
			PlannedOrderOpeningDate:        data.PlannedOrderOpeningDate,
			PlannedOrderLastChangeDateTime: data.PlannedOrderLastChangeDateTime,
			ProductionStartDate:            data.ProductionStartDate,
			ProductionEndDate:              data.ProductionEndDate,
			SalesOrder:                     data.SalesOrder,
			SalesOrderItem:                 data.SalesOrderItem,
			Customer:                       data.Customer,
			WBSElementInternalID:           data.WBSElementInternalID,
			WBSElementExternalID:           data.WBSElementExternalID,
			WBSDescription:                 data.WBSDescription,
			AccountAssignmentCategory:      data.AccountAssignmentCategory,
			Reservation:                    data.Reservation,
			PlannedOrderLongText:           data.PlannedOrderLongText,
			MRPController:                  data.MRPController,
			ProductionSupervisor:           data.ProductionSupervisor,
			PurchasingGroup:                data.PurchasingGroup,
			PurchasingOrganization:         data.PurchasingOrganization,
			FixedSupplier:                  data.FixedSupplier,
			PurchasingDocument:             data.PurchasingDocument,
			PurchasingDocumentItem:         data.PurchasingDocumentItem,
			QuotaArrangement:               data.QuotaArrangement,
			QuotaArrangementItem:           data.QuotaArrangementItem,
			SupplierName:                   data.SupplierName,
			PlannedOrderIsFirm:             data.PlannedOrderIsFirm,
			PlannedOrderIsConvertible:      data.PlannedOrderIsConvertible,
			PlannedOrderBOMIsFixed:         data.PlannedOrderBOMIsFixed,
			PlannedOrderCapacityIsDsptchd:  data.PlannedOrderCapacityIsDsptchd,
			CapacityRequirement:            data.CapacityRequirement,
			BillOfOperationsVariant:        data.BillOfOperationsVariant,
			CapacityRequirementOrigin:      data.CapacityRequirementOrigin,
			BillOfOperationsType:           data.BillOfOperationsType,
			BillOfOperationsGroup:          data.BillOfOperationsGroup,
			LastScheduledDate:              data.LastScheduledDate,
			ScheduledBasicEndDate:          data.ScheduledBasicEndDate,
			ScheduledBasicEndTime:          data.ScheduledBasicEndTime,
			ScheduledBasicStartDate:        data.ScheduledBasicStartDate,
			ScheduledBasicStartTime:        data.ScheduledBasicStartTime,
			SchedulingType:                 data.SchedulingType,
		})
	}

	return header, nil
}

func ConvertToComponent(raw []byte, l *logger.Logger) ([]Component, error) {
	pm := &responses.Component{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Component. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	component := make([]Component, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		component = append(component, Component{
			PlannedOrder:                   data.PlannedOrder,
			Reservation:                    data.Reservation,
			ReservationItem:                data.ReservationItem,
			BOMItem:                        data.BOMItem,
			BOMItemDescription:             data.BOMItemDescription,
			BillOfMaterialCategory:         data.BillOfMaterialCategory,
			BOMItemSorter:                  data.BOMItemSorter,
			BillOfMaterialItemNumber:       data.BillOfMaterialItemNumber,
			BillOfMaterialInternalID:       data.BillOfMaterialInternalID,
			BillOfMaterialVariant:          data.BillOfMaterialVariant,
			BOMItemCategory:                data.BOMItemCategory,
			Material:                       data.Material,
			MatlCompRequirementDate:        data.MatlCompRequirementDate,
			GoodsMovementEntryQty:          data.GoodsMovementEntryQty,
			EntryUnit:                      data.EntryUnit,
			RequiredQuantity:               data.RequiredQuantity,
			BaseUnit:                       data.BaseUnit,
			WithdrawnQuantity:              data.WithdrawnQuantity,
			DebitCreditCode:                data.DebitCreditCode,
			ComponentScrapInPercent:        data.ComponentScrapInPercent,
			QuantityIsFixed:                data.QuantityIsFixed,
			Plant:                          data.Plant,
			StorageLocation:                data.StorageLocation,
			SupplyArea:                     data.SupplyArea,
			MRPController:                  data.MRPController,
			PlannedOrderLastChangeDateTime: data.PlannedOrderLastChangeDateTime,
		})
	}

	return component, nil
}
