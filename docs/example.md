# Example of repository file

```yml
---
session: <name>

before: <cmd>
stop: <cmd>

env:
    - VAR: env

windows:
    - name: <name>
      cwd: "./"
      panes:
        - type: horizontal
          cwd: ".."
          commands:
            - ls
            - zsh
        - type: horizontal
          cwd: ".."
          commands:
            - ls
            - zsh
    
            
```
