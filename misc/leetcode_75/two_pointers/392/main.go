package main

func dequeue(queue []rune) []rune {
    if len(queue) > 1{
        return queue[1:len(queue)]
    }
    return nil
}

func isSubsequence(s string, t string) bool {
    if s == "" {
        return true
    }
    queue := []rune(s)
    for _, r := range t {
        if r == queue[0] {
            queue = dequeue(queue)
        }
        if queue == nil {
            return true
        }
    }
    return false
}