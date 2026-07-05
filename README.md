## Fluxo de uma Requisição

```mermaid
graph TD
    A["🌐 POST /api/transactions"] --> B["📥 Handler<br/>Valida input"]
    B --> C["🗄️  Repository<br/>Executa query"]
    C --> D["🔍 PostgreSQL<br/>Query"]
    D --> E["📤 Repository<br/>Retorna struct"]
    E --> F["✅ Handler<br/>Retorna JSON"]
    
    style A fill:#e1f5ff
    style F fill:#c8e6c9
```

## Entidades e relacionamentos (Tabelas)

<img width="769" height="621" alt="image" src="https://github.com/user-attachments/assets/2a68a3ae-c7ff-4ed5-8413-11ae66ecb15e" />

## 🧪 Casos de Uso Pra Estudar
Criar Conta:
Validar CPF (regra de negócio)
Inserir no banco
Retornar ID

Transferência (o mais complexo):
Validar saldo suficiente
Bloquear race conditions (isolation level)
Atualizar ambas contas em transação
Registrar na tabela de transações
Aqui eu estudo: ACID, locks, deadlocks

Listar Transações com Filtros:
Por data, status, tipo
Paginação
Ordenação
Estuda: query optimization, indexes
