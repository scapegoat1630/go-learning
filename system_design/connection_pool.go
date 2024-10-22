package system_design

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Connection 是一个接口，代表一个连接。
// 你需要根据你的连接类型（如数据库连接）来实现这个接口。
type Connection interface {
	Close() error
	// 其他连接相关的方法...
}

// FakeConnection 是一个模拟的连接实现，用于演示。
type FakeConnection struct {
}

func (fc *FakeConnection) Close() error {
	// 模拟关闭连接的过程
	return nil
}

// ConnectionPool 是一个连接池的结构体。
type ConnectionPool struct {
	mu       sync.Mutex
	cond     *sync.Cond
	conns    []Connection
	maxConns int
	factory  func() (Connection, error)
}

// NewConnectionPool 创建一个新的连接池。
func NewConnectionPool(maxConns int, factory func() (Connection, error)) (*ConnectionPool, error) {
	if maxConns <= 0 {
		return nil, errors.New("maxConns must be greater than 0")
	}
	pool := &ConnectionPool{
		conns:    make([]Connection, 0, maxConns),
		maxConns: maxConns,
		factory:  factory,
	}
	pool.cond = sync.NewCond(&pool.mu)
	// 预先创建一些连接（可选）
	for i := 0; i < cap(pool.conns)/2; i++ {
		conn, err := factory()
		if err != nil {
			return nil, err
		}
		pool.conns = append(pool.conns, conn)
	}
	return pool, nil
}

// Get 从连接池中获取一个连接。
// 如果连接池中没有可用的连接，并且达到了最大连接数，它将阻塞直到有连接可用。
func (p *ConnectionPool) Get() (Connection, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.conns) == 0 {
		if len(p.conns) >= p.maxConns {
			p.cond.Wait()
		} else {
			// 创建新的连接（如果未达到最大连接数）
			conn, err := p.factory()
			if err != nil {
				return nil, err
			}
			p.conns = append(p.conns, conn)
		}
	}

	// 从连接池中取出一个连接
	conn := p.conns[len(p.conns)-1]
	p.conns = p.conns[:len(p.conns)-1]
	p.cond.Signal() // 通知可能有等待的 goroutine

	return conn, nil
}

// Put 将一个连接放回连接池。
// 注意：在实际应用中，你可能需要检查连接是否仍然有效，并在必要时丢弃它。
func (p *ConnectionPool) Put(conn Connection) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 在这里可以添加检查连接是否有效的逻辑
	// 如果连接无效，则关闭它而不是放回池中

	// 将连接放回池中
	p.conns = append(p.conns, conn)
	p.cond.Signal() // 通知可能有等待的 goroutine
}

// Close 关闭连接池中的所有连接。
func (p *ConnectionPool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, conn := range p.conns {
		if err := conn.Close(); err != nil {
			return err
		}
	}
	p.conns = nil // 清空连接池
	return nil
}

func main() {
	// 使用 FakeConnection 作为连接工厂函数
	factory := func() (Connection, error) {
		fmt.Println("Creating new connection")
		return &FakeConnection{}, nil
	}

	// 创建一个最大连接数为 5 的连接池
	pool, err := NewConnectionPool(5, factory)
	if err != nil {
		fmt.Println("Error creating connection pool:", err)
		return
	}
	defer pool.Close()

	// 获取和放回连接
	for i := 0; i < 10; i++ {
		conn, err := pool.Get()
		if err != nil {
			fmt.Println("Error getting connection:", err)
			continue
		}
		fmt.Println("Got connection:", conn)
		// 模拟使用连接
		time.Sleep(time.Second)
		// 将连接放回池中
		pool.Put(conn)
	}
}
