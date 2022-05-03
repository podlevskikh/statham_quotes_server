package hashcash

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "060102150405"
const zeroByte = 48
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Hashcash struct {
	Ver      int       // Версию hashcash, 1 (которая заменила версию 0).
	Bits     int       // Число  "предварительных" (нулевых) битов в хешированном коде.
	Date     time.Time // Время, в которое сообщение было отправлено, в формате ГГММДД[ччмм[сс]].
	Resource string    // Данные об отправителе, например,  IP-адрес или адрес E-mail.
	Rand     string    // Строка случайных чисел, закодированная в формате [[Base64|base-64]].
	Counter  int       // Двоичный счетчик, закодированный в формате base-64.
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewHashcash(resource string) *Hashcash {
	return &Hashcash{
		Ver:      1,
		Bits:     5,
		Date:     time.Now(),
		Resource: resource,
		Rand:     randSeq(10),
		Counter:  0,
	}
}

func Parse(hash string) (*Hashcash, error) {
	splitHash := strings.Split(hash, ":")
	if len(splitHash) != 6 {
		return nil, errors.New("parts count invalid")
	}
	ver, err := strconv.Atoi(splitHash[0])
	if err != nil {
		return nil, errors.Wrap(err, "version format")
	}
	bits, err := strconv.Atoi(splitHash[1])
	if err != nil {
		return nil, errors.Wrap(err, "bits format")
	}
	date, err := time.Parse(DateFormat, splitHash[2])
	if err != nil {
		return nil, errors.Wrap(err, "date format")
	}
	counterDecoded, err := base64.StdEncoding.DecodeString(splitHash[5])
	if err != nil {
		return nil, errors.Wrap(err, "counter base64 format")
	}
	counter, err := strconv.Atoi(string(counterDecoded))
	if err != nil {
		return nil, errors.Wrap(err, "counter format")
	}
	return &Hashcash{
		Ver:      ver,
		Bits:     bits,
		Date:     date,
		Resource: splitHash[3],
		Rand:     splitHash[4],
		Counter:  counter,
	}, nil
}

func (h *Hashcash) ToString() string {
	return fmt.Sprintf(
		"%v:%v:%v:%v:%v:%v",
		h.Ver,
		h.Bits,
		h.Date.Format(DateFormat),
		h.Resource,
		h.Rand,
		base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(h.Counter))))
}

func (h *Hashcash) IsValid() bool {
	s := sha1.New()
	s.Write([]byte(h.ToString()))
	bs := s.Sum(nil)
	hash := fmt.Sprintf("%x", bs)

	if h.Bits > len(hash) {
		return false
	}
	for _, ch := range hash[:h.Bits] {
		if ch != zeroByte {
			return false
		}
	}
	return true
}

func (h *Hashcash) Solute() {
	for !h.IsValid() {
		h.Counter++
	}
}
