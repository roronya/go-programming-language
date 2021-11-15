ヘルプメッセージの表示にcelsiusFlagのString()メソッドを使うから。
https://github.com/golang/go/blob/master/src/flag/flag.go#L534

celsiusFlagのString()はCelsiusのString()を呼び出していて、CelsiusのString()がvalueに°Cをくっつけたstringを返している。