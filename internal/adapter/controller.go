package adapter

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hieronimusbudi/simple-go-api/internal/adapter/docs"
	"github.com/hieronimusbudi/simple-go-api/internal/adapter/mysql"
	"github.com/hieronimusbudi/simple-go-api/internal/adapter/repository"
	"github.com/hieronimusbudi/simple-go-api/internal/application/usecase"
	"github.com/hieronimusbudi/simple-go-api/internal/domain"
	"github.com/hieronimusbudi/simple-go-api/internal/domain/factory"
	"github.com/hieronimusbudi/simple-go-api/internal/domain/valueobject"
	"github.com/hieronimusbudi/simple-go-api/internal/helpers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Controller is a controller
type Controller struct {
	MemberUsecase     usecase.IMemberUsecase
	GatheringUsecase  usecase.IGatheringUsecase
	InvitationUsecase usecase.IInvitationUsecase
}

// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	db := mysql.Connection()

	memberRepository := repository.NewMemberRepository(repository.MemberAdapterRepositoryArgs{DB: db})
	gatheringRepository := repository.NewGatheringRepository(repository.GatheringAdapterRepositoryArgs{DB: db})
	invitationRepository := repository.NewInvitationRepository(repository.InvitationAdapterRepositoryArgs{DB: db})

	memberUsecase := usecase.NewMemberUsecase(usecase.MemberUsecaseArgs{
		MemberRepository: memberRepository,
	})
	gatheringUsecase := usecase.NewGatheringUsecase(usecase.GatheringUsecaseArgs{
		GatheringRepository: gatheringRepository,
	})
	invitationUsecase := usecase.NewInvitationUsecase(usecase.InvitationUsecaseArgs{
		InvitationRepository: invitationRepository,
	})

	controller := Controller{
		MemberUsecase:     memberUsecase,
		GatheringUsecase:  gatheringUsecase,
		InvitationUsecase: invitationUsecase,
	}

	memberRoutes := r.Group("/members")
	memberRoutes.POST("", controller.CreateMember)
	memberRoutes.GET("", controller.GetMembers)
	memberRoutes.GET("/:id", controller.GetMember)
	memberRoutes.PUT("/:id", controller.UpdateMember)
	memberRoutes.DELETE("/:id", controller.DeleteMember)

	gatheringRoutes := r.Group("/gatherings")
	gatheringRoutes.POST("", controller.CreateGathering)
	gatheringRoutes.GET("", controller.GetGatherings)
	gatheringRoutes.GET("/:id", controller.GetGathering)
	gatheringRoutes.PUT("/:id", controller.UpdateGathering)
	gatheringRoutes.DELETE("/:id", controller.DeleteGathering)

	invitationRoutes := r.Group("/invitations")
	invitationRoutes.POST("", controller.CreateInvitation)
	invitationRoutes.GET("", controller.GetInvitations)
	invitationRoutes.GET("/:id", controller.GetInvitation)
	invitationRoutes.PUT("/:id/accept", controller.AcceptInvitation)
	invitationRoutes.PUT("/:id/reject", controller.RejectInvitation)
	invitationRoutes.PUT("/:id/cancel", controller.CancelInvitation)

	docs.SwaggerInfo.Title = "Gathering App API"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// @Tags			Member
// @Summary		Create Member
// @Description	Create Member
// @Accept			json
// @Produce		json
// @Param			payload	body		swaggermodel.Member									true	"Payload"
// @Success		200		{object}	helpers.ResponsePayload{data=swaggermodel.Member}	"Member"
// @Router			/members [post]
func (ctr *Controller) CreateMember(c *gin.Context) {
	member := domain.Member{}
	if err := c.BindJSON(&member); err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	err := member.Validate()
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	member, err = ctr.MemberUsecase.Create(context.Background(), member)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	helpers.NewResponse(c, http.StatusCreated, "success", member)
}

// @Tags			Member
// @Summary		Get Members
// @Description	Get Members
// @Accept			json
// @Produce		json
// @Success		200	{array}	helpers.ResponsePayload{data=swaggermodel.Member}	"Member"
// @Router			/members [get]
func (ctr *Controller) GetMembers(c *gin.Context) {
	members, err := ctr.MemberUsecase.Get(context.Background(), domain.MemberArgs{})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	helpers.NewResponse(c, http.StatusOK, "success", members)
}

