FROM python:latest
WORKDIR /code
ENV FLASK_APP=app.py
ENV FLASK_RUN_HOST=0.0.0.0
EXPOSE 5000
COPY . .
CMD ["python", "app.py"]