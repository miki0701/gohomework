# start server
docker run -d -p 8080:80 cadmusgo/geektime_cloudnative_hw_2

# test
curl -i -H "user: 123" -H "x-forwarded-for:192.168.100.100" http://127.0.0.1:8080/cloneheader