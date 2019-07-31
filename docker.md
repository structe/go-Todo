#
docker pull mysql

/docker/data/mysql

docker run -d --name mysql1 -v /docker/data/mysql:/var/lib/MySQL -e MYSQL_ROOT_PASSWORD=open -p 3500:3306 mysql

docker ps

docker exec -t -i mysql1 /bin/bash

mysql -u root -p

create database meeting;

sudo passwd root

use mysql;

update user set host = '%' where user = 'root';

flush privileges;

