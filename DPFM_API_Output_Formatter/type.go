package dpfm_api_output_formatter

type DeliverySDC struct {
	ConnectionKey       string             `json:"connection_key"`
	Result              bool               `json:"result"`
	RedisKey            string             `json:"redis_key"`
	Filepath            string             `json:"filepath"`
	APIStatusCode       int                `json:"api_status_code"`
	RuntimeSessionID    string             `json:"runtime_session_id"`
	BusinessPartnerID   *int               `json:"business_partner"`
	ServiceLabel        string             `json:"service_label"`
	APIType             string             `json:"api_type"`
	DataConcatenation   *DataConcatenation `json:"DataConcatenation"`
	APISchema           string             `json:"api_schema"`
	Accepter            []string           `json:"accepter"`
	Deleted             bool               `json:"deleted"`
	APIProcessingResult bool               `json:"api_processing_result"`
	APIProcessingError  *string            `json:"api_processing_error"`
}

type DataConcatenation struct {
	Header DeliveryNoticeHeader `json:"DeliveryNoticeHeader"`
	Item   []DeliveryNoticeItem `json:"DeliveryNoticeItem"`
}

type DeliveryNoticeHeader struct {
	ExchangedDeliveryNoticeDocumentIdentifier                      string   `json:"ExchangedDeliveryNoticeDocumentIdentifier" csv:"1"`
	DeliveryNoticeDocument                                         *string  `json:"DeliveryNoticeDocument" csv:"2"`
	ExchangedDocumentContextSpecifiedTransactionIdentifier         *string  `json:"ExchangedDocumentContextSpecifiedTransactionIdentifier" csv:"3"`
	ExchangedDeliveryNoticeDocumentName                            *string  `json:"ExchangedDeliveryNoticeDocumentName" csv:"4"`
	ExchangeDeliveryNoticeDocumentTypeCode                         *string  `json:"ExchangeDeliveryNoticeDocumentTypeCode" csv:"5"`
	ExchangedDeliveryNoticeDocumentIssueDate                       *string  `json:"ExchangedDeliveryNoticeDocumentIssueDate" csv:"6"`
	ExchangedDeliveryNoticeDocumentPurposeCode                     *string  `json:"ExchangedDeliveryNoticeDocumentPurposeCode" csv:"7"`
	ExchangedDeliveryNoticeDocumentRequestedResponseTypeCode       *string  `json:"ExchangedDeliveryNoticeDocumentRequestedResponseTypeCode" csv:"8"`
	ExchangedDeliveryNoticeDocumentVersionIdentifier               *string  `json:"ExchangedDeliveryNoticeDocumentVersionIdentifier" csv:"9"`
	ExchangedDeliveryNoticeDocumentCategoryCode                    *string  `json:"ExchangedDeliveryNoticeDocumentCategoryCode" csv:"10"`
	ExchangedDeliveryNoticeDocumentSubtypeCode                     *string  `json:"ExchangedDeliveryNoticeDocumentSubtypeCode" csv:"11"`
	NoteDeliveryNoticeSubjectText                                  *string  `json:"NoteDeliveryNoticeSubjectText" csv:"12"`
	NoteDeliveryNoticeContentText                                  *string  `json:"NoteDeliveryNoticeContentText" csv:"13"`
	NoteDeliveryNoticeIdentifier                                   *string  `json:"NoteDeliveryNoticeIdentifier" csv:"14"`
	SpecifiedBinaryFileIdentifier                                  *string  `json:"SpecifiedBinaryFileIdentifier" csv:"15"`
	SpecifiedBinaryFileVersionIdentifier                           *string  `json:"SpecifiedBinaryFileVersionIdentifier" csv:"16"`
	SpecifiedBinaryFileNameText                                    *string  `json:"SpecifiedBinaryFileNameText" csv:"17"`
	SpecifiedBineryFileURIIdentifier                               *string  `json:"SpecifiedBineryFileURIIdentifier" csv:"18"`
	SpecifiedBineryFileMIMECode                                    *string  `json:"SpecifiedBineryFileMIMECode" csv:"19"`
	SpecifiedBineryFileEncodingCode                                *string  `json:"SpecifiedBineryFileEncodingCode" csv:"20"`
	SpecifiedBineryFileCode                                        *string  `json:"SpecifiedBineryFileCode" csv:"21"`
	SpecifiedBinaryFileDescriptionText                             *string  `json:"SpecifiedBinaryFileDescriptionText" csv:"22"`
	TradeSellerIdentifier                                          *string  `json:"TradeSellerIdentifier" csv:"23"`
	TradeSellerGlobalIdentifier                                    *string  `json:"TradeSellerGlobalIdentifier" csv:"24"`
	TradeSellerName                                                *string  `json:"TradeSellerName" csv:"25"`
	TradeBillFromPartyRegisteredIdentifier                         *string  `json:"TradeBillFromPartyRegisteredIdentifier" csv:"26"`
	TradeContactSellerIdentifier                                   *string  `json:"TradeContactSellerIdentifier" csv:"27"`
	TradeContactSellerPersonName                                   *string  `json:"TradeContactSellerPersonName" csv:"28"`
	TradeContactSellerDepartmentName                               *string  `json:"TradeContactSellerDepartmentName" csv:"29"`
	SellerTelephoneNumber                                          *string  `json:"SellerTelephoneNumber" csv:"30"`
	SellerFaxNumber                                                *string  `json:"SellerFaxNumber" csv:"31"`
	SellerEmailAddress                                             *string  `json:"SellerEmailAddress" csv:"32"`
	SellerAddressPostalCode                                        *string  `json:"SellerAddressPostalCode" csv:"33"`
	SellerAddress                                                  *string  `json:"SellerAddress" csv:"34"`
	TradeBuyerIdentifier                                           *string  `json:"TradeBuyerIdentifier" csv:"35"`
	TradeBuyerGlobalIdentifier                                     *string  `json:"TradeBuyerGlobalIdentifier" csv:"36"`
	TradeBuyerName                                                 *string  `json:"TradeBuyerName" csv:"37"`
	TradeContactBuyerIdentifier                                    *string  `json:"TradeContactBuyerIdentifier" csv:"38"`
	TradeContactBuyerPersonName                                    *string  `json:"TradeContactBuyerPersonName" csv:"39"`
	TradeContactBuyerDepartmentName                                *string  `json:"TradeContactBuyerDepartmentName" csv:"40"`
	BuyerTelephoneNumber                                           *string  `json:"BuyerTelephoneNumber" csv:"41"`
	BuyerFaxNumber                                                 *string  `json:"BuyerFaxNumber" csv:"42"`
	BuyerEmailAddress                                              *string  `json:"BuyerEmailAddress" csv:"43"`
	BuyerAddressPostalCode                                         *string  `json:"BuyerAddressPostalCode" csv:"44"`
	BuyerAddress                                                   *string  `json:"BuyerAddress" csv:"45"`
	ReferencedOrdersDocumentIssureAssignedIdentifier               *string  `json:"ReferencedOrdersDocumentIssureAssignedIdentifier" csv:"46"`
	ReferencedOrdersDocumentIssueDate                              *string  `json:"ReferencedOrdersDocumentIssueDate" csv:"47"`
	ReferencedOrdersDocumentRevisionIdentifier                     *string  `json:"ReferencedOrdersDocumentRevisionIdentifier" csv:"48"`
	ReferencedOrdersDocumentName                                   *string  `json:"ReferencedOrdersDocumentName" csv:"49"`
	ReferencedOrdersDocumentInformationText                        *string  `json:"ReferencedOrdersDocumentInformationText" csv:"50"`
	ReferencedOrdersDocumentInformationPurposeCode                 *string  `json:"ReferencedOrdersDocumentInformationPurposeCode" csv:"51"`
	ReferencedOrdersDocumentInformationSubtypeCode                 *string  `json:"ReferencedOrdersDocumentInformationSubtypeCode" csv:"52"`
	ProjectIdentifier                                              *string  `json:"ProjectIdentifier" csv:"53"`
	ProjectName                                                    *string  `json:"ProjectName" csv:"54"`
	ReferencedSalesOrderDocumentIssureAddignedIdentifier           *string  `json:"ReferencedSalesOrderDocumentIssureAddignedIdentifier" csv:"55"`
	ReferencedSalesOrderDocumentIssueDate                          *string  `json:"ReferencedSalesOrderDocumentIssueDate" csv:"56"`
	ReferencedSalesOrderDocumentRevisionIdentifier                 *string  `json:"ReferencedSalesOrderDocumentRevisionIdentifier" csv:"57"`
	ReferencedSalesOrderDocumentName                               *string  `json:"ReferencedSalesOrderDocumentName" csv:"58"`
	ReferencedSalesOrderDocumentInformationText                    *string  `json:"ReferencedSalesOrderDocumentInformationText" csv:"59"`
	ReferencedSalesOrderDocumentSubtypeCode                        *string  `json:"ReferencedSalesOrderDocumentSubtypeCode" csv:"60"`
	TradeShipToPartyIdentifier                                     *string  `json:"TradeShipToPartyIdentifier" csv:"61"`
	TradeShipToPartyGlobalIdentifier                               *string  `json:"TradeShipToPartyGlobalIdentifier" csv:"62"`
	TradeShipToPartyName                                           *string  `json:"TradeShipToPartyName" csv:"63"`
	TradeShipToPartyContactIdentifier                              *string  `json:"TradeShipToPartyContactIdentifier" csv:"64"`
	TradeShipToPartyContactPersonName                              *string  `json:"TradeShipToPartyContactPersonName" csv:"65"`
	TradeShipToPartyContactDepartmentName                          *string  `json:"TradeShipToPartyContactDepartmentName" csv:"66"`
	TradeShipToPartyContactPersonIdentifier                        *string  `json:"TradeShipToPartyContactPersonIdentifier" csv:"67"`
	ShipToPartyTelephoneNumber                                     *string  `json:"ShipToPartyTelephoneNumber" csv:"68"`
	ShipToPartyFaxNumber                                           *string  `json:"ShipToPartyFaxNumber" csv:"69"`
	ShipToPartyEmailAddress                                        *string  `json:"ShipToPartyEmailAddress" csv:"70"`
	ShipToPartyAddressPostalCode                                   *string  `json:"ShipToPartyAddressPostalCode" csv:"71"`
	ShipToPartyAddress                                             *string  `json:"ShipToPartyAddress" csv:"72"`
	LastShipToPartyIdentifier                                      *string  `json:"LastShipToPartyIdentifier" csv:"73"`
	TradeShipFromPartyIdentifier                                   *string  `json:"TradeShipFromPartyIdentifier" csv:"74"`
	TradeShipFromPartyName                                         *string  `json:"TradeShipFromPartyName" csv:"75"`
	TradeLogiName                                                  *string  `json:"TradeLogiName" csv:"76"`
	TradeLogiContactIdentifier                                     *string  `json:"TradeLogiContactIdentifier" csv:"77"`
	TradeLogiContactPersonName                                     *string  `json:"TradeLogiContactPersonName" csv:"78"`
	TradeLogiContactDepartmentName                                 *string  `json:"TradeLogiContactDepartmentName" csv:"79"`
	TradeLogiContactPersonIdentifier                               *string  `json:"TradeLogiContactPersonIdentifier" csv:"80"`
	LogiTelephoneNumber                                            *string  `json:"LogiTelephoneNumber" csv:"81"`
	SupplyChainEventIdentifier                                     *string  `json:"SupplyChainEventIdentifier" csv:"82"`
	InstructedTemperatureControlCode                               *string  `json:"InstructedTemperatureControlCode" csv:"83"`
	TradeTaxCalculatedAmount                                       *float32 `json:"TradeTaxCalculatedAmount" csv:"84"`
	TradeTaxTypeCode                                               *string  `json:"TradeTaxTypeCode" csv:"85"`
	TradeTaxBasisAmount                                            *float32 `json:"TradeTaxBasisAmount" csv:"86"`
	TradeTaxCategoryCode                                           *string  `json:"TradeTaxCategoryCode" csv:"87"`
	TradeTaxCategoryName                                           *string  `json:"TradeTaxCategoryName" csv:"88"`
	TradeTaxRatePercent                                            *float32 `json:"TradeTaxRatePercent" csv:"89"`
	TradeTaxGrandTotalAmount                                       *float32 `json:"TradeTaxGrandTotalAmount" csv:"90"`
	TradeTaxCalculationMethod                                      *string  `json:"TradeTaxCalculationMethod" csv:"91"`
	TradeSettlementMonetarySummationTotalTaxAmount                 *float32 `json:"TradeSettlementMonetarySummationTotalTaxAmount" csv:"92"`
	TradeDeliveryNoticeSettlementMonetarySummationGrandTotalAmount *float32 `json:"TradeDeliveryNoticeSettlementMonetarySummationGrandTotalAmount" csv:"93"`
	TradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount   *float32 `json:"TradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount" csv:"94"`
	TradeDeliveryNoticeMonetarySummationIncludingTaxesTotalAmount  *float32 `json:"TradeDeliveryNoticeMonetarySummationIncludingTaxesTotalAmount" csv:"95"`
}

