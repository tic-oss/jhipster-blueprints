# jhipster-blueprints

The JHipster team maintains several official blueprints. These blueprints have two main goals:

   > Enhance JHipster with new features using different languages and/or support

   > Demonstrate how the main generator behavior can be modified to fit anyone’s needs

For more Infromation

refer: https://www.jhipster.tech/modules/official-blueprints/

---
## How to generate jhispter blueprint?

1) Create a folder with the following prefix "generator-jhipster-" followed by the your generator name 

> mkdir generator-jhipster-go

2) Build your generator, using the jhipster blueprint generator

> cd generator-jhipster-go

> jhipster generate-blueprint

3) Follow the instructions and fill the prompts as per your need.

![Alt text](https://i.imgur.com/yV91YGd.png)

Let's say we want to override server sub-generator, select the server and press Enter
![Alt text](https://i.imgur.com/D5zNonh.png)

![Alt text](https://i.imgur.com/I4UnclC.png)

![Alt text](https://i.imgur.com/upBLebw.png)

![Alt text](https://i.imgur.com/oraYxUt.png)

---
## How to point your blueprint to local jhipster?

Make sure you have updated the generator-jhispter source directory
to the locally installed generator-jhispter in package*.json under dependencies section of the blueprint.

> "generator-jhipster": "7.10.0",


FOR EXAMPLE :-


> "generator-jhipster": "file:/home/rakesh/jhipster/generator-jhipster",

---

## How to link the blueprint?

Execute the following command:-

> npm install && npm link generator-jhipster-go 

## How to run the blueprint?
Run the following command:-

> jhipster --blueprints generator-jhipster-go
 
 ---
## How to use blueprint with jdl?

```
application {
    config {
        baseName be,
        applicationType microservice,
        packageName com.cmi.tic,
        authenticationType oauth2,
        databaseType sql,
        prodDatabaseType postgresql,
        devDatabaseType postgresql,
        serviceDiscoveryType eureka,
        serverPort 9001,
        blueprints [go]
    } 
}
```

Create a file  with the extension of the jdl, for example test.jdl 
and add the above piece of code into the jdl file.

Run the following command:-
> jhipster jdl ./test.jdl --skip-install --skip-git --skip-jhipster-dependencies