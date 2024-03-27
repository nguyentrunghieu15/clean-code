// Mục đích là giảm tài nguyên cần thiết cho cuhuuwogn trình

// Giả sử mỗi một vật thể đều có một màu sắc
// Mỗi một màu sắc sẽ có red , blue, green tổng sẽ có 255*255*255 mau
// Tuy nhien thì số lượng vật thể là không đếm được, nếu mỗi vật thể đều lưu 3 kiểu int trong bản thân nó thì
// sẽ không có đủ ram

// Áp dụng flyweight để tích kiệm ram thì tha sẽ tạo class là Color và mỗi vật thể chỉ
// lưu trữ tham chiếu đến màu của nó

package main

import "fmt"

type Color struct {
	Red   int
	Green int
	Blue  int
}

type Thing struct {
	Color *Color
}

func main() {
	redColor := &Color{255, 0, 0}
	greenColor := &Color{0, 255, 0}
	blueColor := &Color{0, 0, 255}

	n1 := Thing{Color: redColor}
	n2 := Thing{Color: redColor}
	n3 := Thing{Color: greenColor}
	n4 := Thing{Color: blueColor}

	fmt.Println(n1, n2, n3, n4)
}