type DeliveryNoticeItem struct {
	ExchangedDeliveryNoticeDocumentIdentifier                                      string   `json:"ExchangedDeliveryNoticeDocumentIdentifier" csv:"1"`
	DeliveryNoticeDocumentItemlineIdentifier                                       string   `json:"DeliveryNoticeDocumentItemlineIdentifier" csv:"96"`
	DeliveryNoticeDocumentItemlineStatusCode                                       *string  `json:"DeliveryNoticeDocumentItemlineStatusCode" csv:"97"`
	DeliveryNoticeDocumentItemlineStatusReasonCode                                 *string  `json:"DeliveryNoticeDocumentItemlineStatusReasonCode" csv:"98"`
	NoteDeliveryNoticeItemSubjectText                                              *string  `json:"NoteDeliveryNoticeItemSubjectText" csv:"99"`
	NoteDeliveryNoticeItemContentText                                              *string  `json:"NoteDeliveryNoticeItemContentText" csv:"100"`
	NoteDeliveryNoticeItemIdentifier                                               *string  `json:"NoteDeliveryNoticeItemIdentifier" csv:"101"`
	ReferencedSalesOrderDocumentIssuerAssignedIdentifier                           *string  `json:"ReferencedSalesOrderDocumentIssuerAssignedIdentifier" csv:"102"`
	ReferencedSalesOrderDocumentItemLineIdentifier                                 *string  `json:"ReferencedSalesOrderDocumentItemLineIdentifier" csv:"103"`
	ReferencedSalesOrderDocumentRevisionIdentifier                                 *string  `json:"ReferencedSalesOrderDocumentRevisionIdentifier" csv:"104"`
	ReferencedOrdersDocumentIssureAssignedIdentifier                               *string  `json:"ReferencedOrdersDocumentIssureAssignedIdentifier" csv:"105"`
	ReferencedOrdersDocumentItemLineIdentifier                                     *string  `json:"ReferencedOrdersDocumentItemLineIdentifier" csv:"106"`
	ReferencedOrdersDocumentRevisionIdentifier                                     *string  `json:"ReferencedOrdersDocumentRevisionIdentifier" csv:"107"`
	TradePriceTypeCode                                                             *string  `json:"TradePriceTypeCode" csv:"108"`
	TradeDeliveryPriceChargeAmount                                                 *float32 `json:"TradeDeliveryPriceChargeAmount" csv:"109"`
	TradePriceBasisQuantity                                                        *float32 `json:"TradePriceBasisQuantity" csv:"110"`
	TradePriceBasisUnitCode                                                        *string  `json:"TradePriceBasisUnitCode" csv:"111"`
	OrderQuantityInBaseUnit                                                        *float32 `json:"OrderQuantityInBaseUnit" csv:"112"`
	DeliveryQuantityInBaseUnit                                                     *float32 `json:"DeliveryQuantityInBaseUnit" csv:"113"`
	SupplyChainTradeDeliveryPackageQuantity                                        *float32 `json:"SupplyChainTradeDeliveryPackageQuantity" csv:"114"`
	SupplyChainTradeDeliveryProductUnitQuantity                                    *float32 `json:"SupplyChainTradeDeliveryProductUnitQuantity" csv:"115"`
	SupplyChainTradeDeliveryPerPackageUnitQuantity                                 *float32 `json:"SupplyChainTradeDeliveryPerPackageUnitQuantity" csv:"116"`
	SupplyChainTradeDeliveryDespatchedQuantity                                     *float32 `json:"SupplyChainTradeDeliveryDespatchedQuantity" csv:"117"`
	SupplyChainTradeDeliveryRequestedQuantity                                      *float32 `json:"SupplyChainTradeDeliveryRequestedQuantity" csv:"118"`
	SupplyChainTradeOrderRequestedPackageQuantity                                  *float32 `json:"SupplyChainTradeOrderRequestedPackageQuantity" csv:"119"`
	SupplyChainTradeDeliveryRequestedPackageQuantity                               *float32 `json:"SupplyChainTradeDeliveryRequestedPackageQuantity" csv:"120"`
	SupplyChainTradeDeliveryRemainingRequestedQuantity                             *float32 `json:"SupplyChainTradeDeliveryRemainingRequestedQuantity" csv:"121"`
	SupplyChainTradeDeliveryBilledQuantity                                         *float32 `json:"SupplyChainTradeDeliveryBilledQuantity" csv:"122"`
	ItemTradeDeliverToPartyIdentifier                                              *string  `json:"ItemTradeDeliverToPartyIdentifier" csv:"123"`
	ItemTradeDeliverToPartyGlobalIdentifier                                        *string  `json:"ItemTradeDeliverToPartyGlobalIdentifier" csv:"124"`
	ItemTradeDeliverToPartyName                                                    *string  `json:"ItemTradeDeliverToPartyName" csv:"125"`
	ItemTradeDeliverToPartyContactPersonName                                       *string  `json:"ItemTradeDeliverToPartyContactPersonName" csv:"126"`
	ItemTradeDeliverToPartyContactDepartmentName                                   *string  `json:"ItemTradeDeliverToPartyContactDepartmentName" csv:"127"`
	ItemDeliverToPartyPhoneNumber                                                  *string  `json:"ItemDeliverToPartyPhoneNumber" csv:"128"`
	ItemDeliverToPartyFaxNumber                                                    *string  `json:"ItemDeliverToPartyFaxNumber" csv:"129"`
	ItemDeliverToPartyEMailAddress                                                 *string  `json:"ItemDeliverToPartyEMailAddress" csv:"130"`
	ItemDeliverToPartyAddressPostalCode                                            *string  `json:"ItemDeliverToPartyAddressPostalCode" csv:"131"`
	ItemDeliverToPartyAddress                                                      *string  `json:"ItemDeliverToPartyAddress" csv:"132"`
	SupplyChainDeliveryEventIdentifier                                             *string  `json:"SupplyChainDeliveryEventIdentifier" csv:"133"`
	SupplyChainDeliveryEventOccurrenceDate                                         *string  `json:"SupplyChainDeliveryEventOccurrenceDate" csv:"134"`
	LastShipToPartyDeliveryDate                                                    *string  `json:"SupplyChainEventTypeCode" csv:"135"`
	SupplyChainEventTypeCode                                                       *string  `json:"LastShipToPartyDeliveryDate" csv:"136"`
	SupplyChainEventRequirementOccurrenceDate                                      *string  `json:"SupplyChainEventRequirementOccurrenceDate" csv:"137"`
	LogisticsLocationIdentification                                                *string  `json:"LogisticsLocationIdentification" csv:"138"`
	LogisticsLocationName                                                          *string  `json:"LogisticsLocationName" csv:"139"`
	SupplyChainEventPlannedOccurrenceDate                                          *string  `json:"SupplyChainEventPlannedOccurrenceDate" csv:"140"`
	TradeTaxTypeCode                                                               *string  `json:"TradeTaxTypeCode" csv:"141"`
	ItemTradeTaxBasisAmount                                                        *float32 `json:"ItemTradeTaxBasisAmount" csv:"142"`
	ItemTradeTaxCategoryCode                                                       *string  `json:"ItemTradeTaxCategoryCode" csv:"143"`
	TradeTaxCategoryName                                                           *string  `json:"TradeTaxCategoryName" csv:"144"`
	ItemTradeTaxRateApplicablePercent                                              *float32 `json:"ItemTradeTaxRateApplicablePercent" csv:"145"`
	ItemTradeTaxGrandTotalAmount                                                   *float32 `json:"ItemTradeTaxGrandTotalAmount" csv:"146"`
	ItemTradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount               *float32 `json:"ItemTradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount" csv:"147"`
	ItemTradeSettlementMonetarySummationTaxAmount                                  *float32 `json:"ItemTradeSettlementMonetarySummationTaxAmount" csv:"148"`
	ItemTradeDeliveryNoticeSettlementMonetarySummationIncludingTaxesNetTotalAmount *float32 `json:"ItemTradeDeliveryNoticeSettlementMonetarySummationIncludingTaxesNetTotalAmount" csv:"149"`
	TradeProductIdentifier                                                         *string  `json:"TradeProductIdentifier" csv:"150"`
	TradeProductGlobalIdentifier                                                   *string  `json:"TradeProductGlobalIdentifier" csv:"151"`
	TradeProductSellerAssignedIdentifier                                           *string  `json:"TradeProductSellerAssignedIdentifier" csv:"152"`
	TradeProductBuyerAssignedIdentifier                                            *string  `json:"TradeProductBuyerAssignedIdentifier" csv:"153"`
	TradeProductManufacturerAssignedIdentifier                                     *string  `json:"TradeProductManufacturerAssignedIdentifier" csv:"154"`
	TradeProductName                                                               *string  `json:"TradeProductName" csv:"155"`
	TradeProductDescription                                                        *string  `json:"TradeProductDescription" csv:"156"`
	TradeProductTypeCode                                                           *string  `json:"TradeProductTypeCode" csv:"157"`
	TradeProductEndItemTypeCode                                                    *string  `json:"TradeProductEndItemTypeCode" csv:"158"`
	TradeProductSizeCode                                                           *string  `json:"TradeProductSizeCode" csv:"159"`
	TradeProductSizeDescriptionText                                                *string  `json:"TradeProductSizeDescriptionText" csv:"160"`
	TradeManufacturerIdentifier                                                    *string  `json:"TradeManufacturerIdentifier" csv:"161"`
	TradeManufacturerName                                                          *string  `json:"TradeManufacturerName" csv:"162"`
	ReferencedLogisticsPackageUnitQuantity                                         *float32 `json:"ReferencedLogisticsPackageUnitQuantity" csv:"163"`
	ReferencedLogisticsPackageQuantityUnitCode                                     *string  `json:"ReferencedLogisticsPackageQuantityUnitCode" csv:"164"`
	ReferencedLogisticsPackageTypeCode                                             *string  `json:"ReferencedLogisticsPackageTypeCode" csv:"165"`
	ReferencedLogisticsPackageIdentifier                                           *string  `json:"ReferencedLogisticsPackageIdentifier" csv:"166"`
}
