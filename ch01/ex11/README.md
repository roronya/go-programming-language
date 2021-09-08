alexa.com/topsiteから50個のURLを取得し、vimで適当にかさ増して、5194個のURLのリストを作った

```shell
go run ./main.go $(cat sitelist.txt) > log
```

以上のように実行し眺めた。
ソフトウェアは5194回のfetchを行い、91回は正常にfetchを終えたが、残りの5083回はエラーになった。

## エラーの種類と頻度
エラーの種類は以下のような2種類が発生した。

Get "http://Csdn.net": EOF
Get "http://Stackoverflow.com": read tcp 127.0.0.1:55507->151.101.1.69:80: read: connection reset by peer
Get "http://Twitch.tv": dial tcp 151.101.2.167:80: i/o timeout

それぞれEOFが16回、connection reset by peerが250回、timeout 4817回発生した。

## 時系列での振る舞い
ソフトウェアを動作開始させた直後は正常に動作を終えるものとconnection reset by peerが入り混じった状態になった。
その後は大量にtimeoutが発生しほとんど終盤までそれが続いた。このときたまにEOFも発生している。
終盤はまた正常に動作を終えるものとconnection reset by peerが入り混じった状態になった。
