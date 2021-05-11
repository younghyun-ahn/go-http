# go-http

http.Client 패키지는 http 패키지 와 마찬가지로 Get, Head, Post, PostForm 이 구현돼 있으며 동일하게 사용 할 수 있다.
http 패키지는 내부적으로 http.Client struct instance 의 각 메서드에 대한 alias로 되어 있다.(http.Client 를 호출한다)

* http.Request : 헤더가 필요하거나 GET/POST/HEAD 외 사용
* http.Client : 쿠키/프록시 사용
* http : GET/POST/HEAD

**유형**|**메서드**|**Go API**
:-----:|:-----:|:-----:
GET을 이용한 정보 획득|GET|http.Get
쿼리가 있는 정보 획득|GET|http.Get
HEAD를 이용한 헤더 획득|HEAD|http.Head
x-www-form-urlencoded로 폼 전송|POST|http.PostForm
POST로 파일 전송|POST|http.Post
multipart/form-data로 파일 전송|POST|http.PostForm
쿠기 송수신|GET/HEAD/POST|http.Client
프록시|GET/HEAD/POST|http.Client
파일 시스템 접근|GET/HEAD/POST|http.Client
자유로운 메서드 전송|ALL|http.Request/http.Client
헤더 전송|ALL|http.Request/http.Client
