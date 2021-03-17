# api.inssa.club_clubhouse

api.inssa.club의 clubhouse 관련 서비스입니다.  
serverless framework를 기반으로하여 serverless.yml 이 작성되어 있습니다.  
로컬에서 debug 모드 및 release 모드 serverless 환경에서 apex gateway를 통한 release 모드를 지원합니다.  

## 프로젝트 구조
 - cmd/server
    - 메인 프로젝트 디렉토리
    - 그 안의 구조는 생략
 - configs
    - db 연결, 환경변수와 같은 설정 파일이 담겨있는 곳

## 환경 변수
실행에 필요한 환경 변수는 [configs/envs.go](./configs/envs.go) 에서 확인 할 수 있습니다.  

## <a name="execution"></a>실행
프로젝트의 root에서, 필요한 패키지들을 다운 받기 위해 다음의 명령을 실행해주세요.
```sh
go mod tidy
```
다음의 명령어로 프로젝트를 실행 할 수 있습니다.
```sh
go run cmd/server/main.go
```
## 바이너리 빌드
### 리눅스
```sh
make build
```
### 그 외
```sh
go build -ldflags="-s -w" -o bin/main cmd/server/main.go
```
## API 문서
### 실행하여 문서를 얻기
코드를 [실행](#execution) 하여 /swagger/index.html 로 이동하면, 그 곳에서 문서를 얻고 api들을 테스트 할 수 있습니다.  
기본 적으로 8080 포트에서 열리기 때문에, <http://localhost:8080/swagger/index.html>에서 접근 할 수 있습니다.

### 직접 문서 파일 얻기
코드를 실행하지 않고 직접 json이나 yaml로 된 문서 파일을 얻고 싶다면, [cmd/server/docs](./cmd/server/docs) 에서 docs.json과 docs.yaml을 얻을 수 있습니다.