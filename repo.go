package main

import ("fmt"

)


var currentId int

var inputs Inputs

func RepoFindInput(id int) Input {
    for _, t := range inputs {
        if t.Id == id {
            return t
        }
    }
    // return empty Todo if not found
    return Input{}
}

func RepoCreateInput(t Input) Input {
    currentId += 1
    t.Id = currentId
    inputs = append(inputs, t)
    return t
}

func RepoDestroyTodo(id int) error {
    for i, t := range inputs {
        if t.Id == id {
            inputs = append(inputs[:i], inputs[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}