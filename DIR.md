/example            # 프로젝트 루트 디렉토리
  ├── /cmd          # 애플리케이션의 엔트리 포인트
  │    └── /example
  │         └── main.go  # 'example' 애플리케이션의 시작 파일
  ├── /pkg          # 재사용 가능한 패키지
  │    ├── /api     # API 관련 유틸리티 패키지
  │    └── /db      # 데이터베이스 관련 코드
  ├── /internal     # 비공개 코드 (외부에서 import 불가)
  │    └── /config  # 구성 관리 패키지
  ├── go.mod        # 모듈 정보 파일
  ├── go.sum        # 의존성 체크섬 파일
