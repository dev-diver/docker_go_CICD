FROM python:3.8-slim

WORKDIR /app

RUN pip install flask flask-cors

COPY app.py .

CMD ["python", "app.py"]