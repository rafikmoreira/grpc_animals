.PHONY: install-deps generate clean

# Instalar dependências necessárias
install-deps:
	@echo "Instalando protoc-gen-go e protoc-gen-go-grpc..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "Dependências instaladas!"

# Gerar código a partir dos arquivos proto
generate:
	@echo "Gerando código Go a partir dos arquivos proto..."
	@chmod +x generate.sh
	@./generate.sh

# Limpar arquivos gerados
clean:
	@echo "Limpando arquivos gerados..."
	@find pkg -name "*.pb.go" -delete
	@find pkg -name "*_grpc.pb.go" -delete
	@echo "Limpeza concluída!"

