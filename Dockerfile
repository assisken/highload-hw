FROM tiangolo/uwsgi-nginx-flask:python3.8

WORKDIR /app

COPY ./Pipfile Pipfile.lock ./
RUN pip install pipenv \
 && pipenv install

COPY ./app ./
EXPOSE 8000

CMD ["/usr/local/go/bin/go", "./main.go"]
