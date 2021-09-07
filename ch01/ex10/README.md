wikipediaの長いページランキングを見て一番上をfetchしてみる
https://ja.wikipedia.org/wiki/%E7%89%B9%E5%88%A5:%E9%95%B7%E3%81%84%E3%83%9A%E3%83%BC%E3%82%B8

一番長いページは「英雄伝説 軌跡シリーズの登場人物」だった
https://ja.wikipedia.org/wiki/%E8%8B%B1%E9%9B%84%E4%BC%9D%E8%AA%AC_%E8%BB%8C%E8%B7%A1%E3%82%B7%E3%83%AA%E3%83%BC%E3%82%BA%E3%81%AE%E7%99%BB%E5%A0%B4%E4%BA%BA%E7%89%A9

二回続けて実行したが報告される時間に大きな変化はなかった。
ファイルに差分はなかった。

```shell
$ go run ./main.go "https://ja.wikipedia.org/wiki/英雄伝説_軌跡シリーズの登場人物"
0.87s 2725724 https://ja.wikipedia.org/wiki/英雄伝説_軌跡シリーズの登場人物
0.87s elapsed

$ go run ./main.go "https://ja.wikipedia.org/wiki/英雄伝説_軌跡シリーズの登場人物"
0.86s 2725724 https://ja.wikipedia.org/wiki/英雄伝説_軌跡シリーズの登場人物
0.86s elapsed

$ du -sh *                                                                                                                                                          ✘ 1 
2.6M 0
2.6M 1
4.0K README.md
4.0K main.go

$ diff 0 1
```
