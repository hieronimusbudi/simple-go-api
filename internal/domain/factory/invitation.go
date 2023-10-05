package factory

import "github.com/hieronimusbudi/simple-go-api/internal/domain"

type Invitation struct{}

func (of Invitation) Generate(
	invitations []domain.Invitation,
	gatherings []domain.Gathering,
	members []domain.Member,
) (result []domain.Invitation) {
	mapMemberByID := map[int64]domain.Member{}
	mapGatheringByID := map[int64]domain.Gathering{}
	for _, m := range members {
		mapMemberByID[m.ID] = m
	}
	for _, g := range gatherings {
		mapGatheringByID[g.ID] = g
	}
	for _, inv := range invitations {
		_, ok := mapMemberByID[inv.Member.ID]
		if ok {
			inv.Member = mapMemberByID[inv.Member.ID]
		}
		_, ok = mapGatheringByID[inv.Gathering.ID]
		if ok {
			inv.Gathering = mapGatheringByID[inv.Gathering.ID]
		}
		result = append(result, inv)
	}
	return
}
