FROM python:alpine
# FROM python:slim-buster
ADD app /app
WORKDIR /app
RUN pip install -r requirements.txt
CMD ["python", "main.py"]
