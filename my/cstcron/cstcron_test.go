package cstcron

import "testing"

// TODO[Dokiy] 2022/3/1: to be continued
// 1 以每秒创建一个二维数组，每个元素（子数组）保存即将执行的配置组id
// 2 使用一个参数map来保存配置数组id对应的入参
// 3 配置schedule执行的方法
// 4 每秒获取二维数组里面的配置组id，并进行遍历，将参数map里的参数取出，执行配置的schedule执行方法。
//   （设置gap，设置批量处理方法）
// 5 添加配置，需要将参数添加到map，获取当天剩余所有时间的所有点，换算成秒，添加到二维数组中
//   删除配置，需要将参数从map删除，获取当天剩余所有时间的所有点，换算成秒，从二维数组中删除
// 6 根据初始化方法，设置初始的map参数和二维数组
func BenchmarkName2(b *testing.B) {
	var m = make([][]int, 86400)
	for i := 0; i < b.N; i++ {
		//for _, v := range m {
		//	for _, _ = range v {
		//
		//	}
		//}
		for i := 0; i < 86400; i++ {
			if i%5 == 0 {
				m[i] = make([]int, 1000)
			} else {
				m[i] = nil
			}
		}
	}
}