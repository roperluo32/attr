package report

import "log"

// Report 上报函数
func Report(name string) error {
	log.Printf("report name:%v", name)

	return nil
}
