## Codegen
A lot of operations in the project require codegen, such as generating APIs from swagger documentation, generating ent schema, and generating dependency injection implementations. 

The following commands can be used to generate the required code.

### Generate the swagger documentation
```bash
bash bin/apigen.sh
```

### Generate the ent schema
```bash
bash bin/dbgen.sh
```

### Generate dependency injection
```bash
bash bin/gogen.sh
```

### Generate all
All the above codegen commands can be combined into one:
```bash
bash bin/allgen.sh
```