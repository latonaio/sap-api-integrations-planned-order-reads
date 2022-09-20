package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-planned-order-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetPlannedOrder(plannedOrder, material, mRPPlant, plant string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(plannedOrder)
				wg.Done()
			}()
		case "HeaderMaterialPlant":
			func() {
				c.HeaderMaterialPlant(material, mRPPlant)
				wg.Done()
			}()
		case "ComponentMaterialPlant":
			func() {
				c.ComponentMaterialPlant(material, plant)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(plannedOrder string) {
	data, err := c.callPlannedOrderSrvAPIRequirementHeader("A_PlannedOrderHeader", plannedOrder)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPlannedOrderSrvAPIRequirementHeader(api, plannedOrder string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PLANNED_ORDER_SRV", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, plannedOrder)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) HeaderMaterialPlant(material, mRPPlant string) {
	data, err := c.callPlannedOrderSrvAPIRequirementHeaderMaterialPlant("A_PlannedOrderHeader", material, mRPPlant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPlannedOrderSrvAPIRequirementHeaderMaterialPlant(api, material, mRPPlant string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PLANNED_ORDER_SRV", api}, "/")
	param := c.getQueryWithHeaderMaterialPlant(map[string]string{}, material, mRPPlant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ComponentMaterialPlant(material, plant string) {
	data, err := c.callPlannedOrderSrvAPIRequirementComponentMaterialPlant("A_PlannedOrderComponent", material, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPlannedOrderSrvAPIRequirementComponentMaterialPlant(api, material, plant string) ([]sap_api_output_formatter.Component, error) {
	url := strings.Join([]string{c.baseURL, "API_PLANNED_ORDER_SRV", api}, "/")
	param := c.getQueryWithComponentMaterialPlant(map[string]string{}, material, plant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToComponent(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, plannedOrder string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PlannedOrder eq '%s'", plannedOrder)
	return params
}

func (c *SAPAPICaller) getQueryWithHeaderMaterialPlant(params map[string]string, material, mRPPlant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Material eq '%s' and MRPPlant eq '%s'", material, mRPPlant)
	return params
}

func (c *SAPAPICaller) getQueryWithComponentMaterialPlant(params map[string]string, material, plant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Material eq '%s' and Plant eq '%s'", material, plant)
	return params
}
