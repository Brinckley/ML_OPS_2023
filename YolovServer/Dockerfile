FROM python:3.9

WORKDIR /code

COPY ./requirements.txt /code/requirements.txt
RUN pip install --no-cache-dir --upgrade -r /code/requirements.txt

EXPOSE 8089

COPY app /code/app
ENTRYPOINT ["uvicorn", "app.main:app", "--host", "0.0.0.0"]