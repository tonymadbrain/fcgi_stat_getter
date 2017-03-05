# FCGI stat getter for Zabbix monitoring

# ENG

Script for getting statistics from PHP-FPM server for Zabbix external check. Implementation of this script - http://doam.ru/fcgi_monitoring_with_zabbix/ on Golang.

## Setting up
### PHP-FPM

In PHP-FPM Pool config enable `status` and `ping`:

```Bash
pm.status_path = /status
ping.path = /ping
```

And restart server.

### Zabbix

We need to compile binary for that platform where zabbix is running. Use this command:

````Bash
env GOOS={OS} GOARCH={ARCH} go build -v github.com/tonymadbrain/fcgi_stat_getter
````

Where:

{OS} - OS type:

* Mac os - darwin
* Windows - windows
* Linux - linux
* FreeBSD - freebsd

{ARCH} - arhitecture:

* x86_64 - amd64
* x86 - 386
* ARM - arm  (linux only)

Example for Linux x86_64:
```Bash
$ env GOOS=linux GOARCH=amd64 go build -v github.com/tonymadbrain/fcgi_stat_getter
```

Then, copy the binary to Zabbix server into `/usr/lib/zabbix/externalscripts` folder, make him executable with `chmod +x fcgi_stat_getter`, set zabbix owner with  `chown zabbix:zabbix fcgi_stat_getter`.
Next, import template `zbx_fcgi_template.xml` in Zabbix frontend and attach him to server(s).

Done! Wait for data.

# RUS

Скрипт для получения статистики из PHP-FPM, который можно использовать в Zabbix. Реализация на Go вот этого скрипта - http://doam.ru/fcgi_monitoring_with_zabbix/.

## Настройка
### PHP-FPM

В конфиге PHP-FPM пула нужно включить статус и пинг:

```Bash
pm.status_path = /status
ping.path = /ping
```

И сделать restart сервера.

### Zabbix

Нужно скомпилировать бинарник под ту платформу, на которой запущен Zabbix сервер, для этого нужно использовать команду:

````Bash
env GOOS={OS} GOARCH={ARCH} go build -v github.com/username/fcgi_stat_getter
````

{OS} - тип операционной системы, может быть:

* Mac os - darwin
* Windows - windows
* Linux - linux
* FreeBSD - freebsd

{ARCH} - архитектура, может быть:

* x86_64 - amd64
* x86 - 386
* ARM - arm  (linux only)

Example:
```Bash
$ env GOOS=linux GOARCH=amd64 go build -v github.com/tonymadbrain/fcgi_stat_getter
```

Закинуть бинарник на сервер Zabbix в каталог `/usr/lib/zabbix/externalscripts`, сделать его исполняемым - `chmod +x fcgi_stat_getter`, сделать владельцем файла Zabbix - `chown zabbix:zabbix fcgi_stat_getter`. Затем нужно импортировать шаблон `zbx_fcgi_template.xml` в Zabbix фронтенде и прикрепить его к нужному серверу.

Ждать данных.
