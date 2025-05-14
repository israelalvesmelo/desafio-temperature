# Desafio FullCycle temperature

Este projeto é um desafio do curso FullCycle, onde o objetivo é criar uma API que retorna a temperatura de uma cidade a partir do CEP.

## Tecnologias

- [Go](https://go.dev/)
- [Chi](https://go-chi.io/#/)
- [Docker](https://www.docker.com/)
- [Google Cloud Run](https://cloud.google.com/run)


## APIs usadas

Esse projeto usa as seguintes APIs:
- [Weather API](https://www.weatherapi.com)
- [ViaCEP](https://viacep.com.br)

Sendo que a API do weather necessita de uma chave para funcionar. Essa chave pode ser obtida assim que você se [cadastra](https://www.weatherapi.com/signup.aspx) na API do weather.
Com a chave em mãos, você pode adicionar ela no arquivo `env.json` na raiz do projeto no campo `api_key`.
É possivel usar a chave `6ec8b1eb5d064e83a5804852251305` para testar o projeto, mas não garantimos que a chave esteja válida no momento que você executar o projeto.

## Como executar localmente

- Execute o comando `docker compose up -d` para subir o container docker com o projeto.
- Após subir o container do docker, você pode executar a API usando `curl` abaixo:
```c
    curl --location 'http://localhost:8080/temperature?cep=01310-100'
```
Resposta:
```json
{
    "temp_C": 16.1,
    "temp_F": 60.980000000000004,
    "temp_K": 289.1
}
```

## Como executar em produção (Google Cloud Run)

- Para executar o projeto diretamente do Google Cloud Run, você pode executar a API usando `curl` abaixo:
```c
curl --location 'https://desafio-temperature-dohpifgd7q-uc.a.run.app/temperature?cep=01310-100'
```
Resposta:
```json
{
    "temp_C": 16.1,
    "temp_F": 60.980000000000004,
    "temp_K": 289.1
}
```

## Como executar os testes de integração

- Para executar os testes de integração, você pode executar o comando abaixo no terminal:
    - Local:
      ```shell
        API_URL=http://localhost:8080 go test -v ./tests/integration
      ```
    - Produção:
      ```shell
        API_URL=https://desafio-temperature-dohpifgd7q-uc.a.run.app go test -v ./tests/integration
      ```



