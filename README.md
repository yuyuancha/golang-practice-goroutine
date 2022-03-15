# golang-practice-goroutine

## 簡述

透過 `go` 關鍵字，來了解 `goroutine` 的運作方式，可以多工執行節省執行時間。

## 步驟及說明

* 紀錄程式開始時間。
* 透過 `go` 關鍵字執行多個隨機等待 0-2 秒的函式，並搭配 `sync.WaitGroup` 的 `.Wait()` 來使程式等待執行完畢。
* 紀錄程式結束時間，並計算花費時間。

> 藉此了解 `goroutine` 的執行方式，可用於節省執行時間。

## 安裝 docker 環境及執行程式

* clone GitHub repository
```
$ get clone https://github.com/yuyuancha/golang-practice-goroutine.git
```

* 透過 docker-compose 開啟 docker 容器
```
$ docker-compose up -d
```

* 執行 main.go
```
$ docker-compose exec app go run main.go
```

* 關閉 docker 容器
```
docker-compose down
```
