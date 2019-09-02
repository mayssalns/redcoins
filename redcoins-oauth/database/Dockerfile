# Create mysql server based on the official image from the dockerhub
FROM mysql:5.7

#ENV MYSQL_DATABASE oauth
#ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE oauth

COPY ./sql-scripts/ /docker-entrypoint-initdb.d/



