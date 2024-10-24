package main

import (
    "fmt"
    "gerenciador-tarefas/tasks"
    "os"
    "strconv"
)

func printUsage() {
    fmt.Println("Uso:")
    fmt.Println("  add [descrição]        - Adicionar uma nova tarefa")
    fmt.Println("  list                   - Listar todas as tarefas")
    fmt.Println("  complete [id]          - Marcar uma tarefa como concluída")
    fmt.Println("  remove [id]            - Remover uma tarefa")
}

func main() {
    if len(os.Args) < 2 {
        printUsage()
        return
    }

    command := os.Args[1]

    switch command {
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Por favor, forneça a descrição da tarefa.")
            return
        }
        description := os.Args[2]
        err := tasks.AddTask(description)
        if err != nil {
            fmt.Println("Erro ao adicionar tarefa:", err)
        } else {
            fmt.Println("Tarefa adicionada com sucesso!")
        }
    case "list":
        taskList, err := tasks.ListTasks()
        if err != nil {
            fmt.Println("Erro ao listar tarefas:", err)
            return
        }
        for _, task := range taskList {
            status := "Pendente"
            if task.Completed {
                status = "Concluída"
            }
            fmt.Printf("%d - %s [%s]\n", task.ID, task.Description, status)
        }
    case "complete":
        if len(os.Args) < 3 {
            fmt.Println("Por favor, forneça o ID da tarefa.")
            return
        }
        id, _ := strconv.Atoi(os.Args[2])
        err := tasks.CompleteTask(id)
        if err != nil {
            fmt.Println("Erro ao completar a tarefa:", err)
        } else {
            fmt.Println("Tarefa concluída com sucesso!")
        }
    case "remove":
        if len(os.Args) < 3 {
            fmt.Println("Por favor, forneça o ID da tarefa.")
            return
        }
        id, _ := strconv.Atoi(os.Args[2])
        err := tasks.RemoveTask(id)
        if err != nil {
            fmt.Println("Erro ao remover a tarefa:", err)
        } else {
            fmt.Println("Tarefa removida com sucesso!")
        }
    default:
        printUsage()
    }
}
