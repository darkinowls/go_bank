package api

//type createTransferRequest struct {
//	Owner    string `json:"owner" binding:"required"`
//	Currency string `json:"currency" binding:"required,oneof=EUR USD" `
//}
//
//func (server *Server) createAccount(ctx *gin.Context) {
//	var req createAccountRequest
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	arg := db.CreateAccountParams{
//		Owner:    req.Owner,
//		Balance:  0,
//		Currency: req.Currency,
//	}
//	account, err := server.store.CreateAccount(ctx, arg)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//	ctx.JSON(http.StatusOK, account)
//}
//
//type getAccountParam struct {
//	ID int64 `uri:"id" binding:"required,min=1"`
//}
//
//func (server *Server) getAccount(ctx *gin.Context) {
//	var param getAccountParam
//
//	if err := ctx.ShouldBindUri(&param); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//	account, err := server.store.GetAccount(ctx, param.ID)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			ctx.JSON(http.StatusNotFound, errorResponse(err))
//			return
//		}
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//	ctx.JSON(http.StatusOK, account)
//}
//
//type ListAccountQuery struct {
//	PageId   int32 `form:"pageId" binding:"required,min=1"`
//	PageSize int32 `form:"pageSize" binding:"required,min=5,max=20"`
//}
//
//func (server *Server) listAccount(ctx *gin.Context) {
//
//	var req ListAccountQuery
//
//	if err := ctx.ShouldBindQuery(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//	log.Println(req.PageId, req.PageSize)
//	accounts, err := server.store.ListAccounts(ctx, db.ListAccountsParams{
//		Limit:  req.PageSize,
//		Offset: (req.PageId - 1) * req.PageSize,
//	})
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//	ctx.JSON(http.StatusOK, accounts)
//}
//
//func (server *Server) countAccount(context *gin.Context) {
//	count, err := server.store.CountAccounts(context)
//	if err != nil {
//		context.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//	context.JSON(http.StatusOK, count)
//}