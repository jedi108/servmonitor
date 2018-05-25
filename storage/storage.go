/*
 1 - filename: берем год_месяц_день
 2 - Storage.day: сохраняем день
 3 - проверяем Storage.день = день
 4 - ПишемДанные (
	   если День = день, то пищем
	   иначе создаем новый файл
	)
 5 - создаем файл если нет

*/

package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const (
	templateFile = "%v%02d%02d.csv"
)

func IsFileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetFileNameFromDate() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf(templateFile, year, month, day)
}

type Storage struct {
	fileName    string
	startDay    int
	fappend     *os.File
	IsDebugEcho bool
	csvWriter   *csv.Writer
	// fremove     *os.File
}

func NewStorage(fileName string) (*Storage, error) {
	_, _, day := time.Now().Date()
	storage := &Storage{fileName: fileName, startDay: day}

	fA, err := os.OpenFile(storage.fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fA, err = os.Create(fileName)
		if err != nil {
			return storage, err
		}
		return storage, err
	}
	storage.fappend = fA
	return storage, err
}

func (s *Storage) InsertData(datas []string) error {
	_, _, day := time.Now().Date()
	if s.startDay != day {
		s.fappend.Close()

		year, month, day := time.Now().AddDate(0, 0, -2).Date()
		removeFileName := fmt.Sprintf(templateFile, year, month, day)
		if IsFileExists(removeFileName) {
			os.Remove(removeFileName)
		}

		ss, err := NewStorage(GetFileNameFromDate())
		if err != nil {
			return err
		}
		s = ss
		defer ss.Close()
	}

	s.csvWriter = csv.NewWriter(s.fappend)
	defer s.csvWriter.Flush()

	// []string
	//
	err := s.csvWriter.Write(datas)
	if err != nil {
		return err
	}

	// [][]string
	//
	// for _, value := range datas {
	// 	err := writer.Write(value)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (s *Storage) Close() {
	s.fappend.Close()
	//s.fremove.Close()
}

// Remove first line from file
func (s *Storage) PopLine() error {
	// os.OpenFile(storage.fileName, os.O_RDWR|os.O_CREATE, 0666)

	// fi, err := s.fremove.Stat()
	// if err != nil {
	// 	return err
	// }
	// buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))
	// _, err = s.fremove.Seek(0, os.SEEK_SET)
	// if err != nil {
	// 	return err
	// }
	// _, err = io.Copy(buf, s.fremove)
	// if err != nil {
	// 	return err
	// }
	// line, err := buf.ReadString('\n')
	// if err != nil && err != io.EOF {
	// 	return err
	// }
	// _, err = s.fremove.Seek(0, os.SEEK_SET)
	// if err != nil {
	// 	return err
	// }
	// nw, err := io.Copy(s.fremove, buf)
	// if err != nil {
	// 	return err
	// }
	// err = s.fremove.Truncate(nw)
	// if err != nil {
	// 	return err
	// }
	// err = s.fremove.Sync()
	// if err != nil {
	// 	return err
	// }
	// _, err = s.fremove.Seek(0, os.SEEK_SET)
	// if err != nil {
	// 	return err
	// }
	// if s.IsDebugEcho {
	// 	fmt.Println("pop: ", line)
	// }
	return nil
}
