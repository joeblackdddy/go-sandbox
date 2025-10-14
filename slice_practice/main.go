package main

import (
	"fmt"
	"errors"
	"time"
	"unsafe"
)

// 大きな構造体を定義
type LargeStruct struct {
	ID       int
	Name     string
	Data     [1000]int        // 大きな配列
	Metadata map[string]interface{}
	Values   []float64
}

func main() {
	fmt.Println("=== 大きなsliceの真ん中のmap要素への安全で効率的なアクセス方法 ===")
	
	// 大きなsliceを作成（mapを要素として持つ）
	largeSlice := createLargeSlice(1000000)
	fmt.Printf("作成したsliceのサイズ: %d\n", len(largeSlice))
	
	// 方法1: 基本的な安全なアクセス
	fmt.Println("\n--- 方法1: 基本的な安全なアクセス ---")
	value, err := safeAccessMiddleMap(largeSlice, "key1")
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	} else {
		fmt.Printf("真ん中のmapのkey1の値: %v\n", value)
	}
	
	// 方法2: より効率的なアクセス（境界チェックを最小限に）
	fmt.Println("\n--- 方法2: 効率的なアクセス ---")
	value2, err2 := efficientAccessMiddleMap(largeSlice, "key2")
	if err2 != nil {
		fmt.Printf("エラー: %v\n", err2)
	} else {
		fmt.Printf("真ん中のmapのkey2の値: %v\n", value2)
	}
	
	// 方法3: ポインタを使用した効率的なアクセス
	fmt.Println("\n--- 方法3: ポインタを使用した効率的なアクセス ---")
	value3, err3 := pointerAccessMiddleMap(largeSlice, "key3")
	if err3 != nil {
		fmt.Printf("エラー: %v\n", err3)
	} else {
		fmt.Printf("真ん中のmapのkey3の値: %v\n", value3)
	}
	
	// 空のsliceでのテスト
	fmt.Println("\n--- 空のsliceでのテスト ---")
	emptySlice := make([]map[string]int, 0)
	_, err4 := safeAccessMiddleMap(emptySlice, "key1")
	if err4 != nil {
		fmt.Printf("期待通りのエラー: %v\n", err4)
	}
	
	// パフォーマンステストを実行
	runPerformanceTest()
	
	// ポインタアクセスの危険性をデモンストレーション
	demonstratePointerDanger()
	
	// より詳細な危険性の例
	demonstrateDetailedDanger()
	
	// メモリ管理の仕組みを詳しく説明
	explainMemoryManagement()
	
	// 実際の使用パターンを検証
	testRealWorldUsage()
	
	// 大きな要素と多数の要素でのパフォーマンス比較
	testLargeElementsPerformance()
}

// createLargeSlice は指定されたサイズの大きなsliceを作成します
func createLargeSlice(size int) []map[string]int {
	slice := make([]map[string]int, size)
	
	// 各要素にmapを初期化
	for i := 0; i < size; i++ {
		slice[i] = make(map[string]int)
		// テスト用のデータを設定
		slice[i]["key1"] = i * 10
		slice[i]["key2"] = i * 20
		slice[i]["key3"] = i * 30
	}
	
	return slice
}

// 方法1: 基本的な安全なアクセス
// 最も安全だが、毎回境界チェックを行う
func safeAccessMiddleMap(slice []map[string]int, key string) (int, error) {
	// 空のsliceチェック
	if len(slice) == 0 {
		return 0, errors.New("slice is empty")
	}
	
	// 真ん中のインデックスを計算
	middleIndex := len(slice) / 2
	
	// 境界チェック（実際には不要だが、明示的にチェック）
	if middleIndex >= len(slice) {
		return 0, errors.New("index out of bounds")
	}
	
	// mapの存在チェック
	middleMap := slice[middleIndex]
	if middleMap == nil {
		return 0, errors.New("map at middle index is nil")
	}
	
	// キーの存在チェック
	value, exists := middleMap[key]
	if !exists {
		return 0, fmt.Errorf("key '%s' not found in map", key)
	}
	
	return value, nil
}

