package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)
// templateを用いて実装する予定
func handleClockNow(w http.ResponseWriter, r *http.Request) {
	// 現在時刻をHTMLで出力
	fmt.Fprintf(w, `
    <!DOCTYPE html>
    <html>
		<body>
        %s
    	</body>
	</html>
    `, time.Now().Format("2006-01-02 Monday 15:04:05"))
	/*
	Format
	2006   : stdLongYear
	01     : stdZeroMonth
	02     : stdZeroDay
	Monday : stdLongWeekDay
	15     : stdHour
	04     : stdZeroMinute
	05     : stfZeroSecond
 	*/
}
func main() {
	// time にアクセスした際に処理するハンドラーを登録
	http.HandleFunc("/time", handleClockNow)

	// サーバーをポート8000で起動
	log.Fatal(http.ListenAndServe(":8000", nil))
}
