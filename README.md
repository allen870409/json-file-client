Go言語育成カリキュラムのJSON FILE SERVERのCLIENT

JSON FILE SERVERは：https://ghe.ca-tools.org/qiao-yicheng/json-file-server


表示 go run main.go get /{path}.json
例： go run main.go get /test/test1.json

ファイルへ保存 go run main.go get -o={出力先} /{path}.json
例： go run main.go get -o=..\file\testclient1.json /test/test1.json

作成 go run main.go create /{path}.json ${json}
例： go run main.go create /test/test_create1.json {\"Id\":-1,\"Name\":\"test\",\"Completed\":true}

ファイルから作成 go run main.go create -i={入力先} /{path}.json 
例： go run main.go create -i=..\file\test\test_create.json /test/test_create2.json

削除 go run main.go delete /{path}.json
例： go run main.go delete /test/test_create.json

更新 go run main.go update /{path}.json ${json}
例： go run main.go update /test/test_create2.json {\"Id\":2,\"Name\":\"test_update!!!!!!!!!!!!!\",\"Completed\":true}

ファイルから更新 go run main.go update -i={入力先} /{path}.json ${json}
例： go run main.go update /test/test_create2.json ..\file\test_update.json

パス一覧表示 go run main.go list /{path} -limit={最大取得数}
例： go run main.go list /test/ -limit=2



