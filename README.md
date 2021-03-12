# NewMesos-go
## 개발환경 구축 
- Docker Compose 를 이용한 개발환경 구축 
```shell
$ make start // mesos master, agent, zookeeper, etc.. 
$ make stop  // 컨테이너 중지 및 삭제 
```
- go mod init, tidy, vendor 로 dependency 관리

## Testing
```shell
$ make test
```
- Benchmark
- Test Coverage 
```
go test -cover

// coverage 파일 출력
go test -coverprofile=coverage.out

// coverage 파일 html 로 변환 
go tool cover -html=coverage.out

// gocov과 gocov-html 이용해서 리포트 파일 보기 좋게
// gocov, gocov-html 설치
go get github.com/axw/gocov/gocov
go get github.com/matm/gocov-html

gocov test ./ > handler.json
gocov-html handler.json > handler.html
```
## 진행 상황
RESERVATION_REFINEMENT, CapabilityReservationRefinement=1 과 관련하여 [mesos 매뉴얼](http://mesos.apache.org/documentation/latest/reservation/) 을 참고해야함.
mesos scheduler 개발 시작. mesos executor 의 경우 컨테이너 환경에서는 개발이 어려움. 이 문제에 대해서는 추가적인 대책이 필요함. 
테스트는 TDD 방식으로 진행한다. 
- [참고1](https://joinc.co.kr/w/man/12/golang/TDD) TDD, 
- [참고2](https://blog.golang.org/race-detector) race (21-03-11)

## License
This project is [Apache License 2.0](LICENSE).