FROM golang:latest as builder
LABEL authors="SvenDowideit@home.org.au"
LABEL mail="elijahxb@163.com"
EXPOSE 80
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
WORKDIR /app
ADD ../go.mod .
ADD ../go.sum .
RUN go mod download
FROM scratch as prod
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/release/gin_demo /
COPY --from=build /go/release/conf ./conf
CMD ["./gin_demo"]