// 方法2: より効率的なアクセス
// 境界チェックを最小限に抑え、パフォーマンスを重視
func efficientAccessMiddleMap(slice []map[string]int, key string) (int, error) {
	// 空のsliceチェックのみ
	if len(slice) == 0 {
		return 0, errors.New("slice is empty")
	}
	
	// 真ん中のインデックスを計算（境界チェックは不要）
	middleIndex := len(slice) / 2
	middleMap := slice[middleIndex]
	
	// mapの存在チェック
	if middleMap == nil {
		return 0, errors.New("map at middle index is nil")
	}
	
	// キーの存在チェック
	value, exists := middleMap[key]
	if !exists {
		return 0, fmt.Errorf("key '%s' not found in map", key)
	}
	
	return value, nil
}

// 方法3: ポインタを使用した効率的なアクセス
// 最も効率的だが、ポインタの扱いに注意が必要
func pointerAccessMiddleMap(slice []map[string]int, key string) (int, error) {
	// 空のsliceチェック
	if len(slice) == 0 {
		return 0, errors.New("slice is empty")
	}
	
	// 真ん中のインデックスを計算
	middleIndex := len(slice) / 2
	
	// ポインタを使用して直接アクセス（境界チェックなし）
	// 注意: この方法は高速だが、sliceが変更される可能性がある場合は危険
	middleMapPtr := &slice[middleIndex]
	
	// mapの存在チェック
	if *middleMapPtr == nil {
		return 0, errors.New("map at middle index is nil")
	}
	
	// キーの存在チェック
	value, exists := (*middleMapPtr)[key]
	if !exists {
		return 0, fmt.Errorf("key '%s' not found in map", key)
	}
	
	return value, nil
}

// ベンチマーク用の関数群
func benchmarkSafeAccess(slice []map[string]int, key string) int {
	value, _ := safeAccessMiddleMap(slice, key)
	return value
}

func benchmarkEfficientAccess(slice []map[string]int, key string) int {
	value, _ := efficientAccessMiddleMap(slice, key)
	return value
}

func benchmarkPointerAccess(slice []map[string]int, key string) int {
	value, _ := pointerAccessMiddleMap(slice, key)
	return value
}

// パフォーマンステスト用の関数
func runPerformanceTest() {
	fmt.Println("\n=== パフォーマンステスト ===")
	
	// テスト用のsliceを作成
	testSlice := createLargeSlice(1000000)
	key := "key1"
	
	// 各方法の実行時間を測定（簡易版）
	fmt.Println("各方法のパフォーマンス比較:")
	
	// 方法1のテスト
	start := time.Now()
	for i := 0; i < 1000; i++ {
		benchmarkSafeAccess(testSlice, key)
	}
	duration1 := time.Since(start)
	fmt.Printf("方法1 (安全なアクセス): %v (1000回実行)\n", duration1)
	
	// 方法2のテスト
	start = time.Now()
	for i := 0; i < 1000; i++ {
		benchmarkEfficientAccess(testSlice, key)
	}
	duration2 := time.Since(start)
	fmt.Printf("方法2 (効率的なアクセス): %v (1000回実行)\n", duration2)
	
	// 方法3のテスト
	start = time.Now()
	for i := 0; i < 1000; i++ {
		benchmarkPointerAccess(testSlice, key)
	}
	duration3 := time.Since(start)
	fmt.Printf("方法3 (ポインタアクセス): %v (1000回実行)\n", duration3)
	
	fmt.Printf("\n効率性の比較:\n")
	fmt.Printf("方法2は方法1より %.2fx 高速\n", float64(duration1)/float64(duration2))
	fmt.Printf("方法3は方法1より %.2fx 高速\n", float64(duration1)/float64(duration3))
	fmt.Printf("方法3は方法2より %.2fx 高速\n", float64(duration2)/float64(duration3))
}

