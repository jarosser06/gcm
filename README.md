Go Cloud Monitoring
===================
Not a fan of pyrax the CLI didn't feel very intiutave plus I prefer
compiled binaries for shipping software.  This CLI tool has single
commands and an interactive shell.

###External Dependencies
This tool uses gpm for dependency management so you need to
ensure that it is installed in order to build from source.

https://github.com/pote/gpm

###Commands
object verb optional attribute

List Monitoring Account Information

```shell
gcm -t <token> -limits
```

List All entities

```shell
gcm -t <token> -entities
```

List Entity Information

```shell
gcm -t <token> -entity <entity_id>
gcm -t <token> -entity <entity_id> -a
```

List Checks

```shell
gcm -t <token> -entity <entity_id>
```

Interactive Shell
```shell
gcm>auth set endpoint <endpoint>
gcm>auth set user <user>
gcm>auth set type <rsa/password>

-- Example with RSA --
gcm>auth set user jim
gcm>auth set type rsa
gcm>auth
Enter RSA Pin+Token: mypin234647824
Auth Successful or Auth Failed
jim@gcm>

-- Example with Existing Token --
gcm>auth set user jim
gcm>auth token <token_id>
Auth Successful or Auth Failed
jim@gcm>list entities
jim@gcm>entity set id <entity_id>
jim@gcm#entity_id>show
jim@gcm#entity_id>list checks

-- Test Check --
jim@gcm#entity_id>check test <check_id>
jim@gcm#entity_id#check_id>test

jim@gcm#entity_id>exit
jim@gcm>
```
