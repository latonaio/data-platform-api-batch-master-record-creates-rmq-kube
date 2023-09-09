package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-production-order-confirmation-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) headerProductExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	headers := make([]dpfm_api_input_reader.Header, 0, 1)
	headers = append(headers, input.Header)
	for _, header := range headers {
		product := getheaderProductMasterGeneralExistenceConfKey(mapper, &header, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(product) {
				wg2.Done()
				return
			}
			res, err := c.productMasterGeneralExistenceConfRequest(product, mapper, input, existenceMap, mtx, log)
			if err != nil {
				mtx.Lock()
				*errs = append(*errs, err)
				mtx.Unlock()
			}
			if res != "" {
				*exconfErrMsg = res
			}
			wg2.Done()
		}()
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

//func (c *ExistenceConf) itemProductExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//
//	items := input.Header.Item
//	for _, item := range items {
//		product := getitemProductMasterGeneralExistenceConfKey(mapper, &item, exconfErrMsg)
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(product) {
//				wg2.Done()
//				return
//			}
//			res, err := c.productMasterGeneralExistenceConfRequest(product, mapper, input, existenceMap, mtx, log)
//			if err != nil {
//				mtx.Lock()
//				*errs = append(*errs, err)
//				mtx.Unlock()
//			}
//			if res != "" {
//				*exconfErrMsg = res
//			}
//			wg2.Done()
//		}()
//	}
//	wg2.Wait()
//	if exReqTimes == 0 {
//		*existenceMap = append(*existenceMap, false)
//	}
//}

func (c *ExistenceConf) productMasterGeneralExistenceConfRequest(product string, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"ProductMasterGeneral": product,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[Returns](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	req.ProductMasterReturn.General.Product = product
	req.Accepter = []string{"General"}

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}

	return "", nil
}

func getheaderProductMasterGeneralExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, exconfErrMsg *string) string {
	var product string

	//switch mapper.Field {
	//case "Product":
	//	if header.Product == nil {
	//		product = ""
	//	} else {
	//		product = *header.Product
	//	}
	//}
	return product
}

//func getitemProductMasterGeneralExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) string {
//	var product string
//
//	switch mapper.Field {
//	case "Product":
//		if item.ComponentProduct == nil {
//			product = ""
//		} else {
//			product = *item.ComponentProduct
//		}
//	}
//	return product
//}

func productMasterConfKeyExistence(res map[string]interface{}, tableTag string) bool {
	req, err := jsonTypeConversion[Returns](res)
	if err != nil {
		return false
	}

	if tableTag == "ProductMasterGeneral" {
		return req.ProductMasterReturn.General.ExistenceConf
	}

	return false
}
