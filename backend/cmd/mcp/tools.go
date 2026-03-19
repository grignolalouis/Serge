package main

import (
	"encoding/json"
	"fmt"
	"serge/backend/internal/controllers"
)

// MCPRequest represents the JSON-RPC request
type MCPRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  MCPParams   `json:"params"`
	ID      interface{} `json:"id"`
}

// MCPParams contains the tool call parameters
type MCPParams struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

// MCPResponse represents the JSON-RPC response
type MCPResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *MCPError   `json:"error,omitempty"`
	ID      interface{} `json:"id"`
}

// MCPError represents an error response
type MCPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ToolResponse represents the tool result
type ToolResponse struct {
	Content []ToolContent `json:"content"`
}

// ToolContent represents tool content
type ToolContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// MCPHandler handles MCP protocol requests
type MCPHandler struct {
	omsRouter *controllers.OmsRouter
	tmsRouter *controllers.TmsRouter
	wmsRouter *controllers.WmsRouter
	srmRouter *controllers.SrmRouter
}

// NewMCPHandler creates a new MCP handler
func NewMCPHandler(oms *controllers.OmsRouter, tms *controllers.TmsRouter, wms *controllers.WmsRouter, srm *controllers.SrmRouter) *MCPHandler {
	return &MCPHandler{
		omsRouter: oms,
		tmsRouter: tms,
		wmsRouter: wms,
		srmRouter: srm,
	}
}

// HandleRequest handles incoming MCP requests via stdio
func (h *MCPHandler) HandleRequest(req *MCPRequest) *MCPResponse {
	// Handle initialize request
	if req.Method == "initialize" {
		return h.handleInitialize(req)
	}

	// Handle tools/list request
	if req.Method == "tools/list" {
		return h.handleToolsList(req)
	}

	// Handle tools/call request
	if req.Method == "tools/call" {
		return h.handleToolCall(req)
	}

	return &MCPResponse{
		JSONRPC: "2.0",
		Error: &MCPError{
			Code:    -32601,
			Message: "Method not found",
		},
		ID: req.ID,
	}
}

// handleInitialize handles the initialize request
func (h *MCPHandler) handleInitialize(req *MCPRequest) *MCPResponse {
	return &MCPResponse{
		JSONRPC: "2.0",
		Result: map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities": map[string]interface{}{
				"tools": map[string]interface{}{},
			},
			"serverInfo": map[string]interface{}{
				"name":    "serge-supply-chain",
				"version": "1.0.0",
			},
		},
		ID: req.ID,
	}
}

// handleToolsList lists all available tools
func (h *MCPHandler) handleToolsList(req *MCPRequest) *MCPResponse {
	tools := []map[string]interface{}{
		{
			"name":        "oms_get_order",
			"description": "Get order details by order number",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"order_number": map[string]interface{}{
						"type":        "string",
						"description": "The order number to retrieve",
					},
				},
				"required": []string{"order_number"},
			},
		},
		{
			"name":        "tms_get_shipment_by_order",
			"description": "Get shipment details by order number",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"order_number": map[string]interface{}{
						"type":        "string",
						"description": "The order number to retrieve shipment for",
					},
				},
				"required": []string{"order_number"},
			},
		},
		{
			"name":        "wms_get_inventory_by_sku",
			"description": "Get inventory details by SKU",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"sku": map[string]interface{}{
						"type":        "string",
						"description": "The SKU to retrieve inventory for",
					},
				},
				"required": []string{"sku"},
			},
		},
		{
			"name":        "srm_get_suppliers_by_category",
			"description": "Get suppliers by category",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"category": map[string]interface{}{
						"type":        "string",
						"description": "The supplier category to filter by",
					},
				},
				"required": []string{"category"},
			},
		},
		{
			"name":        "srm_get_purchase_orders_by_sku",
			"description": "Get purchase orders by SKU",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"sku": map[string]interface{}{
						"type":        "string",
						"description": "The SKU to retrieve purchase orders for",
					},
				},
				"required": []string{"sku"},
			},
		},
	}

	return &MCPResponse{
		JSONRPC: "2.0",
		Result:  map[string]interface{}{"tools": tools},
		ID:      req.ID,
	}
}

