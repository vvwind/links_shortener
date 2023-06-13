# links_shortener
A simple grpc-service which can help shortening links

#Запуск DB => sudo docker-compose up --build

#Запуск сервера => go run server/server.go

Флаг -m задает вараинт использования памяти:
mem - внутренняя память
db - база данных


#Запуск клиента => go run client/client.go

После флага -sh указывается ссылка для сокращения.
После флага -rev указывается ссылка для обратного преобразования.

Допускается испльзование только одного флага на один запрос.




