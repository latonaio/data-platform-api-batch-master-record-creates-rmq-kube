package sub_func_complementer

import (
	"context"
	dpfm_api_input_reader "data-platform-api-production-order-confirmation-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-production-order-confirmation-creates-rmq-kube/config"
	"encoding/json"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type SubFuncComplementer struct {
	ctx context.Context
	c   *config.Conf
	rmq *rabbitmq.RabbitmqClient
	db  *database.Mysql
}

func NewSubFuncComplementer(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient, db *database.Mysql) *SubFuncComplementer {
	return &SubFuncComplementer{
		ctx: ctx,
		c:   c,
		rmq: rmq,
		db:  db,
	}
}

func (c *SubFuncComplementer) ComplementHeader(input *dpfm_api_input_reader.SDC, subfuncSDC *SDC, l *logger.Logger) (*NumberRange, error) {
	s := &SDC{}
	nr, err := c.ComplementProductionOrderID(input, l)

	if err != nil {
		return nil, xerrors.Errorf("complement orderID error: %w", err)
	}
	res, err := c.rmq.SessionKeepRequest(nil, c.c.RMQ.QueueToSubFunc()["Headers"], input)
	if err != nil {
		return nil, err
	}
	res.Success()

	err = json.Unmarshal(res.Raw(), s)
	if err != nil {
		return nil, err
	}

	//err = c.IncrementLatestNumber(numRange, l)
	//if err != nil {
	//	return xerrors.Errorf("increment latest number error: %w")
	//}

	subfuncSDC.SubfuncResult = s.SubfuncResult
	subfuncSDC.SubfuncError = s.SubfuncError
	subfuncSDC.Message.Header = s.Message.Header
	return nr, err
}

func getBoolPtr(b bool) *bool {
	return &b
}

func (c *SubFuncComplementer) ComplementProductionOrderID(input *dpfm_api_input_reader.SDC, l *logger.Logger) (*NumberRange, error) {
	rows, err := c.db.Query(
		`SELECT ProductionOrder,
			   ProductionOrderItem,
			   Operations,
			   OperationsItem,
			   OperationID,
			   ConfirmationCountingID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_confirmation_header_data
		WHERE ConfirmationCountingID = (SELECT MAX(ConfirmationCountingID) FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_confirmation_header_data)
		AND (ProductionOrder, ProductionOrderItem, Operations, OperationsItem, OperationID) = ((?, ?, ?, ?, ?));`,
		input.Header.ProductionOrder,
		input.Header.ProductionOrderItem,
		input.Header.Operations,
		input.Header.OperationsItem,
		input.Header.OperationID,
	)
	if err != nil {
		return nil, xerrors.Errorf("DB Query error: %w", err)
	}
	nr := NumberRange{}
	if !rows.Next() {
		input.Header.ConfirmationCountingID = 1
		//return nil, xerrors.Errorf("number range does not exist")
		return &nr, nil
	}
	err = rows.Scan(
		&nr.ProductionOrder,
		&nr.ProductionOrderItem,
		&nr.Operations,
		&nr.OperationsItem,
		&nr.OperationID,
		&nr.ConfirmationCountingID,
	)
	if err != nil {
		return nil, xerrors.Errorf("DB Scan error: %w", err)
	}

	nr.ConfirmationCountingID++
	input.Header.ConfirmationCountingID = nr.ConfirmationCountingID
	return &nr, nil
}

func (c *SubFuncComplementer) IncrementLatestNumber(nr *NumberRange, l *logger.Logger) error {
	//_, err := c.db.Query(
	//	`UPDATE DataPlatformCommonSettingsMysqlKube.data_platform_number_range_latest_number_data
	//	SET LatestNumber = ?
	//	WHERE  (NumberRangeID, ServiceLabel, FieldNameWithNumberRange) = ((?,?,?));`, nr.LatestNumber, nr.NumberRangeID, nr.ServiceLabel, nr.FieldNameWithNumberRange,
	//)
	//if err != nil {
	//	return xerrors.Errorf("DB Query error: %w", err)
	//}

	return nil
}