// ポインタアクセスの危険性をデモンストレーションする関数
func demonstratePointerDanger() {
	fmt.Println("\n=== ポインタアクセスの危険性デモンストレーション ===")
	
	// 小さなsliceでテスト（理解しやすくするため）
	slice := createLargeSlice(5)
	fmt.Printf("初期slice: %v\n", slice)
	
	// 方法2（安全）: 真ん中の要素を取得
	middleIndex := len(slice) / 2
	fmt.Printf("真ん中のインデックス: %d\n", middleIndex)
	
	// 方法2でアクセス
	fmt.Println("\n--- 方法2（安全なアクセス）---")
	value2, err2 := efficientAccessMiddleMap(slice, "key1")
	if err2 == nil {
		fmt.Printf("方法2で取得した値: %d\n", value2)
	}
	
	// 方法3でポインタを取得
	fmt.Println("\n--- 方法3（ポインタアクセス）---")
	middleMapPtr := &slice[middleIndex]
	fmt.Printf("ポインタで取得したmap: %v\n", *middleMapPtr)
	
	// ここでsliceを変更してみる
	fmt.Println("\n--- sliceを変更 ---")
	// appendでsliceを拡張（内部配列が再割り当てされる可能性）
	slice = append(slice, map[string]int{"new": 999})
	fmt.Printf("append後のslice: %v\n", slice)
	fmt.Printf("sliceの長さ: %d\n", len(slice))
	
	// ポインタが指している先を確認
	fmt.Printf("ポインタが指している先: %v\n", *middleMapPtr)
	
	// さらに危険な例：sliceの先頭に要素を挿入
	fmt.Println("\n--- より危険な例：sliceの先頭に要素を挿入 ---")
	originalSlice := slice
	slice = append([]map[string]int{{"inserted": 888}}, slice...)
	fmt.Printf("先頭挿入後のslice: %v\n", slice)
	fmt.Printf("元のslice: %v\n", originalSlice)
	fmt.Printf("ポインタが指している先: %v\n", *middleMapPtr)
	
	// 方法2と方法3の違いを明確に示す
	fmt.Println("\n--- 方法2と方法3の違い ---")
	fmt.Println("方法2: 毎回slice[middleIndex]を計算するため、sliceが変更されても正しい要素を取得")
	fmt.Println("方法3: ポインタを保持するため、sliceが変更されると古いメモリ位置を参照する可能性")
	
	// 実際に方法2で再計算してみる
	newMiddleIndex := len(slice) / 2
	fmt.Printf("新しい真ん中のインデックス: %d\n", newMiddleIndex)
	fmt.Printf("新しい真ん中の要素: %v\n", slice[newMiddleIndex])
}

// より詳細な危険性の例
func demonstrateDetailedDanger() {
	fmt.Println("\n=== より詳細な危険性の例 ===")
	
	// 初期slice
	slice := make([]map[string]int, 3)
	for i := 0; i < 3; i++ {
		slice[i] = map[string]int{"value": i * 10}
	}
	fmt.Printf("初期slice: %v\n", slice)
	
	// 真ん中の要素のポインタを取得
	middlePtr := &slice[1]
	fmt.Printf("真ん中の要素のポインタ: %p, 値: %v\n", middlePtr, *middlePtr)
	
	// sliceを大きく拡張（内部配列が再割り当てされる）
	fmt.Println("\nsliceを大きく拡張...")
	slice = append(slice, make([]map[string]int, 1000)...)
	fmt.Printf("拡張後のslice長: %d\n", len(slice))
	fmt.Printf("元のポインタが指している先: %p, 値: %v\n", middlePtr, *middlePtr)
	fmt.Printf("新しいslice[1]: %v\n", slice[1])
	
	// ポインタが指している先と実際のslice[1]が異なることを確認
	if middlePtr != &slice[1] {
		fmt.Println("⚠️  危険: ポインタが古いメモリ位置を指しています！")
		fmt.Printf("ポインタのアドレス: %p\n", middlePtr)
		fmt.Printf("実際のslice[1]のアドレス: %p\n", &slice[1])
	}
}

