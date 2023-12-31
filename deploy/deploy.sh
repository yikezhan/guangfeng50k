echo "拉取代码开始"
git pull origin main
echo "打包主程序"
go build -o main ./cmd/main.go
echo "杀死原有程序"
kill -9 "$(pgrep -f main)"
echo "主程序权限"
chmod 777 main
echo "运行新打包的程序"
nohup ./main > start.log 2>&1 &
echo "重启结束"
echo "OK"
