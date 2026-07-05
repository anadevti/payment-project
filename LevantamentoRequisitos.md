## Criacao de conta (user), o que deve ter ao criar uma conta?
- Nome
- email
- telefone
- endereco
- data de nascimento
- CPF

## O que uma conta tem?
- id da conta
- user id da conta ──→ chave estrangeira
- numero da conta
- amount
- tipo da transacao (pix, ted)

🎯 Ordem de Implementação
Fase 1:

- Desenhar schema SQL completo
- Criar Models em Go (User, Account, Transaction structs)
- Escrever migrations

Fase 2:
- Criar Repository (queries CRUD)
- Criar Handlers (endpoints)
- Testar