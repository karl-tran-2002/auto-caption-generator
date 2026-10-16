package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.1" )

func C16oDDI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5dyQ3NrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJW49T22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9dkkxMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llEWebbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bKdiV63kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXG9llNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YiY50ddmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APJ3w1g6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VsqrL62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSHZTazgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WwJPZweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WuEnfwHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r16BTTRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMEVV6GyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3HR1V75eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tCCjJjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NI2l6PAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9E8HhSXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64YulvOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func buLvADIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FtBcOgYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yANEGK3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rp4076EbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozIVph8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51KlypemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jojpvzpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yv1t26IlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8HzMI5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7aQT2omqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsCkCgl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BnxbNieWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UP1bE6QFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQq0TRuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NavWaAoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35FjQrrjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sG5oeZzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTWO4pjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eMIROoF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSSTY4xLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QmHZgIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DU7UeEVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j2s7M5kmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSCDmHLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHS6WZkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ZZL2UUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzbQJe6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afpsFwBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2ex7b2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func by6fZOGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OR7Ygl0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rebAiSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCZB5R8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Q0c8RsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9izpORnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DvbT5WXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NelJkkgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oVBEueTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQLg5F70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkkdEm03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBuXKfbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HjidogYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 74gMBWuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YOAYLYtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7Qd0y9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jH7fbBtVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqSJl7kLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MUPtrrWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhF5PxjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewELIXPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFvnlaPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywa4JzV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func irvS32rWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2BJAU1rdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hb2Knu1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NcGp6zAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywBC5LpnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0l4TpYWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDxyitMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zScZklbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yAbnoFdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acpt5RVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mrWtSn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WfHtE735Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 15sFlmLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSIWTMl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Esql8fkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RgCj5U7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wb5mw2qXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdzkwRMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8M7XsroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CpcOmniIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vGISmENOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YtMH2fzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 665AcD3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owXrWWK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVqQ9FdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTLJP8yvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNKxYCqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yalNqWbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mlD7UHIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0CJT8vcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHxce508Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVhIsXxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYvzSIZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vA7vLS7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2ZWSDVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EVpws8izWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ie4ULZgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lRI3Ify4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHmUGbmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BbJ3SDbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SAilmhVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFiSsqfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXQvEhIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKPaBiPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6SloM7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAYpz4gpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1R4wy9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmXM2Lm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oB25ioeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4FkpDEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGIKS0N0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PkBR4tirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 79EXYkTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3JjGneKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpM5Pl6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7teGZXQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkGH0AzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njIVxETEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ept5bmOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qvBuREdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNsFbJ4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ai3aXkZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9qr87bXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIoqCw7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7LqEkXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uy2ffG78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlG1Wu3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGvslk6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y88exe0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KkZyLHCtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbf6Glx3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ytI3g9aIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uqn53U4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Io71rDchWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsbBRW6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uf3wSJLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kjk6F0vDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dnVfUikeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lMqJwqMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVmyjetYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VoQDbnZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyYAl6GtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDTuRVjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8PnYqaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTPeQWXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mOXnN77KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kztwGdP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U530OoA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zzlYnF7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhChn8VUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bYyG4kHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPcrfeMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEj0pz9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IBLROyTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cvlp5mmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvcavuMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUSSrelIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5So02M0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q5y930j4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92AHcSBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBLsijJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UXmsSgQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e09UDn7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zUoMDW0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tb92rBH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9fqfk3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h93ICjwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func watYrKcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mlt7qGe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diBvwp44Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9H53tXE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jX0dzLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U4eJ9pZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmfCmLW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F3a7WAY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mkvURyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p85Zka6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkqI88TRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iXHqwR5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vsu9rBTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I1GRv6oOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZY2zXed5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zocwSyXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvlnMSFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n9BgflNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LaFaI0G3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2q80epvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzLVxDSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hljFBtVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5czEfzkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sc6rdrZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLGIY8yEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z6lPwVd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhZwus4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFexwSDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lushJxJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5CmKoRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PbVIXjiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfdzER34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUWAB3GXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmJLRuumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UDewfwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kPFgpqJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7TjylkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKVHLN3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zyM0F2i1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwfnfjHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IkXeMEOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8j1YDyrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aF3FOgaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OFSbsaG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCSU1DNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIo7l5avWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RjVqZD6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OaKfkE0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7wqFvwlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92Kl3yDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKKIBcFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0A5jcNyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hICCmM2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SI4jfPgtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdYHCBwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ssPPITOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EB0b1ecaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIThcz24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBF0ZX7NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VligZod7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wU54GpuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zhaueud4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPxrzyiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Go3vaUY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nRudEHwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CisdDCg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5D1rvmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4MqIsfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T44j6REAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1O8nHx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kdWEQRZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYYLfXRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kotx8pGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WiOgEH32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WAjtyNiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X7aGxo7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIsqSYBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Tjz5VDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qcK0sE3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVWgUmTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9abFUpmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCVdCn3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODUSb6FWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0q59yQZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func POYLwVNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DVOmSAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Y3zrt4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIs7A7zPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TciLnzDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nQbSwEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zqo6j1URWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8wFVEL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7QPhFwDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRTv7GZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TphuNIF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCmbb8LoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fb8WBHBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBObwQZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wZ0SyAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rNwnlmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l5udnnIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0K0ElOELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ixs1K13QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evvpmSziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wg1GYLcJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJEFxeDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2J0BpYanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cob85JBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tscWmrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IISAT4QhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPpS2uhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1lRgDIocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2geWS2xeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ObF4x5QsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func shHCi0anWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JtWLgzE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdH2tVDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GQXg5wIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDrZlfryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLbHqHH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgSJY4QuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRZc3JolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33kk200JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNe5M8kDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymBvCNUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWfz2gHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Go6mdMAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCRrqWG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uu4ODeQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4ybo1ZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWDBmSo4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6PjtEJs7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpfWC9KbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d79JLbI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqGTuK8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsma1kDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbmpKO6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtMFFmVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LzwSjJ3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mnpxiEEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sdv6tJ1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6F8zKAlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UxxuWEMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvAROWkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XyiLynicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xf1mH5UPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUmBbI3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k9xlIG70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btJrlSvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3PdOBkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjPXkfK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7HKac6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ryFskFvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbUHk7grWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TwoaN5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ufOmKyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8rxDyk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8rQqDzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OeKAoJHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbCJ1WytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5MzgBPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4JtK8yNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KaRLGxTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsF7akSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2INpK60zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iiMc8jS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWjJhgMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXe4in3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grLengQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 94MvrFnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWP1ChHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPCSfVf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fbG3ekkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMA8ErRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 245wbr0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWXD3F7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qp8pOiQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fW6f9lY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C1PrjEFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6cgUEZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCrKpV2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ihTr6EgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPCMbvnZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hm0HMataWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MpaetVtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1M0IYYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWtsyfPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQ2bbVe4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAofvjtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBevOWXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TavCtvhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ca0mty9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ju1ReswZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GgbmPSASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITD7uRePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UnTYCbXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APB1hBgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rcZzJyrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzqrFEEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wDOxOg29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ShkF8p6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ON9e8aqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1ByOaSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func usjfaIDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otPuzjDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BiSrDL7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qtZesBbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GePPlRNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C1494rDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p1KViIh3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IB2bd4GAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HB9SUTn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XEaMLSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q83w2KAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpnKLBAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8kP5oxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q3RvovV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWhvPGxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aihu0ZoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cAUrqw43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAc4DM7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPRNJW9TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VToRC34NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9stO0ECmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jqrkHSdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func neH4AEKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fh0J3gadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LIGZidEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func do5iRvjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vxDOJBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNiHq8QmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4JOW5keWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Wg1IQyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8pHFa4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8C1KSZkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CtrHRdyhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QnkEsIVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fg47Zvl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KmFL3uSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2FUAjebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lzliHw9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHfjcjwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8wcZrUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJfJdLF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jcokbStOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TpUoslKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFXPyC9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYUNC4NQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUaLO2jpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHuoXqMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgsMAOQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ectvWotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ql3n4SSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4kcV8fm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UrFVCG8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsJUhkZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dg40eK4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sv7Xe5L0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vKUBhDfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tYdss9yTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgMrps9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sY6JRc3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6Hk4DCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func laCQVIr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9hnQYsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89uwU1XqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoWoAYZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8MmHNWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBFUzeQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P33RYY8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUuFvHcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NKYpTfd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v04fIBN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YR9Fua6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKnC6rTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jqs2KFAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dG915T67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3JxjVdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fV7ThsIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDyqObIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWYm9rDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UvyMfOyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UdduzGaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gUsYyBHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZOT2gV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gINqHTdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKFsEFt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OS1VjjfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUxXh0PqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bm40DeCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2ye0O5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRWJQvJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptBzFOdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STrNckr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TidjQkI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qF1HQA8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDoX8hx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uevJ3WiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zgJiL9ECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G33pZpRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VwnToNxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12AtFL49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func edV3qDWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgjZ3BD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHWDfqc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNpQgzJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FvFLwijHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rt5Bu54HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Th3TUfPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQVcyLxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2K5epWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func deJfeZ1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xKYeOga6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhdRKBveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5d8p5fAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbpY1AO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lckp2Rm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmpc3YqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rB4EWYbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzrWwW1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVmgWMBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IvPz5p99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmVhPWQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XNjbZeqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xl81Z1aFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2sOtfJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cRFOs99GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRbnIvUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dmeBIm5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQLeyNBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3GwLOkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rmcopew9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func krh4fe9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WA00bUZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZzCK0y4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bF1GKkMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SdVEnJyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3T59fmBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhmgWEYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVFqcwWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ci3yXr1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1UYuomZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDJXeS2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZDFIFYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bghpOLJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2LLYTLiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3faoa49BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVYRla3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lo8kxaWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rentHPvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2SmPtt2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HuuV57ufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gfue2y2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HRrSmpZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PRGvA6K5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XrJO5GCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3GTCGKmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUSFv1ygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUqaaDaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvfWSRZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6E6AbBE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mk8bjftOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0pgKXxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAwwXAxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJ3BdY5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8XYRdncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVyZE5aLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ea3TQ2AJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QKlcl7nHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JyExHh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Os6AFfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5xdSmz2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzFwSuFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlkSjNy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NVNqEtZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDPLBsjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilZd1uopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEBJboUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTyul01MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZaJiCQFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5y5GFhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WwBAa1TuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPQAhVI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yxKTGKBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3R0gYHxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2bHqd3JJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func itkG9uHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1r1dFu4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ohr249Z0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4SVq2TTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0MiSMyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptUQsyvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jz9fO0KSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D4LF6jcFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5wscCA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmtzrqH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWJfy4pEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0X6nzqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jxy8WHJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cSt4Zn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5aton3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W5B4qV3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEam5OUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KY6mRyMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDihPmcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tf5w5P2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5YW4lyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MX1DjYZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rrseEOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Crg4j0m1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func myfjYmVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IBkn6ahuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeHjczRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hQbTSfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oeziC5DfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bU52koPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h26Ldip3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCvCRITNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UcDKi7vmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wl12bhV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLux3hdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2GJMgDvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJwJ7LFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCsUiQmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ukJPGGAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vP02F7hFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzQ2gpPJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8nRSQnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kYgLzSyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8h29VGTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Ik5O8xjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lOac2DmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPm3ROrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKXmHVM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vvgUfrCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFMqIm8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5CWjoglWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRiaX1K4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cMPOhPlKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCWeWftBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeVLjJcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QwivhHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IgUhLqdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tfJ9W4D9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TADmgmTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1VkFLtKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KFpe5a0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCbsrQBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OljVjcn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ulp86sHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4sp3bcK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uWO6J8auWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Egi6Dk96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khNGG1UIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJafOjrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p5Awj9mlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tMddEGbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kn83OzwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnQpTahQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fw6d1TxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqNlxh2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n72LKj1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNx24WTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvreW4BsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9FbmNNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yeod93xFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1XVpw7PZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUpG9u86Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MFuh7aaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13Vk8gutWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func euvHoZt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4mOOneu2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAggHmXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOrxa8FsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oT54kBa4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Z6nTnQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dH1KoMcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTjnggUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBBgRTG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYJwhvPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vesyzjijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwquoCtJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wL1ara6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufX9gWGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hc4nHpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWBPrCNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvDfO93wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qaoRW1j6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SnyJ7IgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLARkSO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bf30F3csWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VtAjUj7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3GlzfafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func geB0qVmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1f9khEn7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZbNZS5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyLUDZNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBCPg5G2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WrvphR3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QF7knlhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9KJHzceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOqeAMUoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLOO6XewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqxDvoEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2j2BUz0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r87l2KhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OkRmNyeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKHk0vpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51xc4DMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lwr7OtyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OlKq8XRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5U3eJUFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEjpFn02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9OexvUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WIi7kr7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1utn5t8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvEYAPh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuDpe7mIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2V72FKH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OLH1ZmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IaJnoBUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVf4Ou4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WKc3HqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z1sdsUHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9XUsOWByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OAGFw2GfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gh1kPkYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YiI1ehg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8oX9ddnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2TOXcU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bxvyLFL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H26zv8LXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8c6rccrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFnOZdQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRxD6d2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jtmf97e8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVkL1yxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cU7KuqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FkSAKQjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 88ZGF90OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2k6cep4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCIMldxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W691hgSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0yuZuOAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zya8ykp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEiAZZ6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HmV7ciByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2tL5BefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w65uMzWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MG1NQnLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVYH9kWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z8oiq0AFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bf2xkkmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHvSrQD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1LQCMxk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GQf8zEARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6FrtKfPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgWyZ0rqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZfLLwMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4Xu4AE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VG9bcDw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zj3EQgFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKVnOlWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCY0SDspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NG2qv9JjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ouz0BcShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5apzzEDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rM6zui4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UoLW4mIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func reCocQ5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ma7ZQFICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PW6VWrxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJvM467zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gfwJucsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g97gzh95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6akb5v8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TPUVw5oJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7UeyvUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwxJkraBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qk4Uo2dQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func opJUVw6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tEOXxUxGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYijxv03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k96u2LeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ibDe03rhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5Zn5WARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func imAKmEhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nXrJgKvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6COFOhUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acVeQFKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func st0wc4zDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgBR8l8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSWV7JKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8IN42YtbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ify1DLbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mM2LNIWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJeNrTZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDbcSflaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pc5xeL6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dYsWKPbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o79mGgXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lZufLe76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQsOBUdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iAa4hZQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xzrWWCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func siisyFLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLxCYHZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAEf8GYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5JvEpQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjQXHpJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQd2XMyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23t32eqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhKissmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGuggjeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3ar9Es0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8tp3ElNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMeLPd8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btQjocBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kIhxyoyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Bf0r5kTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rUqVSggpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0w1Ebf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P92rrI2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qW7yWiZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ynQjmnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kAyO61tVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8He64CMdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z1WkWCD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjqEYevpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func excsqbshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DP7BxLwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2O19hgu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76ATnDtiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdp5kJyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tUZVqnp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Eos2nkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uuc5DOlzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nf5uiWMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgJfPXdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jEnfiKT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PKkfBe5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uq1KVgjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e9oKhSuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMy3EtSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBzcbKqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SaveH09pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dyw70piXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cu8JlbOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VAOBRrZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mBRWLJHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNRz0RarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCz1732IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIXuaNWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Fdt9wvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQyKkBJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDMZrGlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLngG0glWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQOcvR1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JgzJ3gMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAhMbtkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GpfBIKNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q6xMsZudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gplAKR51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FltrNoOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFktyEnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func expCeptMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 615O5HIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2iwVbUk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ppknV3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3n8dQg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRLbfepVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghNTD4y2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGsUdFLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVB96A6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhuKiuyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vSG8dK0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ncru1VBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FBrbQQuaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8kUdU3ggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYAR6uvrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wD8S7AjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yB17Yq2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OuK0WA38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func arzN9fxcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDmWkLaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzQsFNk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFVVqEOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rqTOXmUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5kKOLnqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZUnP7327Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QkPzF7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1dAo8F5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dD7ENA2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vs6QgGyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMMW3eXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZI2NrzpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsdCbdLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtJ2mf0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZrmYcOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHgFMgFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 37KHhzRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4L3XHCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eguBstejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8GEt6LdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xfj104bBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DTf5HDvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxQPzp8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmMDqvmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxFrBmbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A8sx7j85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpsnK8PKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFb7eEPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmPADOn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func el6RufkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CGf99iHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKyYbbWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uaWmzrzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3tUrE24nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Y7N9WGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OgFcQD7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5moRX8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jc68nZMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZAVzb79yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2blxYQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmLsjrNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func poAlatp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSBaqfhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Br69rWwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CekZ97u2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvyznoHpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIU3eQ12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dgv1OkjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gMBTpEOhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SbfRgZ2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKuvOUd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lyt74ixBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qalYtQTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DCzlt1BfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqoN3WzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T1l9ZiJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkTCQL0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9kxgq4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVsojXf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HA5WDthxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U6KaeEj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3UnpHvSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTjA59iNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zEOpa7QeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhY9Gf2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wj7HKGN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qbe8q1YsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D8WhB8CqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0paVGul9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RukTsyF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AB1xVKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8O4mCo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeOhUF3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vh5UnFzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bcwpCAgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bz8uiq3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3byPC4PDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msqegshbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DAxL5AgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4QR2EdWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VETsfgd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHSkf4HsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FFs6y7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GANmMaZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S98upgCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SeUij29sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SZEaIJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5G9oJ5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2JfMGoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNyJ2DMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vSsZw5KNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvBdyvdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GkvqiHxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lqGh2C2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGwVdt0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func poQruHV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNcqlm1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FAsF1NjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wM3fiMXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLJxtJiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qX6UIpk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aiWtboDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBSkZRnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7WnUIc5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdOKMhNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S97DO65JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULxLAca6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnmoZWPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ynwkPsYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSzerLRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82pPm4YcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLcb82JqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hnpj73nyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TdShK2o6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OGtOHw2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EnAKItomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qpa2IB4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A69Q5jh3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func On3OgUhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrxJmQH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXUvPVkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9dKPYjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5NUDDKqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rUeebH8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nv47SzayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoPM6szsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BqmOi0SRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pobdAOiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydFpy4AxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yv7o15ysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func klDRHTd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hP8UDAFYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GMKJG8PcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncSajU46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9X9oWiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1GSbkKPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tT2nW0agWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LxRgOupWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NzXiESTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gwPteCwzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UfZhFLYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bN9Fm2mJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SHQeXKsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YEkKNErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WjPjGTjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kp0CzHldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jLN9hUM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPBTOoJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNb5CotiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUWKxloqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEzzTIoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8RKe6DgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcUY3NHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nORMXe3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y2CJeDQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YQKWnM1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Roj4jKJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Hwx3TzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcVasByxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdaORWfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BseqT8ITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7rXmFJdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDnwr4zUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tOf9La38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGckA0GGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJqAO2cPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFroX7I3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0HW8xfW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDlObO5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mTTyOXuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zb53o3pUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1AwYpqtiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKd9yR74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VyaTDZr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jeb1vUqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUOylHovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vH1yRFKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjgl1rSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xV5aE5OzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGEORPUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kO33Dp73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxoL1VOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfo9HcG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tf68isuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRyIrS3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVwjfFHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gguCT1ySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q6Qd0UXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMzlFwN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWlWwHCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUgAiMzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g9WBKozkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VvdXblCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50v7cxDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKoEMfmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zv4DZzboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ReOi3v7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7z1qPl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JBOqNnRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3UczbZfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjhduhEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgbslLanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0oRqljaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1gB3LMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJnonXfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TAT2oHriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4PW67iyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4YJa6rg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7WN2hd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjzrFmffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmTPOSXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvZnJzA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLeud3aPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yx1n7v1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JfWF49pFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35oxwMVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qoa6gY5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdETD7KMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8vnrnIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cp1ljPoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXFpkGAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGIC7h0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DuAXaOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GQg7p5ITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nL0e3FJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNFGLn3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lRVeTdt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5MdklSxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lcp9GlvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OUedoadNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CvjivtwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgIh116QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vG8w2RICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OWC4b0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qANVgA0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KtNaDVHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F64KXah8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FW4VYqfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzWsVxq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func al4gOtkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIolKfyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQClfnt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwX03b3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PR6xfUuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TA6Pvcz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mdgdin0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydwwMMOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgO7SBBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8sySFJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsqtJHafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBqwLvXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcfxpANpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y1BIwEtcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6tnrIVQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RSupuiPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6bTfwdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ik8fR8oKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mBvJfoi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h2WOFzOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GIYcFY6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZYn3aWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8k8KateJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Y9mUQKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C95CPBUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gaHYjL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hf8yy9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2Hr4JkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IsKEX01FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZcB55ZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbYfWDbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4NlXlcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ubs01jBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cj34xoNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWeX6tIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOYpfu5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LE0zZLWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func friloqZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZHxiMgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWf5AhjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9UqzpVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6vSJHZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QBgKXO5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVPnqxGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OpmQsAqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TsdbZq0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHLjIW69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GOTUJPgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQtlKEXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NsPeVfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFKo2WTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqULTCQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbHrQwiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3E7Mws6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NWAjsMbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2S73TVvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQuYX2kvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c1F6Nil5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sHLch9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwNNzEWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KixhNXcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBJr9gQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func klBjxKFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YUGEWp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Zui89RiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lId0e1PVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vydRpVfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psfXTGseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwOIpciTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ou9KOW75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1uoUHrrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zA3TLkORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h4fATCsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oRhwRrjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func shH8z0dvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fw8vn9YNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrDQmj0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gepRrMTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnw9YYMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMYntj9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJMF5zXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 72oVUsQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAaBRRHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpO7GGsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qb7KxOB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aiLiUwmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Mosaw1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V74L6DMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6BBh4qu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4u9JgkubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQEjBSyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0rtfob0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQRbd9hzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ugk9NkeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YaYFLqfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHIrFGVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5iXsKTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3Qt2UWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MKdusiHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qoVd739FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U2ZoA3FOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bl2J90krWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wrTlz4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vv7rD8XVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SeA84UtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peqkDO7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWGhu8nKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xv4Kh08ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUDK9uSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TnfeTV4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22tR0ltfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BiTVdmJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Zuh1zjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLsHmAU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mpDtFnsQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hgXonRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BiIC48r8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhm20CRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7q4HFKO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AFYbiTveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTjscZ0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASSMA0iXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2ntWqNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4WO4zDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmLTqclZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98NKgmHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOqtedxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKnA5lyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0UmjCJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rKWLozTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6c7MRsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P83mMWW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBwiaLeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func osm3ZQsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CS6WQEYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Heo9ZTfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5CQkiQ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hAbUlavLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRSL0crtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Woy8ABr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RzLId0QXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