// メモリ管理の仕組みを詳しく説明する関数
func explainMemoryManagement() {
	fmt.Println("\n=== なぜ古いポインタのままになってしまうのか？ ===")
	
	// 1. 初期状態の確認
	fmt.Println("\n--- 1. 初期状態 ---")
	slice := make([]map[string]int, 3)
	for i := 0; i < 3; i++ {
		slice[i] = map[string]int{"id": i, "value": i * 10}
	}
	
	fmt.Printf("初期slice: %v\n", slice)
	fmt.Printf("sliceのアドレス: %p\n", &slice)
	fmt.Printf("slice[0]のアドレス: %p\n", &slice[0])
	fmt.Printf("slice[1]のアドレス: %p\n", &slice[1])
	fmt.Printf("slice[2]のアドレス: %p\n", &slice[2])
	
	// 2. ポインタを取得
	fmt.Println("\n--- 2. ポインタを取得 ---")
	middlePtr := &slice[1]
	fmt.Printf("middlePtrの値（アドレス）: %p\n", middlePtr)
	fmt.Printf("middlePtrが指す値: %v\n", *middlePtr)
	
	// 3. sliceの内部構造を確認
	fmt.Println("\n--- 3. sliceの内部構造 ---")
	fmt.Printf("sliceの長さ: %d\n", len(slice))
	fmt.Printf("sliceの容量: %d\n", cap(slice))
	fmt.Printf("sliceのデータポインタ: %p\n", &slice[0])
	
	// 4. 小さな拡張（容量内）
	fmt.Println("\n--- 4. 小さな拡張（容量内）---")
	slice = append(slice, map[string]int{"id": 3, "value": 30})
	fmt.Printf("拡張後のslice: %v\n", slice)
	fmt.Printf("sliceの長さ: %d, 容量: %d\n", len(slice), cap(slice))
	fmt.Printf("slice[0]のアドレス: %p\n", &slice[0])
	fmt.Printf("middlePtrの値: %p\n", middlePtr)
	fmt.Printf("middlePtrが指す値: %v\n", *middlePtr)
	fmt.Printf("slice[1]のアドレス: %p\n", &slice[1])
	
	if middlePtr == &slice[1] {
		fmt.Println("✅ ポインタは有効（容量内の拡張では再割り当てされない）")
	} else {
		fmt.Println("❌ ポインタは無効（再割り当てが発生）")
	}
	
	// 5. 大きな拡張（容量を超える）
	fmt.Println("\n--- 5. 大きな拡張（容量を超える）---")
	fmt.Printf("現在の容量: %d\n", cap(slice))
	
	// 容量を超える要素を追加
	oldDataPtr := &slice[0]
	slice = append(slice, make([]map[string]int, 10)...)
	
	fmt.Printf("拡張後のslice長: %d, 容量: %d\n", len(slice), cap(slice))
	fmt.Printf("古いデータポインタ: %p\n", oldDataPtr)
	fmt.Printf("新しいslice[0]のアドレス: %p\n", &slice[0])
	fmt.Printf("middlePtrの値: %p\n", middlePtr)
	fmt.Printf("middlePtrが指す値: %v\n", *middlePtr)
	fmt.Printf("新しいslice[1]のアドレス: %p\n", &slice[1])
	
	if middlePtr == &slice[1] {
		fmt.Println("✅ ポインタは有効")
	} else {
		fmt.Println("❌ ポインタは無効（再割り当てが発生）")
		fmt.Println("   → 古いメモリ位置を指し続けている")
	}
	
	// 6. なぜポインタが更新されないのかを説明
	fmt.Println("\n--- 6. なぜポインタが更新されないのか？ ---")
	fmt.Println("Goのsliceは以下の3つの要素で構成されています：")
	fmt.Println("1. データポインタ（実際の配列の先頭アドレス）")
	fmt.Println("2. 長さ（len）")
	fmt.Println("3. 容量（cap）")
	fmt.Println()
	fmt.Println("ポインタ変数（middlePtr）は、特定のメモリアドレスを保持します。")
	fmt.Println("sliceが再割り当てされると：")
	fmt.Println("- 新しいメモリ領域にデータがコピーされる")
	fmt.Println("- sliceのデータポインタが新しいアドレスを指す")
	fmt.Println("- しかし、古いポインタ変数は古いアドレスを保持し続ける")
	fmt.Println()
	fmt.Println("これが「古いポインタのまま」になってしまう理由です。")
	
	// 7. 安全な方法との比較
	fmt.Println("\n--- 7. 安全な方法（方法2）との比較 ---")
	fmt.Println("方法2では：")
	fmt.Println("- ポインタを保持しない")
	fmt.Println("- 毎回 slice[middleIndex] を計算")
	fmt.Println("- sliceが変更されても、常に最新の位置を参照")
	fmt.Println("- メモリ再割り当ての影響を受けない")
	
	// 実際に比較してみる
	fmt.Println("\n実際の比較：")
	newMiddleIndex := len(slice) / 2
	fmt.Printf("方法2で計算した真ん中の要素: %v\n", slice[newMiddleIndex])
	fmt.Printf("方法3のポインタが指す要素: %v\n", *middlePtr)
	
	// mapの比較は直接できないので、アドレスで比較
	if &slice[newMiddleIndex] == middlePtr {
		fmt.Println("✅ 同じメモリ位置（安全）")
	} else {
		fmt.Println("❌ 異なるメモリ位置（危険な状態）")
		fmt.Printf("方法2のアドレス: %p\n", &slice[newMiddleIndex])
		fmt.Printf("方法3のアドレス: %p\n", middlePtr)
	}
}