// @Tags			Member
// @Summary		Get Member By ID
// @Description	Get Member By ID
// @Accept			json
// @Produce		json
// @Param			id	path		int													true	"member ID"
// @Success		200	{object}	helpers.ResponsePayload{data=swaggermodel.Member}	"Member"
// @Router			/members/{id} [get]
func (ctr *Controller) GetMember(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	members, err := ctr.MemberUsecase.GetByID(context.Background(), id)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	helpers.NewResponse(c, http.StatusOK, "success", members)
}

// @Tags			Member
// @Summary		Update Member
// @Description	Update Member
// @Accept			json
// @Produce		json
// @Param			id		path		int							true	"Member ID"
// @Param			payload	body		swaggermodel.Member			true	"Payload"
// @Success		200		{object}	helpers.ResponsePayload{}	"Member"
// @Router			/members/{id} [put]
func (ctr *Controller) UpdateMember(c *gin.Context) {
	member := domain.Member{}
	if err := c.BindJSON(&member); err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	err := member.Validate()
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	_, err = ctr.MemberUsecase.GetByID(context.Background(), id)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	member.ID = id
	err = ctr.MemberUsecase.Update(context.Background(), member)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	member, err = ctr.MemberUsecase.GetByID(context.Background(), id)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	helpers.NewResponse(c, http.StatusOK, "success", member)
}

// @Tags			Member
// @Summary		Delete Member
// @Description	Delete Member
// @Accept			json
// @Produce		json
// @Param			id	path		int							true	"Member ID"
// @Success		200	{object}	helpers.ResponsePayload{}	"Member"
// @Router			/members/{id} [delete]
func (ctr *Controller) DeleteMember(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	_, err = ctr.MemberUsecase.GetByID(context.Background(), id)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	err = ctr.MemberUsecase.Delete(context.Background(), domain.MemberArgs{ID: id})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	helpers.NewResponse(c, http.StatusOK, "success", nil)
}

// @Tags			Gathering
// @Summary		Create Gathering
// @Description	Create Gathering
// @Accept			json
// @Produce		json
// @Param			payload	body		swaggermodel.Gathering									true	"Payload"
// @Success		200		{object}	helpers.ResponsePayload{data=swaggermodel.Gathering}	"Gathering"
// @Router			/gatherings [post]
func (ctr *Controller) CreateGathering(c *gin.Context) {
	gathering := domain.Gathering{}
	if err := c.BindJSON(&gathering); err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	err := gathering.Validate()
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	// creator will also treated as attendee
	creator, err := ctr.MemberUsecase.GetByID(context.Background(), gathering.Creator.ID)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	gathering.Attendees = append(gathering.Attendees, creator)
	gathering, err = ctr.GatheringUsecase.Create(context.Background(), gathering)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	memberIDs := []int64{}
	for _, m := range gathering.Attendees {
		if m.ID != creator.ID {
			memberIDs = append(memberIDs, m.ID)
		}
	}
	members, err := ctr.MemberUsecase.Get(context.Background(), domain.MemberArgs{
		IDs: memberIDs,
	})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	members = append(members, creator)
	gatheringFactory := factory.Gathering{}
	gathering = gatheringFactory.Generate([]domain.Gathering{gathering}, members)[0]
	helpers.NewResponse(c, http.StatusCreated, "success", gathering)
}

// @Tags			Gathering
// @Summary		Get Gatherings
// @Description	Get Gatherings
// @Accept			json
// @Produce		json
// @Success		200	{array}	helpers.ResponsePayload{data=swaggermodel.Gathering}	"Gathering"
// @Router			/gatherings [get]
func (ctr *Controller) GetGatherings(c *gin.Context) {
	gatherings, err := ctr.GatheringUsecase.Get(context.Background(), domain.GatheringArgs{})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	memberIDs := []int64{}
	for _, g := range gatherings {
		memberIDs = append(memberIDs, g.Creator.ID)
		for _, m := range g.Attendees {
			memberIDs = append(memberIDs, m.ID)
		}
	}
	members, err := ctr.MemberUsecase.Get(context.Background(), domain.MemberArgs{
		IDs: memberIDs,
	})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	gatheringFactory := factory.Gathering{}
	gatherings = gatheringFactory.Generate(gatherings, members)
	helpers.NewResponse(c, http.StatusOK, "success", gatherings)
}

