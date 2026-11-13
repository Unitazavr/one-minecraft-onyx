package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.5" )

func W8TQS4ryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0wtp9z1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pcTfcTFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8f6RVCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMkjgcboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdwEHBnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LH4e7s22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTjIshOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qYWeyQrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rfTfmLhpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8EASW7xtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4YFt8MtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6l2LwJzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3bLYt2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYVSb42pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPFhWIVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hAvsfYGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8Eapf0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLz1H6haWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Abcm40GlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PbyIPVByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g16HZ1PmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1Bs91dvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oj2UWCjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yW8DBWXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZeLaPTygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IERms6cbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QF13ZDpOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOXOFbsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A27Xf9Q7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUiiTgUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtNHwxcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ShONXVxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHcEZX12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bRUq9Y3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PxPmq8L0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyCEvSBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cVfcmRu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ha1BVBKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ia8ZsNcMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yIizCLzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlIxiUVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVNYKE7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUUxf5TkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDe1sAKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzJfYHgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func idN2mgEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t1UWGHCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2psuqqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hmr2LmH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oj52wjv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4JYBLBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlSTOGnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2PLSo0JbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BswGHLG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X86IRCxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJHiOKTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txwLqfEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ys7ZAoUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1bVfucpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jf5dEO4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNYuu9SjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1h5RbhQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwPDnc1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wcogb6CbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1hf1KeBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vV6BjA6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9cr4AeRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxgnNn2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0J1ee0ylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IB2n54xjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bT1E1q9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Shbf2Aa8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVvmFeIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WOAmcS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVy5Ik8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RU5t0lAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zru6QClyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHzsB5aLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23BP8UPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fl95fnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TQgLUbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAfMYBxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c3Uhd53hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 75r1wM6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ch0WbcVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LaGXEKGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEwRkZxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjlE3jLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uvw6p6BCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22Q9WsAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3KzT6rfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hspu7OeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AONUGQXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLRdQZYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20g8PcblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gcNwvkX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MyGfNx8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0TCIIl5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYXbGslMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHTcRXCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9uM8U6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWxwjR16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cumTXsJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yk5xPdWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dW0PQuPJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tx6fSgHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exaxbRAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmFHOBHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c38XLbj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UN6RkQiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S51n0XkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6lksnlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGJC1y6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5zvZaMqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func saS2faoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7391KAt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IcFkYB2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pBd6qTS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xLRmaYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ecQG7M4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d91gHFLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OU51ekbBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRlPRIQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7AYESg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HryClNQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ijbwQGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WAnFDQCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmJJ6dnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XtJa2IFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2X0EkPgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A0BnUNxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fb0CA7dEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50pI0D05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIXBSTpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ApZFGznrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 396ltf9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzJwlTyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpIhDGaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FtxK6kUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWCL2bsDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rARTxQHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uiTbdkriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ctRpNELbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtehgbEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WU7xIhtjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OId8SInbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kihwISPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgIQVvTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNuyol27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dAeTordsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rAtXJCZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j6BXvQKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XHWJq60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewVq1ecOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRl6a34tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDKgwx14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XR5wQSqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DsjqFSrrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mf2kQfFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLt4QZXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3o7AROOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UTKy6kmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8nnEgRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func opPgjQsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hd03i9kPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRysqlnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZK2ECyZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f5KP2XQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suWo79DYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLdxLpgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHx0tbGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KC7x83ZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2F2pEREtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJb4BbgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8e3NhcWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ulkhkC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZoZyjWfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 499ItWrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqMKQOQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNaEcs80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o4tuIXMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXBzlhegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cYWOupAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pmj9WCdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LW8bQaOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8uFOrlLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vcZmFyPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vpJzOh9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rAMwPkDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gSOIsXZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k94iQjVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BOEM95cgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8idzKcRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5cJmA1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnIz2gDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sz59Y3fxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tC4dJevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1folOT09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKOZyaeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLoY3EggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2uvL69RKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaxviyXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QO02qDOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r88y61EeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gxoliLqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T69N9e1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2InJONMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05lsgqEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nKM3090Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LdmXKalWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TADmIZw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func akOl1jMdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GfGoM87ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NgroCstWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTjGMCRgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hT4IrcJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzLrum5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ge83dtitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ukc5k4fUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0d5aeIFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RyB3siz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZztrGepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCVhQYFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IWjDMJSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KmgxNv7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8rkLanXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IO2rSoviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jEZTNvtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHvOrSz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JuOlCkt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Gr4kIwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vnZBVxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKLySI5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 34VDEE03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uobCsBFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eGQDtrD6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7wmNJTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vmb2luntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8RXzsaRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rnjDZojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2aytGLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lHpXEEctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Fjl7UV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7IxwEobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPgcl78JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNYWUcH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZNy2yptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICdCVgHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfnzkpvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38EHhK49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R4QOUGuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZi8jsBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SaKzUOUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GpQ2IqrjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZCBkF0cnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t0oQl6ueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CcojQFfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZnj7ft7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oTTgZVECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITIPW1FCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8V2fg9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Xs1u4BVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJ41vf7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgFCw4c7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCuNprbYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXxQCut4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlpP9RKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTckkPEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXKfEv8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttr7XHYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wg6wRMwKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izgZM0BDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDns8RlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCCursT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fsq5b0FUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFDmg29VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cES1ewtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOx1ISukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YgaVk52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wuD6qFiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7M1KD4vvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AlCon9DJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JAPdCWVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQmdIFbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwkMlj5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7B3KgjwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 020RBBfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CywsNsOGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkqEXWkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WK3RegzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAqSk99SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGVaSKs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaSbYYQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2nzrKflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNh1NJgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McKTNHKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxDB8YD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxmIF4CKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PS1mRzGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9SRquHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Sf0A83yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sef1bTlOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPaeft2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKWVoGObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dFkTBEflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMk7MP0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoTgDImtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func znLBoiZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldQbrT68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8DFu5bwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKg3Cm3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78qrvc1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBGTvzF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YofOteZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GpG2IKGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K4nsYvKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ChAOeCR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwoFWYfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4YNpmxMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPggkr54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQWoWXiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGR61FtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c0NnpqSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eGzsQ91bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLQXO3mYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfnHkfRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hy1MuVtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5MqHHHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mGc3lVrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1vLCpQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c6O7QEwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ArifiBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cN0vRuKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ygzZQG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1u040ZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUYsMqSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tppL9ScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXViutWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ieoSCDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvBrI3DrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHCzeedsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0tNa7neWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6FgCQuByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPDxLeinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HFvfItSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VCjl0KhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zlfa5Mj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fq7rjptkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JrfQSzOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDA0m6wHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XM7dl2QzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRILJ6MpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qsyswwOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sz1Gih5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7iaCJxLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5zKfXLkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PuQPyKhpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KT5arfr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0ruVPMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55AobPAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3sTXst7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTncvn8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 630IyfWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func plXp2L3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzYmvsMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XWrCWAkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gSuYUmJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qpxo1lD8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Izh4Q8OMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8K4XhIYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WwHin9umWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yIVMLGzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71n9Tr70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKxFA3KTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6BNEbWRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMUldoghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rpT3hN14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PG0E6StKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2GNvvFQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KjNycA6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adtUMqKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jde4xjOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fW0wM2bMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RsLyrHlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FRDQOndNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNM08n71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9jiG9sfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XC7OhZLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ufsg63hQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJSq4VORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2nSB8paWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFR5A0IZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCWm1AHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LwHHdUzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VNhOyZrOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K4XUf5zyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eA0W2B88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHEr8vKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g59HRhO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqAoP9gmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VK7TDMKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eymqE4EiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VbVw6Z0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2d2vOHUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgQRHs3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3HBybhdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGPsqbiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QDHYX6K8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yoefo92EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5vuLpOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I4NYTJoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkDXWcCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtNpe8TPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJwuxtNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GyGJzjGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3sWvtJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TpsZRvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3h8MSz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGjOhuOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2fND4xTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OsCB37SuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LoE83iNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsj5n39eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZuGp6sNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ufe5cPYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sbb90DIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPQuIGDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zl4pRpTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jcKHs66jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Yb8yexyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygj8AsijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BEn2uSzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5NWwY21fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IkqRdGHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V53m69TSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mtHOCuEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VX6NNVa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3SvLzQjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AF1COVg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vqtrDkX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rT3mJZaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aoFYPpwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEDP7uTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQiZ9NzpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CUtp3dnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2m3eOxRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1dUTdfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmDiX3E5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QLEPAPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvtML0U4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9B3F0ybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZddKCzCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfn5yorhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ze3XQXsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAEiBtHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmqrWMbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D8wB0Z2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iiGCKWZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rUkToB6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQPMCyVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9ql3XIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJ9g6hd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcnby0GeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZWVbdDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTkQDZKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PiT5e3y2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVTfckCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sl760kVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gb0HOT8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uHlXj7KPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZ3Vqvi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IwSMG6OCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvlFaZuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5KhDOfzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9F38zcfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4waYv21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SbUxxObGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OcRKNoe6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0nxuvJlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uT0RdaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6466j3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vTEL7oLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3ZNJ6WXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95OLMc6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ci2ypQU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vAvLWuKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwM07Z0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N43C5BKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L74HRFqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kh5bMHLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gPSrIp7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BoUCGmVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCBtO5b2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rouz8d2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6kAovHQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B87RvdDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TAkbnuL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boGB4YdWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYrYf9neWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jlA0Er2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZKFeq16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGQEbouNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WAtcI6U5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func md73eR7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AK0syx7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Re27xgqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OBSu5Gc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFRxZCKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HxTWI4foWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvRKpC3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dz4nlqsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfjExLsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZBDWOU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YSsKOOLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0BMxBDEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U4J9B3OtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YSFBXGanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWFcfULMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFWRXanqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pBntI5Q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALP9YO10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwsOjMFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njQS8AD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UyCpKTHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MOycLvH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIdbx7EeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V072Klp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbBikgaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Oc6kaLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78e6KgEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jbzXvc0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThFgrNElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9a2HrYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YWrxqF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2Aecw6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S36dQkcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pdgwoyaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39TBGp9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NPnlz3nDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2PLpLXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Le2hiQiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADBbtO8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AY2zXf2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pUFW8fyMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jy3yVwatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n9JaFa7NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLFDr8lKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmT6nrI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lG7yDcKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lE3gRMskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ux2xzgYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vmxq2gIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMxs4ybaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7ZvEyXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tTE7VpHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwRQJTcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func awgXeJWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XnObQ84BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 01Nd2cpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4JJ8G0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJoB1AmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLXVxBv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ovq68KmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIpWnWs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGojBZ7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDkXfmZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KekNT2lFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func penwJTAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HUn6G4inWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c1Fli8BxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QcogkwSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PbZ1mTTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 565vvnWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nDFTWozaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NLmV8eahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3SYQ9ge0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rk8Hqz7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tv35ErUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Y4gTCxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MtLkkmzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCoxdlubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pevGv2mrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0pJMzdNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqCQFdDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQQCo7iXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSBA47mrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFddh0mfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A1CA9dXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7D6GEEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bonHh7hbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fb3JWAkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u1tJbGf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nAjSKTinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y58VYL2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmGlHpDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func skAwdSUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zSvAQLKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5CTIPbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LOTfLszBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFsZYPLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajlLxzDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFZNRamEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1S7T9u61Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISIsDW8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vAe4MGMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oShyoyrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BnsN7hukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3PohuI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o64w3FOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRXReeHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6RbzrIqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qCKEcbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5pW9Df0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ec95qyPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RLvGeE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OctW87UqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVWXJdI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func piy5jYkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sc53KXe3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5eglUfllWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Tr9aLh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dW2vQI2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDAdzm0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 91T4dn34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8Zi9qivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XnF8A23wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHwGo1rwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hBAig2r2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZUxnhIO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PvJFpLvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNxhWH0wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eEqdU2YwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6Qf6TQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QaVqUJ6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3t0dcQdeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oSxNumIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ha9PD7pWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BpeHcbvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e70aUq2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8EPrjrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWXrVfbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04ov7oAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJT1grG3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dijtuIHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xOlqDIT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1E36s4f6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNGoBxKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGO6n2euWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yd7vK4ePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wdi0Er07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8fqRblYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdnCGeSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6jsSXQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4KHm8yKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eNvq0KkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y86ji6pxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gEcRfHKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDUOisYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TqFG5EXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ClboKc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSYnNc2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCEH97pfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgHn7UtJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNfVYHpSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJHj9pr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PsImEQqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VV4Z9cP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8BbIZ1lrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func smASpxBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zc7GkfvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8luLTJb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dT38ieKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wd8T9cLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IthT0VDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9oEoEdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04Qt7rybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JjLXv0g0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEXtwtpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6BxXIazaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1a3DOGsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdkXRhq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7fRuQ9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wM70fsdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUBklvvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ro961JwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O38VF3WmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6abnmfgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9tzXYpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXCDTM3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GkQD0LU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cG6dSevRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wi7Pb3MQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrMEeB0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cAZjCDG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9UwqBv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SlRCc19aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfUB6peqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKmT3fqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozhEbknZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vV8F489UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MV1C5QhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ezu73rbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2RHuJn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJURMBeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YN3ulIDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lF2roCcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mbAMigD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOqyQhedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9R8MH2XwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLA3vtUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YyE9DGh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iv4qBaxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsdue9mMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LIZo0HtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfTofPMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func paWlWr3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gnFyQVj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TeM0PfoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BOC4u2Y1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WWLtTDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dgY7hlTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FpDDjrjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gweBrXjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjovrcQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EBcOSmp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxD0tGlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lguCjsUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func epiuRgTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SCc7fMZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qltJBvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ee9XpV7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CVRd7LaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QS296gDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLCkbNARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eO9Sere3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yipFydMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iiKNqGYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMEFj9gwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fcwzc5mZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGwMBWgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7oo2olVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLtuzYtgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEtFsxH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tu74jMJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ai2rxI6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YVlpoFuAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dXdFxYgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvDPB7DwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2Tlzi3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ecX6YRgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZUdfjTcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJEhfwlrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vOJKRqu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EC8jUT51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5rGV7EqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zrYXtOvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2LzgJsVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KJ5FgKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D9DAdz97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cP9ryQLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxboh2maWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4iBL2r3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwOzUUTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M04JlU0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzCq4pevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n84aigmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tkCHLi5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KtrDpBB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LB4djkEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1w9Zx2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YSUX4A1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrDWraBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qe4pPoeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltBtHWCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYwhCOd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkicQ8rUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPd7dL8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x035iJqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jX7mgQjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1mCW2XdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rU4HcuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7QfjBPX2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 899e2OtcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBPzdryTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QOczwnW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4stRcUhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wu1CFvV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yxjb2tpRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rpW2eoOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dbq7LtSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uNXJWGolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUYX504TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bz652pHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ssS7YgUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vb8OLEI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJNc0ajnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GeLvRq5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvOweCTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jaAbEvxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1L81Sqn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mfMrf6P3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8yFQ5ovHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HSXmI1MlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v34SpJEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func addG6WCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6arF793jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func he6jOizUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aJCPgX0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybdIQDiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2X2PKFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7VIgoQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uc9oCGXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToJ5YMUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dspzmpwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6rts0p9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i980GX8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23sjelFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lko3NU5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3xA9rCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtjuORUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z9SrZwNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HWovD8GAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xr6I2be7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Punzl5TgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tTDWVLMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OaWB3M7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQP3ZERpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrjXTLb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWL1uZ0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5aI6HhnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9tRvRRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0bVbAHI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hh00n0JSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsrv0ZumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7suOvjT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKarqQNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUnhAGPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vG7GT20OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hxkCaJ0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTZsPyjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MnzJ5EARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cz5ApKwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BGSoBfooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUB6TJvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRJmcsw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LolnmYH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cek9xnKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8pnkmIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2YuM0xc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJa8sBr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pg51quKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4NQADHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sv7sLGFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gUk4aHeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFLNTECWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GhBTLioHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGPbdEdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1va1rqDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKUcGz1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enK7z7UXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6NCtwv0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3vyoClGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVgdMhyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLvozOOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1Ev9VFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFM2IaIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCGwbWZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vzv8D6zZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wbmYF0ziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FSyJgm71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUgzsJLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fa9vhrNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dm6NW8pAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lyVilJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJoVPt2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqTWoM2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CtWGAdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xm7zpFZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjQ1qrVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b6q6bh1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06Dt7d2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7ZG4UdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXqghqr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNDLRxXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XyZGkZ6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJuZaIdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ahV2HLWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Lm1iHSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGFXjoTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aI7xiyYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTuXoYL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dqQBpCZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLvu1quRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Td8CDCA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nypHCoOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTMHJVyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZO0w31xMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tAbbAnurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ooR9461yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRsRfTkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func USOHsMN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkM5UrwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9L9eEgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7tlUZttaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cI1hJgRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func so6FShl3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func as6ebBujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4NYBJGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Qs9iQ87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oD6UJpf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4YjjKaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFFUytNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsfDSwtzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVdnLTn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SLIoNpYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STL2fEN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2X9f5iSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qy8uh96kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DyI7iOSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IH1FFptDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDR8sWCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

