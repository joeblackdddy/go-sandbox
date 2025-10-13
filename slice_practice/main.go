package main

import "fmt"
import "unsafe"
import "reflect"

func main() {
	fmt.Println("Hello, World!")

	var names []string
	names = []string{"Jiro"}

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&names))
	fmt.Printf("初期状態: %v, len: %d, cap: %d, data: %p, slice変数アドレス: %p\n", 
		names, len(names), cap(names), unsafe.Pointer(sliceHeader.Data), &names)
	
	// 既存要素の変更
	addName(names, "John")
	sliceHeader = (*reflect.SliceHeader)(unsafe.Pointer(&names))
	fmt.Printf("addName後: %v, len: %d, cap: %d, data: %p, slice変数アドレス: %p\n", 
		names, len(names), cap(names), unsafe.Pointer(sliceHeader.Data), &names)

	// appendの使用
	addNameByAppend(names, "John")
	fmt.Printf("addNameByAppend後: %v, len: %d, cap: %d, data: %p\n", 
		names, len(names), cap(names), (*unsafe.Pointer)(unsafe.Pointer(&names)))
	
	// ポインタを返すバージョン
	names = addNameByAppendReturn(names, "Jane")
	fmt.Printf("addNameByAppendReturn後: %v, len: %d, cap: %d, data: %p\n", 
		names, len(names), cap(names), (*unsafe.Pointer)(unsafe.Pointer(&names)))
}

func addName(names []string, name string) {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&names))
	fmt.Printf("関数内 addName（変更前）: %v, len: %d, cap: %d, data: %p, slice変数アドレス: %p\n", 
		names, len(names), cap(names), unsafe.Pointer(sliceHeader.Data), &names)
	names[len(names)-1] = name
	sliceHeader = (*reflect.SliceHeader)(unsafe.Pointer(&names))
	fmt.Printf("関数内 addName（変更後）: %v, len: %d, cap: %d, data: %p, slice変数アドレス: %p\n", 
		names, len(names), cap(names), unsafe.Pointer(sliceHeader.Data), &names)
}

func addNameByAppend(names []string, name string) {
	names = append(names, name)
	fmt.Printf("関数内 addNameByAppend: %v, len: %d, cap: %d, data: %p\n", 
		names, len(names), cap(names), (*unsafe.Pointer)(unsafe.Pointer(&names)))
}

func addNameByAppendReturn(names []string, name string) []string {
	names = append(names, name)
	fmt.Printf("関数内 addNameByAppendReturn: %v, len: %d, cap: %d, data: %p\n", 
		names, len(names), cap(names), (*unsafe.Pointer)(unsafe.Pointer(&names)))
	return names
}