// @Tags			Gathering
// @Summary		Get Gathering By ID
// @Description	Get Gathering By ID
// @Accept			json
// @Produce		json
// @Param			id	path		int														true	"Gathering ID"
// @Success		200	{object}	helpers.ResponsePayload{data=swaggermodel.Gathering}	"Gathering"
// @Router			/gatherings/{id} [get]
func (ctr *Controller) GetGathering(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	gathering, err := ctr.GatheringUsecase.GetByID(context.Background(), id)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	memberIDs := []int64{gathering.Creator.ID}
	for _, m := range gathering.Attendees {
		memberIDs = append(memberIDs, m.ID)
	}
	members, err := ctr.MemberUsecase.Get(context.Background(), domain.MemberArgs{
		IDs:              memberIDs,
		IsIncludeDiscard: true,
	})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	gatheringFactory := factory.Gathering{}
	gatherings := gatheringFactory.Generate([]domain.Gathering{gathering}, members)
	helpers.NewResponse(c, http.StatusOK, "success", gatherings[0])
}

// @Tags			Gathering
// @Summary		Update Gathering
// @Description	Update Gathering
// @Accept			json
// @Produce		json
// @Param			id		path		int								true	"Gathering ID"
// @Param			payload	body		swaggermodel.UpdateGathering	true	"Payload"
// @Success		200		{object}	helpers.ResponsePayload{}		"Gathering"
// @Router			/gatherings/{id} [put]
func (ctr *Controller) UpdateGathering(c *gin.Context) {
	gathering := domain.Gathering{}
	if err := c.BindJSON(&gathering); err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	err := gathering.Validate()
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	_, err = ctr.GatheringUsecase.GetByID(context.Background(), id)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	gathering.ID = id
	err = ctr.GatheringUsecase.Update(context.Background(), gathering)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	gathering, err = ctr.GatheringUsecase.GetByID(context.Background(), id)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	helpers.NewResponse(c, http.StatusOK, "success", gathering)
}

// @Tags			Gathering
// @Summary		Delete Gathering
// @Description	Delete Gathering
// @Accept			json
// @Produce		json
// @Param			id	path		int							true	"Gathering ID"
// @Success		200	{object}	helpers.ResponsePayload{}	"Gathering"
// @Router			/gatherings/{id} [delete]
func (ctr *Controller) DeleteGathering(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	_, err = ctr.GatheringUsecase.GetByID(context.Background(), id)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	err = ctr.GatheringUsecase.Delete(context.Background(), domain.GatheringArgs{ID: id})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	helpers.NewResponse(c, http.StatusOK, "success", nil)
}

// @Tags			Invitation
// @Summary		Create Invitation
// @Description	Create Invitation
// @Accept			json
// @Produce		json
// @Param			payload	body		swaggermodel.Invitation									true	"Payload"
// @Success		200		{object}	helpers.ResponsePayload{data=swaggermodel.Invitation}	"Invitation"
// @Router			/invitations [post]
func (ctr *Controller) CreateInvitation(c *gin.Context) {
	invitation := domain.Invitation{}
	if err := c.BindJSON(&invitation); err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	err := invitation.Validate()
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	member, err := ctr.MemberUsecase.GetByID(context.Background(), invitation.Member.ID)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	gathering, err := ctr.GatheringUsecase.GetByID(context.Background(), invitation.Gathering.ID)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	invitation.Member = member
	invitation.Gathering = gathering
	invitation.Status = valueobject.INVITATION_CREATED
	invitation, err = ctr.InvitationUsecase.Create(context.Background(), invitation)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	invitationFactory := factory.Invitation{}
	invitation = invitationFactory.Generate([]domain.Invitation{invitation}, []domain.Gathering{gathering}, []domain.Member{member})[0]
	helpers.NewResponse(c, http.StatusCreated, "success", invitation)
}

// @Tags			Invitation
// @Summary		Get Invitations
// @Description	Get Invitations
// @Accept			json
// @Produce		json
// @Success		200	{array}	helpers.ResponsePayload{data=swaggermodel.Invitation}	"Invitation"
// @Router			/invitations [get]
func (ctr *Controller) GetInvitations(c *gin.Context) {
	invitations, err := ctr.InvitationUsecase.Get(context.Background(), domain.InvitationArgs{})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	memberIDs := []int64{}
	gatheringIDs := []int64{}
	for _, inv := range invitations {
		memberIDs = append(memberIDs, inv.Member.ID)
		gatheringIDs = append(gatheringIDs, inv.Gathering.ID)
	}
	members, err := ctr.MemberUsecase.Get(context.Background(), domain.MemberArgs{
		IDs: memberIDs,
	})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	gatherings, err := ctr.GatheringUsecase.Get(context.Background(), domain.GatheringArgs{
		IDs: gatheringIDs,
	})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	invitationFactory := factory.Invitation{}
	invitations = invitationFactory.Generate(invitations, gatherings, members)
	helpers.NewResponse(c, http.StatusOK, "success", invitations)
}

