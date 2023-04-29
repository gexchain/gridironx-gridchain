package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gridironx/gridchain/libs/cosmos-sdk/client/context"
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/types/rest"
	"github.com/gridironx/gridchain/libs/tendermint/crypto/ed25519"
	"github.com/gridironx/gridchain/x/common"
	comm "github.com/gridironx/gridchain/x/common"
	"github.com/gridironx/gridchain/x/staking/types"
)

func cm45ParamsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData("custom/staking/parameters", nil)
		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusInternalServerError, common.ErrorABCIQueryFails)
			return
		}
		var params types.Params
		cliCtx.Codec.MustUnmarshalJSON(res, &params)
		cm45p := params.ToCM45Params()
		wrappedParams := types.NewWrappedParams(cm45p)
		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, wrappedParams)
	}
}

func cm45PoolHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData("custom/staking/pool", nil)
		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusInternalServerError, common.ErrorABCIQueryFails)
			return
		}
		var pool types.Pool
		cliCtx.Codec.MustUnmarshalJSON(res, &pool)
		wrappedPool := types.NewWrappedPool(pool)
		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, wrappedPool)
	}
}

func cm45DelegatorHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bech32DelAddr := mux.Vars(r)["delegatorAddr"]

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32DelAddr)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, types.CodeNoDelegatorExisted, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQueryDelegatorParams(delegatorAddr)

		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, common.CodeMarshalJSONFailed, err.Error())
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDelegator), bz)
		if err != nil {
			sdkErr := comm.ParseSDKError(err.Error())
			if sdkErr.Code == types.CodeNoDelegatorExisted {
				delegationResponses := types.NewCM45DelegationResponses(make([]types.CM45DelegationResp, 0))
				cliCtx = cliCtx.WithHeight(height)
				rest.PostProcessResponse(w, cliCtx, delegationResponses)
			} else {
				common.HandleErrorMsg(w, cliCtx, sdkErr.Code, sdkErr.Message)
			}
			return
		}
		// If res is not nil, return a formatted response.
		var delegator types.Delegator
		cliCtx.Codec.MustUnmarshalJSON(res, &delegator)
		delegationResponses := types.FormatCM45DelegationResponses(delegator)
		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, delegationResponses)
	}
}

func cm45DelegatorUnbondingDelegationsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bech32DelAddr := mux.Vars(r)["delegatorAddr"]

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32DelAddr)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, types.CodeNoDelegatorExisted, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQueryDelegatorParams(delegatorAddr)

		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, common.CodeMarshalJSONFailed, err.Error())
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryUnbondingDelegation), bz)
		if err != nil {
			sdkErr := comm.ParseSDKError(err.Error())
			if sdkErr.Code == types.CodeNoUnbondingDelegation {
				// If there is no unbonding delegation, return an empty response instead of en error.
				unbondingResponses := types.NewUnbondingResponses(make([]types.CM45UnbondingResp, 0))
				cliCtx = cliCtx.WithHeight(height)
				rest.PostProcessResponse(w, cliCtx, unbondingResponses)
			} else {
				common.HandleErrorMsg(w, cliCtx, sdkErr.Code, sdkErr.Message)
			}
			return
		}
		// If res is not nil, return a formatted response.
		var undelegationInfo types.UndelegationInfo
		cliCtx.Codec.MustUnmarshalJSON(res, &undelegationInfo)
		unbondingResponses := types.FormatCM45UnbondingResponses(undelegationInfo)
		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, unbondingResponses)
	}
}

func cm45ValidatorHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return cm45QueryValidator(cliCtx, "custom/staking/validator")
}

func cm45QueryValidator(cliCtx context.CLIContext, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bech32ValAddr := mux.Vars(r)["validatorAddr"]

		validatorAddr, err := sdk.ValAddressFromBech32(bech32ValAddr)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, types.CodeBadValidatorAddr, "validator address is invalid")
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQueryValidatorParams(validatorAddr)

		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, common.CodeMarshalJSONFailed, err.Error())
			return
		}

		res, height, err := cliCtx.QueryWithData(endpoint, bz)
		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusInternalServerError, common.ErrorABCIQueryFails)
			return
		}

		//format validator to be compatible with cosmos v0.45.1
		var val types.Validator
		cliCtx.Codec.MustUnmarshalJSON(res, &val)
		pubkey, ok := val.ConsPubKey.(ed25519.PubKeyEd25519)
		if !ok {
			common.HandleErrorMsg(w, cliCtx, common.CodeInternalError, "invalid consensus_pubkey type ")
			return
		}
		cosmosAny := types.WrapCosmosAny(pubkey[:])
		cosmosVal := types.WrapCM45Validator(val, &cosmosAny)
		wrappedValidator := types.NewWrappedValidator(cosmosVal)
		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, wrappedValidator)
	}
}

// HTTP request handler to query list of validators
func cm45ValidatorsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pr, err := rest.ParseCM45PageRequest(r)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, common.CodeArgsWithLimit, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		status := r.FormValue("status")
		if status == "" {
			status = sdk.CM45BondStatusBonded
		}
		// compatible with status of cosmos v0.45.1
		switch status {
		case sdk.CM45BondStatusBonded:
			status = sdk.BondStatusBonded
		case sdk.CM45BondStatusUnbonding:
			status = sdk.BondStatusUnbonding
		case sdk.CM45BondStatusUnbonded:
			status = sdk.BondStatusUnbonded
		}

		page := (pr.Offset / pr.Limit) + 1
		params := types.NewQueryValidatorsParams(int(page), int(pr.Limit), status)
		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, common.CodeMarshalJSONFailed, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryValidators)
		res, height, err := cliCtx.QueryWithData(route, bz)
		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusInternalServerError, common.ErrorABCIQueryFails)
			return
		}

		//format validators to be compatible with cosmos
		var vs []types.Validator
		cliCtx.Codec.MustUnmarshalJSON(res, &vs)
		filteredCosmosValidators := make([]types.CM45Validator, 0, len(vs))
		for _, val := range vs {
			pubkey, ok := val.ConsPubKey.(ed25519.PubKeyEd25519)
			if !ok {
				common.HandleErrorMsg(w, cliCtx, common.CodeInternalError, "invalid consensus_pubkey type ")
				return
			}
			cosmosAny := types.WrapCosmosAny(pubkey[:])
			cosmosVal := types.WrapCM45Validator(val, &cosmosAny)
			filteredCosmosValidators = append(filteredCosmosValidators, cosmosVal)
		}
		wrappedValidators := types.NewWrappedValidators(filteredCosmosValidators)
		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, wrappedValidators)
	}
}
