package factory

import "github.com/hieronimusbudi/simple-go-api/internal/domain"

type Gathering struct{}

func (of Gathering) Generate(
	gatherings []domain.Gathering,
	members []domain.Member,
) (result []domain.Gathering) {
	mapMemberByID := map[int64]domain.Member{}
	for _, m := range members {
		mapMemberByID[m.ID] = m
	}
	for _, g := range gatherings {
		_, ok := mapMemberByID[g.Creator.ID]
		if ok {
			g.Creator = mapMemberByID[g.Creator.ID]
		}
		attendees := []domain.Member{}
		for _, a := range g.Attendees {
			_, ok = mapMemberByID[a.ID]
			if ok {
				a = mapMemberByID[a.ID]
			}
			attendees = append(attendees, a)
		}
		g.Attendees = attendees
		result = append(result, g)
	}
	return
}
