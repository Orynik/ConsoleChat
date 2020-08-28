package server

//Storage : Класс - хранилище
type Storage struct {
	Messages []*Message
}

//Storager : Интерфейс для записи данных в хранилище
//TODO: Вспомнить об пустых интерфейсах и их предназначениях
type Storager interface{}

//Message : Структура для хранения сообщений
type Message struct {
	ID     int
	Author string
	Text   string
	//TODO: Под вопросом TimeStamp - int64/string/time?
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
