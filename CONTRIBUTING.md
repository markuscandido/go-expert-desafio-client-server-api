# Guia de Contribuição

Obrigado por considerar contribuir para o projeto de Cotação de Moedas! Este guia irá ajudá-lo a configurar o ambiente de desenvolvimento e entender como contribuir de forma eficaz.

## Começando

### Pré-requisitos

- Go 1.24.2 ou superior
- Git
- SQLite3 (para testes locais)
- Uma ferramenta para testar APIs (como curl, Postman ou Insomnia)

### Configuração do Ambiente

1. Faça um fork do repositório
2. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/desafio-client-server-api.git
   cd desafio-client-server-api
   ```
3. Instale as dependências:
   ```bash
   go mod download
   ```

## Fluxo de Trabalho

1. Crie uma branch para sua feature/correção:
   ```bash
   git checkout -b feature/nome-da-feature
   # ou
   git checkout -b fix/corrige-problema
   ```

2. Faça suas alterações seguindo as convenções de código

3. Execute os testes:
   ```bash
   # No diretório raiz
   go test ./...
   ```

4. Certifique-se de que o código está formatado corretamente:
   ```bash
   go fmt ./...
   ```

5. Verifique por problemas comuns:
   ```bash
   go vet ./...
   golangci-lint run
   ```

6. Faça commit das suas alterações com mensagens claras e descritivas

7. Envie as alterações para o seu fork:
   ```bash
   git push origin sua-branch
   ```

8. Abra um Pull Request (PR) para a branch `main` do repositório original

## Convenções de Código

### Estilo

- Siga o [Effective Go](https://golang.org/doc/effective_go.html)
- Use `go fmt` para formatação
- Nomes em inglês
- Comentários em inglês

### Estrutura do Código

- Funções pequenas e com responsabilidade única
- Tratamento adequado de erros
- Uso de contextos para cancelamento
- Logs claros e informativos

### Testes

- Crie testes para novas funcionalidades
- Mantenha a cobertura de testes alta
- Use tabelas para testes com múltiplos casos
- Testes devem ser independentes e repetíveis

## Processo de Revisão

1. Crie um PR com uma descrição clara das alterações
2. Adicione labels apropriados
3. Atribua revisores se necessário
4. Responda aos comentários e faça as alterações solicitadas
5. Após aprovação, um mantenedor fará o merge

## Reportando Problemas

Ao relatar um problema, inclua:

1. Descrição clara e concisa do problema
2. Passos para reproduzir
3. Comportamento esperado vs. real
4. Capturas de tela se aplicável
5. Versão do Go e do sistema operacional

## Diretrizes de Segurança

- Nunca exponha chaves de API ou credenciais
- Valide todas as entradas
- Use contextos para timeouts
- Mantenha as dependências atualizadas

## Estilo de Commit

Use o seguinte formato para mensagens de commit:

```
tipo(escopo): descrição curta

Descrição mais detalhada se necessário

Fixes #issue-number
```

Tipos de commit:
- feat: Nova funcionalidade
- fix: Correção de bug
- docs: Alterações na documentação
- style: Formatação, ponto e vírgula, etc. (sem mudança de código)
- refactor: Refatoração de código
- test: Adicionando testes
- chore: Atualização de tarefas, configurações, etc.

## Dúvidas?

Se tiver dúvidas, abra uma issue ou entre em contato com os mantenedores.

## Agradecimentos

Obrigado por dedicar seu tempo para contribuir com este projeto! Sua ajuda é muito valiosa.
