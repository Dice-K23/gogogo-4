package main
import(
	"net/http"
	"io/ioutil"
	"html"
	"fmt"
)
const saveFile = "memo.txt" // データファイルの保存先
func main(){
	// サーバーを起動する --- (*1)
	print("memo server - [URL] http://localhost:8888/\n")
	http.HandleFunc("/", readHandler) // ハンドラを登録
	http.HandleFunc("/w", writeHandler)
	http.ListenAndServe(":8888", nil) // 起動
}
// ルートへアクセスしたときメモを表示 --- (*2)
func readHandler(w http.ResponseWriter, r *http.Request){
	// データファイルを開く
	text, err := ioutil.ReadFile(saveFile)
	if err != nil { text = []byte("メモを記入\n")}
	// HTMLのフォームを返す
	htmlText := html.EscapeString(string(text))
	s := "<html>" +
		 "<style>textarea{width:99%; height:200px; }</style>" +
		 "<form method='POST' action='/w'>" +
		 "<textarea name='text'>" + htmlText + "</textarea>" +
		 "<input type='submit' value='保存' /></form><html>";
	w.Write([]byte(s))
}

// フォーム投稿したとき --- (*3)
func writeHandler(w http.ResponseWriter, r *http.Request){
	// 投稿されたフォームを解析
	r.ParseForm() 
	/*ParseFormメソッド 
		--- 投稿されたフォームから任意のフィールドを取り込むのに必要。
		実行すると「Form[フィールド名]」でフォーム内の値を取り出すことができる。
		取り出した値はstring型のスライス、
		そのフィールドが見当たらない場合は要素数は0となる
	*/
	if (len(r.Form["text"])==0){
		w.Write([]byte("フォームから投稿してね。"))
		return
	}
	text := r.Form["text"][0]
	// データファイルへ書き込む
	ioutil.WriteFile(saveFile, []byte(text), 0644)
	fmt.Println("save:" + text)
	// ルートパージへリダイレクトして戻る --- (*4)
	/*リダイレクト
		テキストを保存した後に、テキストを読み込んで表示するルートページへ
		画面遷移するようにすることでテキストファイルの内容が画面に表示される。
		イメージ図: https://news.mynavi.jp/techplus/article/gogogo-4/images/004.jpg
	*/
	http.Redirect(w, r, "/", 301)
	/*Redirectメソッドの引数について*/
}