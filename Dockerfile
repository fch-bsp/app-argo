
# 01 Gerando uma imagem no Golang  as build criando mais de uma imagem 
FROM golang:1.17 as build
#Criando uma IMG Golang 
WORKDIR /app
COPY . .
# Copia tudo que tem na app para container 
RUN CGO_ENABLED=0 go build -o server main.go 
# copilando -> aplicação = main.go gerando um arquivo server que é nosso executável 


# 02 Cpiando o Arquivo que foi gerado dessa IMG para essa IMG 02 alpine
FROM alpine:3.12
#Criando uma IMG Alpine
WORKDIR /app
COPY --from=build /app/server .
# copiando tudo que está no "SERVER" para dentro do container
CMD ["./server"]
# Quando o container subir executa CMD o arquivo executável 

