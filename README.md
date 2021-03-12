# NewMesos-go
## 개발환경 구축 
- Docker Compose 를 이용한 개발환경 구축 
```shell
$ make start // mesos master, agent, zookeeper, etc.. 
$ make stop  // 컨테이너 중지 및 삭제 
```

## Testing
```shell
$ make test
```

## 진행 상황
RESERVATION_REFINEMENT, CapabilityReservationRefinement=1 과 관련하여 [mesos 매뉴얼](http://mesos.apache.org/documentation/latest/reservation/) 을 참고해야함.
mesos scheduler 개발 시작. mesos executor 의 경우 컨테이너 환경에서는 개발이 어려움. 이 문제에 대해서는 추가적인 대책이 필요함. 
go mod init, tidy, vendor 사용. (21-03-11)

## License
This project is [Apache License 2.0](LICENSE).