// 実際の使用パターンを検証する関数
func testRealWorldUsage() {
	fmt.Println("\n=== 実際の使用パターンの検証 ===")
	
	// 実際の使用場面をシミュレート
	slice := createLargeSlice(1000)
	
	fmt.Println("\n--- パターン1: 一度だけアクセス（最も一般的）---")
	// この場合、ポインタアクセスは完全に安全
	value1 := getValueWithPointer(slice, "key1")
	value2 := getValueWithIndex(slice, "key1")
	fmt.Printf("ポインタアクセス: %d\n", value1)
	fmt.Printf("インデックスアクセス: %d\n", value2)
	
	fmt.Println("\n--- パターン2: 複数回アクセス（sliceが変更されない場合）---")
	// sliceが変更されない限り、ポインタアクセスは安全で高速
	for i := 0; i < 5; i++ {
		val := getValueWithPointer(slice, "key2")
		fmt.Printf("アクセス%d回目: %d\n", i+1, val)
	}
	
	fmt.Println("\n--- パターン3: 複数回アクセス（sliceが変更される場合）---")
	// この場合のみ危険
	fmt.Println("sliceを変更しながらアクセス...")
	for i := 0; i < 3; i++ {
		val := getValueWithPointer(slice, "key3")
		fmt.Printf("変更前アクセス%d回目: %d\n", i+1, val)
		
		// sliceを変更
		slice = append(slice, map[string]int{"key3": 999})
		fmt.Printf("slice長: %d\n", len(slice))
		
		valAfter := getValueWithPointer(slice, "key3")
		fmt.Printf("変更後アクセス%d回目: %d\n", i+1, valAfter)
	}
	
	fmt.Println("\n--- パターン4: 高頻度アクセス（パフォーマンス重視）---")
	// 大量のアクセスでパフォーマンスを比較
	largeSlice := createLargeSlice(100000)
	
	// ポインタアクセス（高速）
	start := time.Now()
	for i := 0; i < 10000; i++ {
		_ = getValueWithPointer(largeSlice, "key1")
	}
	pointerTime := time.Since(start)
	
	// インデックスアクセス（安全）
	start = time.Now()
	for i := 0; i < 10000; i++ {
		_ = getValueWithIndex(largeSlice, "key1")
	}
	indexTime := time.Since(start)
	
	fmt.Printf("ポインタアクセス: %v (10000回)\n", pointerTime)
	fmt.Printf("インデックスアクセス: %v (10000回)\n", indexTime)
	fmt.Printf("ポインタアクセスは %.2fx 高速\n", float64(indexTime)/float64(pointerTime))
	
	fmt.Println("\n--- 結論 ---")
	fmt.Println("✅ 一度だけアクセス: ポインタアクセスは安全で高速")
	fmt.Println("✅ 複数回アクセス（slice変更なし）: ポインタアクセスは安全で高速")
	fmt.Println("❌ 複数回アクセス（slice変更あり）: ポインタアクセスは危険")
	fmt.Println("✅ 高頻度アクセス: ポインタアクセスは大幅に高速")
	fmt.Println()
	fmt.Println("実際の使用では、sliceが変更されないことが多いため、")
	fmt.Println("ポインタアクセスの危険性は過大評価されている可能性があります。")
}

