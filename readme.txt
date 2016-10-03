
１．golang ウィンドウズ版をインストール

https://golang.org/dl/


２．goenv.bat を実行

　GOROOT に golang をインストールしたディレクトリを設定
　GOPATH に このでディレクトリを設定
　HTTP_PROXY, HTTPS_PROXY でプロキシ設定 
　
  goenv.bat の set GOROOT=, set GOPATH=, 
  set HTTP_PROXY= set HTTPS_PROXY= を書き換えて実行してください。
  
３．設定の修正

　温度の条件を変更する場合は、
　service.go の
  const (
    // 警告温度
    warningLimit  = 33.5
    // エラー温度
    emergeLimit   = 35.0
  )
　を修正します。
　
　メール情報を変更する場合は
　sendmail.go
　関数 sendMail の定数を書き換えます
  func sendMail(name string, title string, description string) error {
    from     := "daemon@foo dot com"
    to       := "tweet@foo dot com"
    user     := "smtp user@foo dot com"
    password := "your password"
    subject  := title
    server   := "foo dot com"
    port     := "587"
　

４．build.bat を実行

  HnsTemp.exe が生成されます。
  
５．インストール

  HnsTemp.exe と USBMeter.DLL をコピーして
  
  HnsTemp install を実行します。
  サービスとして登録されます。
  HnsTempサービスを開いて自動実行に変更します。
  
  HnsTemp remove でサービスから削除されます。
  


補足：

USBMeter.dll は、下記のキットです。
http://strawberry-linux.com/catalog/items?code=52001
