Go言語育成カリキュラムのJSON FILE SERVERのCLIENT

JSON FILE SERVERは：https://ghe.ca-tools.org/qiao-yicheng/json-file-server


表示 go run main.go get /{path}.json
例： go run main.go get /todos/1.json

ファイルへ保存 go run main.go get -o={出力先} /{path}.json
例： go run main.go get -o=..\file\todos1.json /todos/1.json

作成 go run main.go create /{path}.json ${json}
例： go run main.go create /todos.json {\"Id\":-1,\"Name\":\"test\",\"Completed\":true}

ファイルから作成 go run main.go create -i={入力先} /{path}.json 
例： go run main.go create -i=..\file\todos1.json /todos.json 

削除 go run main.go delete /{path}.json
例： go run main.go delete \todos27.json 

更新 go run main.go update /{path}.json ${json}
例： go run main.go update /todos.json {\"Id\":2,\"Name\":\"test update\",\"Completed\":true}

ファイルから更新 go run main.go update -i={入力先} /{path}.json ${json}
例： go run main.go update /todos.json ..\file\todosupdate.json

パス一覧表示 go run main.go list /{path} -limit={最大取得数}
例： go run main.go list /todos.json -limit=2



