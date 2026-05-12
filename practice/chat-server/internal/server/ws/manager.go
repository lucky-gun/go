// internal/ws/manager.go
package ws

type Manager struct {
        Rooms map[string]*Room
}

func NewManager() *Manager {
        return &Manager{
                Rooms:    make(map[string]*Room),
        }
}
