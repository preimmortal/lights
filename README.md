# smarthome
Smarthome API

# tplink testing
go test -v --cover github.com/preimmortal/smarthome
go test -v --cover github.com/preimmortal/smarthome -run Encrypt

# docker build
docker build -t github.com/preimmortal/smarthome:latest .

# docker run
docker run --network host github.com/preimmortal/smarthome:latest

TODO:
    FE:
        - Disable CORS to allow client to make CORS request to backend

    High Level:
        x CI/CD Pipeline

    Backlog:
        - API calls for other devices
