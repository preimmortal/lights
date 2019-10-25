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
    BE:
        - Figure out how to organize my data models (i.e. DBDevice and Device should be the same model) 
            - Data input and output should rely solely on this data model
    FE:
        - Set up docker for build/deployment

    High Level:
        x CI/CD Pipeline

    Backlog:
        - API calls for other devices
