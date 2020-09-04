package server

//Storage : Класс - хранилище
type Storage struct {
	Messages []*Message
}

//Storager : Интерфейс для записи данных в хранилище
type Storager interface {
	Save() *Storage
	GetAll() []*Message
	GetByID(int) *Message
}

//Message : Структура для хранения сообщений
type Message struct {
	ID        int
	From      string
	To        string
	Text      string
	TimeStamp int64
	Lang      string
}

//NewStorage : Конструктор
func NewStorage() *Storage {
	return &Storage{
		Messages: make([]*Message, 512, 1024),
	}
}

//Save : Сохранить сообщение
func (s *Storage) Save(m *Message) {
	s.Messages = append(s.Messages, m)
}

//GetAll : Получить все сообщения
func (s *Storage) GetAll() []*Message {
	return s.Messages
}

//GetByID возвращает сообщение с переданным id
func (s *Storage) GetByID(id int) *Message {
	return s.Messages[id]
}