// @Tags			Invitation
// @Summary		Get Invitation By ID
// @Description	Get Invitation By ID
// @Accept			json
// @Produce		json
// @Param			id	path		int														true	"Invitation ID"
// @Success		200	{object}	helpers.ResponsePayload{data=swaggermodel.Invitation}	"Invitation"
// @Router			/invitations/{id} [get]
func (ctr *Controller) GetInvitation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	invitation, err := ctr.InvitationUsecase.GetByID(context.Background(), id)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	memberIDs := []int64{invitation.Member.ID}
	gatheringIDs := []int64{invitation.Gathering.ID}
	members, err := ctr.MemberUsecase.Get(context.Background(), domain.MemberArgs{
		IDs:              memberIDs,
		IsIncludeDiscard: true,
	})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	gatherings, err := ctr.GatheringUsecase.Get(context.Background(), domain.GatheringArgs{
		IDs:              gatheringIDs,
		IsIncludeDiscard: true,
	})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	invitationFactory := factory.Invitation{}
	invitations := invitationFactory.Generate([]domain.Invitation{invitation}, gatherings, members)
	helpers.NewResponse(c, http.StatusOK, "success", invitations[0])
}

// @Tags			Invitation
// @Summary		Accept Invitation
// @Description	Accept Invitation
// @Accept			json
// @Produce		json
// @Param			id	path		int							true	"Invitation ID"
// @Success		200	{object}	helpers.ResponsePayload{}	"Invitation"
// @Router			/invitations/{id}/accept [put]
func (ctr *Controller) AcceptInvitation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	invitation, err := ctr.InvitationUsecase.GetByID(context.Background(), id)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if invitation.Status == valueobject.INVITATION_ACCEPT {
		err = errors.New("the member has accepted the invitation")
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	err = ctr.InvitationUsecase.Accept(context.Background(), domain.InvitationArgs{
		ID:          id,
		MemberID:    invitation.MemberID,
		GatheringID: invitation.GatheringID,
		Status:      valueobject.INVITATION_ACCEPT,
	})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	helpers.NewResponse(c, http.StatusOK, "success", nil)
}

// @Tags			Invitation
// @Summary		Reject Invitation
// @Description	Reject Invitation, used by member to reject the invitation
// @Accept			json
// @Produce		json
// @Param			id	path		int							true	"Invitation ID"
// @Success		200	{object}	helpers.ResponsePayload{}	"Invitation"
// @Router			/invitations/{id}/reject [put]
func (ctr *Controller) RejectInvitation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	invitation, err := ctr.InvitationUsecase.GetByID(context.Background(), id)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if invitation.Status == valueobject.INVITATION_CANCELED {
		err = errors.New("the invitation for this member has canceled")
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	} else if invitation.Status == valueobject.INVITATION_REJECT {
		err = errors.New("the member has rejected the invitation")
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	err = ctr.InvitationUsecase.Reject(context.Background(), domain.InvitationArgs{
		ID:          id,
		MemberID:    invitation.MemberID,
		GatheringID: invitation.GatheringID,
		Status:      valueobject.INVITATION_REJECT,
	})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	helpers.NewResponse(c, http.StatusOK, "success", nil)
}

// @Tags			Invitation
// @Summary		Cancel Invitation
// @Description	Cancel Invitation, will remove member from attendee list
// @Accept			json
// @Produce		json
// @Param			id	path		int							true	"Invitation ID"
// @Success		200	{object}	helpers.ResponsePayload{}	"Invitation"
// @Router			/invitations/{id}/cancel [put]
func (ctr *Controller) CancelInvitation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	invitation, err := ctr.InvitationUsecase.GetByID(context.Background(), id)
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if invitation.Status == valueobject.INVITATION_CANCELED {
		err = errors.New("the invitation for this member has canceled")
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	} else if invitation.Status == valueobject.INVITATION_REJECT {
		err = errors.New("the member has rejected the invitation")
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	err = ctr.InvitationUsecase.Cancel(context.Background(), domain.InvitationArgs{
		ID:          id,
		MemberID:    invitation.MemberID,
		GatheringID: invitation.GatheringID,
		Status:      valueobject.INVITATION_CANCELED,
	})
	if err != nil {
		helpers.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	helpers.NewResponse(c, http.StatusOK, "success", nil)
}
