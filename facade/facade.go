// giả sử 1 thư viện bao gồm quá nhiều chức năng
// Tuy nhiêu hệ thống của mình chỉ cần 1 vài chức năng của thử viện
// Cách tối ưu là nên tạo 1 lớp facade để triển khai 1 vài chức năng đó
// Tách biệt code của hệ thống khỏi thư viên , dễ dàng đọc hiểu hơn

package main

import "fmt"

// 1 ví dụ không phổ hiến như sau
// Giả sử hệ thống chỉ cần in ra màn hinh mà không cần tính năng format, scan thì ta có thể facade lại pakage fmt
type PrinterFacadeOfFmt struct{}

func (s *PrinterFacadeOfFmt) Show(content string) {
	fmt.Print(content)
}
