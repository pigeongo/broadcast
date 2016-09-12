package broadcast

// 广播
type broadcast struct {
    C chan broadcast
    V interface{}
}

// 播音器
type broadcaster struct {
    L chan chan (chan broadcast)
    S chan <- interface{}
}

// 接收器
type Receiver struct {
    C chan broadcast
}

func NewBroadcaster() *broadcaster {
    s := make(chan interface{})
    l := make(chan (chan (chan broadcast)))
    go func() {
        current := make(chan broadcast, 1)
        for {
            select {
            case v := <-s:
                if v == nil {
                    current <- broadcast{}
                    return
                }
                c := make(chan broadcast, 1)
                current <- broadcast{C: c, V: v}
                current = c
            case r := <-l:
                r <- current
            }
        }
    }()
    return &broadcaster{L: l, S: s}
}

func (b *broadcaster) Listen() Receiver {
    c := make(chan chan broadcast, 0);
    b.L <- c;
    return Receiver{<-c};
}

func (b *broadcaster) Writer(v interface{}) {
    b.S <- v
}

func (r *Receiver) Reader() interface{} {
    b := <-r.C;
    r.C <- b;
    r.C = b.C;
    return b.V;
}