# Core

## Componentes

### Domain

Todos os modelos de domínio estão nesse diretório (./internal/core/domain). Contém uma "go struct" pra cada entidade que é parte do domínio do problema e que pode ser utilizado por toda aplicação.

**Observação** Nem toda "go struct" é um modelo do domínio. Somente as structs que estão envolvidas na lógica de negócio.

### Ports

As "ports" são armazenadas no diretório `./internal/core/ports`. Ela contem as definições de interfaces utilizadas para se comunicar com os atores.

### Services

Services são as portas de entrada para o `core` e cada uma delas implementa a sua `port` correspondente. Eles estão armazenados no diretório `./internal/core/services`