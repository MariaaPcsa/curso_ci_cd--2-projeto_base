FROM ubuntu:latest

EXPOSE 8000

WORKDIR /app

ENV HOST=localhost
ENV PORT=5432
ENV USER=root
ENV PASSWORD=root
ENV DBNAME=root


COPY ./main main

CMD [ "./main" ]