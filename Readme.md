go get -u github.com/gin-gonic/gin

go fmt ./...
go run main.go

    curl -v -X POST http://localhost:8080/create-matrix -H 'content-type: application/json'   -d '{ "range": 5 }'
    curl -v -X POST http://localhost:8080/ship -H 'content-type: application/json'   -d '{ "Coordinates": "1A 2B,3D 3E" }'
    curl -v -X POST http://localhost:8080/shot -H 'content-type: application/json'   -d '{ "coord": "1A" }'
    curl -v -X POST http://localhost:8080/shot -H 'content-type: application/json'   -d '{ "coord": "2B" }'
    curl -v http://localhost:8080/state
    curl -v -X POST http://localhost:8080/clear

+ проверить почему не выставляется knock и destroy на /shot - 2B
+ проверить почему не увеличивается knock и destroy у /state
- проверки:
  + проверка на повторный выстрел
  + на выход кораблей за пределы матрицы
  - корабль где LeftTop меньше RightBottom
- рефакторинг:
  + разнести handlers в папку
  + попробовать вынести логику в сервисы
  - тесты?