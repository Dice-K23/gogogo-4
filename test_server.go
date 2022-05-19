// 学習教材: https://news.mynavi.jp/techplus/article/gogogo-4/
package main
import "net/http"
func main(){
	// URLに対応するハンドラを登録 --- (*1)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/aaa", aaaHandler)
	http.HandleFunc("/bbb", bbbHandler)
	http.HandleFunc("/ccc", cccHandler)
	// サーバーを起動 --- (*2)
	http.ListenAndServe(":8888", nil)
}
// ハンドラを定義 --- (*3)
func rootHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello")) // Helloと出力
}
func aaaHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("aaa")) // aaaと出力
}
func bbbHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("bbb")) // bbbと出力
}
func cccHandler(ccc http.ResponseWriter, r *http.Request){
	ccc.Write([]byte("ccc")) // cccと出力
}