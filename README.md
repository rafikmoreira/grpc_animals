# gRPC Animals

Projeto gRPC para gerenciamento de animais usando Protocol Buffers.

## Gerando Código a partir dos Arquivos Proto

Este projeto usa Protocol Buffers para definir os serviços e mensagens gRPC.
Para gerar o código Go a partir dos arquivos `.proto`, siga os passos abaixo.

### Pré-requisitos

- Go 1.25.5 ou superior
- `buf` instalado (opcional, mas recomendado)
- `protoc` instalado (alternativa ao buf)

### Instalação das Dependências

Primeiro, instale os plugins necessários para gerar código Go:

`make install-deps`

Gerando o Código

Após instalar as dependências, gere o código usando:

`make generate` ou execute o script diretamente:

./generate.sh

### Arquivos Gerados

Os arquivos gerados serão criados em `pkg/proto/v1/`:

- `animals.pb.go` - Contém as mensagens e tipos (Animal, GetAnimalRequest, etc.)
- `animals_grpc.pb.go` - Contém o cliente e servidor gRPC (AnimalsServiceClient, AnimalsServiceServer)

### Limpando Arquivos Gerados

Para remover os arquivos gerados:

`make clean`
