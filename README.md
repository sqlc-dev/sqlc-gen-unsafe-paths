## Usage

```yaml
version: '2'
sql:
- schema: schema.sql
  queries: query.sql
  engine: postgresql
  codegen:
  - out: gen
    plugin: test
plugins:
- name: test
  env:
  - PATH
  wasm:
    url: https://github.com/sqlc-dev/sqlc-gen-test/releases/download/v0.1.0/sqlc-gen-test.wasm
    sha256: 138220eae508d4b65a5a8cea555edd155eb2290daf576b7a8b96949acfeb3790
```