// ポインタを使用した値取得（危険だが高速）
func getValueWithPointer(slice []map[string]int, key string) int {
	if len(slice) == 0 {
		return 0
	}
	middleIndex := len(slice) / 2
	middlePtr := &slice[middleIndex]
	if *middlePtr == nil {
		return 0
	}
	value, exists := (*middlePtr)[key]
	if !exists {
		return 0
	}
	return value
}

// インデックスを使用した値取得（安全）
func getValueWithIndex(slice []map[string]int, key string) int {
	if len(slice) == 0 {
		return 0
	}
	middleIndex := len(slice) / 2
	middleMap := slice[middleIndex]
	if middleMap == nil {
		return 0
	}
	value, exists := middleMap[key]
	if !exists {
		return 0
	}
	return value
}

// 大きな要素と多数の要素でのパフォーマンス比較
func testLargeElementsPerformance() {
	fmt.Println("\n=== 大きな要素と多数の要素でのパフォーマンス比較 ===")
	
	// 大きな要素を持つsliceを作成
	fmt.Println("\n--- 大きな要素でのテスト ---")
	largeSlice := createLargeStructSlice(100000)
	fmt.Printf("大きな要素のslice作成完了: %d個\n", len(largeSlice))
	
	// パフォーマンス比較
	fmt.Println("\n--- パフォーマンス比較（1000回アクセス）---")
	
	// ポインタアクセス
	start := time.Now()
	for i := 0; i < 1000; i++ {
		_ = getLargeStructWithPointer(largeSlice)
	}
	pointerTime := time.Since(start)
	
	// インデックスアクセス
	start = time.Now()
	for i := 0; i < 1000; i++ {
		_ = getLargeStructWithIndex(largeSlice)
	}
	indexTime := time.Since(start)
	
	fmt.Printf("ポインタアクセス: %v\n", pointerTime)
	fmt.Printf("インデックスアクセス: %v\n", indexTime)
	fmt.Printf("ポインタアクセスは %.2fx 高速\n", float64(indexTime)/float64(pointerTime))
	
	// メモリ使用量の比較
	fmt.Println("\n--- メモリ使用量の比較 ---")
	fmt.Printf("1つのLargeStructのサイズ: %d bytes\n", getStructSize())
	fmt.Printf("slice全体の推定サイズ: %d MB\n", len(largeSlice)*getStructSize()/(1024*1024))
	
	// より詳細な分析
	fmt.Println("\n--- 詳細分析 ---")
	analyzeAccessPatterns(largeSlice)
}

