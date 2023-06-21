# Go Application

 ## To run the golang application generated

  + Build the go application 
  ---
      cd go
      go mod tidy
  ---
  + First run the postgress,keycloak and jhipster registry files also start rabbitmq server.
  ---
      cd docker
      docker-compose -f  postgresql.yml up     
      docker-compose -f  keycloak.yml up     
      docker-compose -f  jhipster-registry.yml up   
      docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
  ---


  + Create a .env file in the go sub directory 
  ---
    ### .env file sample  
    - SERVICE_PORT= <port_number>
    - SERVICE_REGISTRY="NETFLIX_EUREKA"
    - SERVICE_REGISTRY_URL=http://admin:admin@localhost:8761/eureka
    - AUTHENTICATION_KEYCLOAK=true 
    - DB_HOST=localhost 
    - DB_USER=go
    - DB_NAME=postgres 
    - DB_PORT=5433
    - KEYCLOAK_URL=http://localhost:9080
    - CLIENT_ID=internal
    - CLIENT_SECRET=internal
    - REALM_NAME=jhipster
    + (DB_URL based on selection of url)
    - DB_URL=postgresql://go@localhost:5433/postgres
    - MESSAGE_BROKER=amqp://guest:guest@localhost:5672
---

  + Now get back to the root directory of go and start the golang service 

      go run .
  ---
  
## Build an image 
   ---
       docker build -t image_name .
   ---
### Run by passing values 
- docker run -it --network="host" -p <port_number>:<port_number> -e        db_url=postgresql://go@localhost:5433/postgres image_name 


### Run by passing 3 ENV values 

```
docker run -it --network="host" -p <port_number>:<port_number> \

-e DB_URL=postgresql://go@localhost:5433/postgres \

-e KEYCLOAK_URL=http://localhost:9080/realms/jhipster \

-e SERVICE_REGISTRY_URL=http://admin:admin@localhost:8761/eureka/apps/ \

be
```
    

