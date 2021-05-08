# WarungPintar2021

## Prequisites
- MySQL Database
- Postman
- kafkacat
```bash
apt-get install kafkacat
```

## Import Database
1. Migrate database from folder masterscv/database/migrations/schema.sql
2. Migrate database from folder inventoryscv/database/migrations/schema.sql
3. Migrate database from folder transactionscv/database/migrations/schema.sql

## Setup Environment for all service
1. Open file .env at folder mastersvc
```hashkell
DB_DRIVER=mysql
DB_USER=${dbuser}
DB_PASSWORD=${dbpassword}
DB_HOST=${dbhostname}
DB_PORT=3306
DB_DATABASE=wp_master
```
2. Change value for this parameter (for hostname,you can use your IP) :
   - **${dbuser}**  
   - **${dbpassword}**
   - **${dbhostname}**
3. Open file .env at folder transactionsvc
```hashkell
DB_DRIVER=mysql
DB_USER=${dbuser}
DB_PASSWORD=${dbpassword}
DB_HOST=${dbhostname}
DB_PORT=3306
DB_DATABASE=wp_transaction
KAFKA_HOST=${kafkahostname}
KAFKA_PORT=19092
KAFKA_NETWORK=tcp
KAFKA_TOPIC=inbound
KAFKA_PARTITION=0
```
4. Change value for this parameter (for hostname,you can use your IP) :
    - **${dbuser}**
    - **${dbpassword}**
    - **${dbhostname}**
    - **${kafkahostname}** 
5. Open file .env at folder inventorysvc
```hashkell
DB_DRIVER=mysql
DB_USER=${dbuser}
DB_PASSWORD=${dbpassword}
DB_HOST=${dbhostname}
DB_PORT=3306
DB_DATABASE=wp_inventory
KAFKA_HOST=${kafkahostname}
KAFKA_PORT=19092
KAFKA_NETWORK=tcp
KAFKA_TOPIC=inbound
KAFKA_PARTITION=0
```
6. Change value for this parameter (for hostname,you can use your IP) :
    - **${dbuser}**
    - **${dbpassword}**
    - **${dbhostname}**
    - **${kafkahostname}**      

## How to run services
1. ensure you have migrate all databases.
2. ensure you have change the environment file.
3. run command to build docker image
```bash
$ docker-compose build
```
4. ensure all images have been build (mastersvc, transactionsvc, inventorysvc, zookeeper, kafka)
5. run command to host all containers
```bash
$ docker-compose up -d
```
6. ensure all service has been running
```bash
$ docker ps
```
7. run this command for checking kafka is available
```bash
$ kafkacat -C -b ${kafkahostname}:19092 -t inbound -p 0
```

## Testing the service
1. open postman
2. import postman file `WarungPintar2021.postman_collection.json`
3. execute postman `transactionsvc - Add inbound`
4. if success you can check the message at **kafkacat** and check the database *wp_inventory* table *inventory* the new record should be there.