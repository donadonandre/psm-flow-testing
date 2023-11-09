FROM openjdk:17-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

EXPOSE 8090

CMD ["java", "-jar", "app.jar"]