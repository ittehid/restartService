# Управление службой Windows (Остановка и Запуск)

## Описание

Данный проект представляет собой утилиту на языке Go, предназначенную для автоматической остановки и запуска Windows службы.  
Рассчитан на использование вместе с планировщиком задач Windows, для автоматизации процесса управления службой.  

Программа выполняет две основные операции:  
1. Остановка службы с использованием команды taskkill.  
2. Перезапуск (запуск) службы с использованием команды net start.  

Программа не предусматривает какую-либо проверку на наличие службы, ее статуса и вообще какое-либо логирование ошибок. Подразумевается, что служба есть и вы в этом уверены.

Убедитесь, что у вас есть права администратора для управления службами.

## Как изменить имя службы?

По умолчанию в программе задано имя службы GATE_LNK.  
Чтобы изменить службу, отредактируйте значение переменной serviceName:

    var serviceName = "Имя_вашей_службы"

Сохраните изменения и пересоберите программу.