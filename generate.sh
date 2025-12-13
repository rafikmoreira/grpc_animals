#!/bin/bash

# Script para gerar código Go a partir dos arquivos proto
# Use: ./generate.sh
# Ou: make generate (que chama este script)

set -e

echo "Gerando código Go usando buf..."

# Tentar usar buf generate primeiro
if command -v buf &> /dev/null; then
    buf generate
    echo "Código gerado com sucesso usando buf!"
else
    echo "buf não encontrado, usando protoc diretamente..."
    
    # Encontrar o caminho dos imports do Google
    GOOGLE_PROTO_PATH=""
    if [ -d "$(go env GOPATH)/src/github.com/google/protobuf" ]; then
        GOOGLE_PROTO_PATH="$(go env GOPATH)/src"
    elif [ -d "$(go env GOMODCACHE)/google.golang.org/protobuf@latest/types" ]; then
        GOOGLE_PROTO_PATH="$(go env GOMODCACHE)"
    fi
    
    # Gerar código usando protoc (usando . para manter estrutura relativa)
    protoc \
      --go_out=. \
      --go_opt=paths=source_relative \
      --go-grpc_out=. \
      --go-grpc_opt=paths=source_relative \
      --proto_path=pkg/proto/v1 \
      ${GOOGLE_PROTO_PATH:+"--proto_path=$GOOGLE_PROTO_PATH"} \
      pkg/proto/v1/animals.proto
    
    echo "Código gerado com sucesso em pkg/proto/v1/"
fi