// handleToolCall handles tool execution
func (h *MCPHandler) handleToolCall(req *MCPRequest) *MCPResponse {
	toolName := req.Params.Name

	var result interface{}
	var errMsg string

	switch toolName {
	case "oms_get_order":
		result, errMsg = h.callOmsGetOrder(req.Params.Arguments)
	case "tms_get_shipment_by_order":
		result, errMsg = h.callTmsGetShipment(req.Params.Arguments)
	case "wms_get_inventory_by_sku":
		result, errMsg = h.callWmsGetInventory(req.Params.Arguments)
	case "srm_get_suppliers_by_category":
		result, errMsg = h.callSrmGetSuppliers(req.Params.Arguments)
	case "srm_get_purchase_orders_by_sku":
		result, errMsg = h.callSrmGetPurchaseOrders(req.Params.Arguments)
	default:
		return &MCPResponse{
			JSONRPC: "2.0",
			Error: &MCPError{
				Code:    -32601,
				Message: fmt.Sprintf("Tool not found: %s", toolName),
			},
			ID: req.ID,
		}
	}

	if errMsg != "" {
		return &MCPResponse{
			JSONRPC: "2.0",
			Error: &MCPError{
				Code:    -32000,
				Message: errMsg,
			},
			ID: req.ID,
		}
	}

	return &MCPResponse{
		JSONRPC: "2.0",
		Result: ToolResponse{
			Content: []ToolContent{
				{
					Type: "text",
					Text: result.(string),
				},
			},
		},
		ID: req.ID,
	}
}

// callOmsGetOrder calls the OMS get order handler
func (h *MCPHandler) callOmsGetOrder(args map[string]interface{}) (interface{}, string) {
	orderNumber, ok := args["order_number"].(string)
	if !ok || orderNumber == "" {
		return nil, "Missing or invalid order_number parameter"
	}

	output, err := h.omsRouter.HandleGetOrderByNumber(
		nil, // context - will be provided by caller
		controllers.GetOrderByNumberInput{OrderNumber: orderNumber},
	)
	if err != nil {
		return nil, fmt.Sprintf("Failed to get order: %v", err)
	}

	b, err := json.Marshal(output)
	if err != nil {
		return nil, fmt.Sprintf("Failed to marshal response: %v", err)
	}

	return string(b), ""
}

// callTmsGetShipment calls the TMS get shipment handler
func (h *MCPHandler) callTmsGetShipment(args map[string]interface{}) (interface{}, string) {
	orderNumber, ok := args["order_number"].(string)
	if !ok || orderNumber == "" {
		return nil, "Missing or invalid order_number parameter"
	}

	output, err := h.tmsRouter.HandleGetShipmentByOrderNumber(
		nil,
		controllers.GetShipmentByOrderNumberInput{OrderNumber: orderNumber},
	)
	if err != nil {
		return nil, fmt.Sprintf("Failed to get shipment: %v", err)
	}

	b, err := json.Marshal(output)
	if err != nil {
		return nil, fmt.Sprintf("Failed to marshal response: %v", err)
	}

	return string(b), ""
}

// callWmsGetInventory calls the WMS get inventory handler
func (h *MCPHandler) callWmsGetInventory(args map[string]interface{}) (interface{}, string) {
	sku, ok := args["sku"].(string)
	if !ok || sku == "" {
		return nil, "Missing or invalid sku parameter"
	}

	output, err := h.wmsRouter.HandleGetInventoryBySKU(
		nil,
		controllers.GetInventoryBySKUInput{SKU: sku},
	)
	if err != nil {
		return nil, fmt.Sprintf("Failed to get inventory: %v", err)
	}

	b, err := json.Marshal(output)
	if err != nil {
		return nil, fmt.Sprintf("Failed to marshal response: %v", err)
	}

	return string(b), ""
}

// callSrmGetSuppliers calls the SRM get suppliers handler
func (h *MCPHandler) callSrmGetSuppliers(args map[string]interface{}) (interface{}, string) {
	category, ok := args["category"].(string)
	if !ok || category == "" {
		return nil, "Missing or invalid category parameter"
	}

	output, err := h.srmRouter.HandleGetSuppliersByCategory(
		nil,
		controllers.GetSuppliersByCategoryInput{Category: category},
	)
	if err != nil {
		return nil, fmt.Sprintf("Failed to get suppliers: %v", err)
	}

	b, err := json.Marshal(output)
	if err != nil {
		return nil, fmt.Sprintf("Failed to marshal response: %v", err)
	}

	return string(b), ""
}

// callSrmGetPurchaseOrders calls the SRM get purchase orders handler
func (h *MCPHandler) callSrmGetPurchaseOrders(args map[string]interface{}) (interface{}, string) {
	sku, ok := args["sku"].(string)
	if !ok || sku == "" {
		return nil, "Missing or invalid sku parameter"
	}

	output, err := h.srmRouter.HandleGetPurchaseOrdersBySKU(
		nil,
		controllers.GetPurchaseOrdersBySKUInput{SKU: sku},
	)
	if err != nil {
		return nil, fmt.Sprintf("Failed to get purchase orders: %v", err)
	}

	b, err := json.Marshal(output)
	if err != nil {
		return nil, fmt.Sprintf("Failed to marshal response: %v", err)
	}

	return string(b), ""
}

