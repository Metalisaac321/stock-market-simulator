package accountMap

import "github.com/Metalisaac321/stock-market-simulator/internal/account"

/*
  public static toPersistence (vinyl: Vinyl): any {
    return {
      album_name: vinyl.albumName.value,
      artist_name: vinyl.artistName.value
    }
  } */

type AccountDto struct {
	Id   string `json:"id"`
	Cash uint   `json:"cash"`
}

func ToDomain(raw AccountDto) account.Account {
	a, _ := account.NewAccount(raw.Id, raw.Cash)
	return a
}

func ToDto(account account.Account) AccountDto {
	return AccountDto{
		Id:   account.Id().Value(),
		Cash: account.Cash().Value(),
	}
}