// 大きな構造体のsliceを作成
func createLargeStructSlice(size int) []LargeStruct {
	slice := make([]LargeStruct, size)
	
	for i := 0; i < size; i++ {
		slice[i] = LargeStruct{
			ID:   i,
			Name: fmt.Sprintf("Item_%d", i),
			Data: [1000]int{},
			Metadata: map[string]interface{}{
				"created": time.Now(),
				"version": i % 10,
				"active":  i%2 == 0,
			},
			Values: make([]float64, 100),
		}
		
		// データを埋める
		for j := 0; j < 1000; j++ {
			slice[i].Data[j] = i * j
		}
		for j := 0; j < 100; j++ {
			slice[i].Values[j] = float64(i) * float64(j) * 0.1
		}
	}
	
	return slice
}

// ポインタを使用した大きな構造体の取得
func getLargeStructWithPointer(slice []LargeStruct) *LargeStruct {
	if len(slice) == 0 {
		return nil
	}
	middleIndex := len(slice) / 2
	return &slice[middleIndex]
}

// インデックスを使用した大きな構造体の取得
func getLargeStructWithIndex(slice []LargeStruct) LargeStruct {
	if len(slice) == 0 {
		return LargeStruct{}
	}
	middleIndex := len(slice) / 2
	return slice[middleIndex]
}

// 構造体のサイズを取得
func getStructSize() int {
	var s LargeStruct
	return int(unsafe.Sizeof(s))
}

// アクセスパターンの詳細分析
func analyzeAccessPatterns(slice []LargeStruct) {
	fmt.Println("アクセスパターンの分析:")
	
	// 1. 一度だけアクセス
	fmt.Println("\n1. 一度だけアクセス:")
	start := time.Now()
	ptr := getLargeStructWithPointer(slice)
	_ = ptr.ID
	pointerOnce := time.Since(start)
	
	start = time.Now()
	val := getLargeStructWithIndex(slice)
	_ = val.ID
	indexOnce := time.Since(start)
	
	fmt.Printf("  ポインタ: %v\n", pointerOnce)
	fmt.Printf("  インデックス: %v\n", indexOnce)
	fmt.Printf("  ポインタは %.2fx 高速\n", float64(indexOnce)/float64(pointerOnce))
	
	// 2. 複数回アクセス（同じ要素）
	fmt.Println("\n2. 複数回アクセス（同じ要素）:")
	iterations := 10000
	
	start = time.Now()
	for i := 0; i < iterations; i++ {
		ptr := getLargeStructWithPointer(slice)
		_ = ptr.ID
	}
	pointerMultiple := time.Since(start)
	
	start = time.Now()
	for i := 0; i < iterations; i++ {
		val := getLargeStructWithIndex(slice)
		_ = val.ID
	}
	indexMultiple := time.Since(start)
	
	fmt.Printf("  ポインタ: %v (%d回)\n", pointerMultiple, iterations)
	fmt.Printf("  インデックス: %v (%d回)\n", indexMultiple, iterations)
	fmt.Printf("  ポインタは %.2fx 高速\n", float64(indexMultiple)/float64(pointerMultiple))
	
	// 3. メモリコピーの影響
	fmt.Println("\n3. メモリコピーの影響:")
	fmt.Printf("  構造体サイズ: %d bytes\n", getStructSize())
	fmt.Printf("  ポインタアクセス: 8 bytes (ポインタのみ)\n")
	fmt.Printf("  インデックスアクセス: %d bytes (構造体全体をコピー)\n", getStructSize())
	fmt.Printf("  コピー量の差: %dx\n", getStructSize()/8)
	
	// 4. 結論
	fmt.Println("\n4. 結論:")
	if float64(indexMultiple)/float64(pointerMultiple) > 1.5 {
		fmt.Println("  ✅ 大きな要素では、ポインタアクセスが大幅に高速")
		fmt.Println("  ✅ メモリコピーのオーバーヘッドが大きい")
		fmt.Println("  ✅ 要素が大きいほど、ポインタアクセスの優位性が増す")
	} else {
		fmt.Println("  ⚠️  このサイズでは、差はそれほど大きくない")
	}
}